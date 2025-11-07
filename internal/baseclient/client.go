package baseclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/propagation"

	"github.com/flussonic/go-flussonic/apierror"
	"github.com/flussonic/go-flussonic/authorization"
	"github.com/flussonic/go-flussonic/reqctx"
)

type BaseClient interface {
	Request(ctx context.Context, request *http.Request, result any) error
}

// baseClient contains shared HTTP client and base URL.
type baseClient struct {
	HTTPClient    *http.Client
	BaseURL       string
	Authorization authorization.AuthKey
	Originator    string
	Retry         uint
}

const TraceIDHeader = "X-Trace-Id"

// New creates a new instance of base client.
func New(httpClient *http.Client, baseURL string, authKey authorization.AuthKey, originator string, retry uint) BaseClient {
	if httpClient == nil {
		httpClient = &http.Client{} // Use default client if not provided
	}

	c := *httpClient

	c.Transport = otelhttp.NewTransport(
		c.Transport,
		otelhttp.WithPropagators(propagation.TraceContext{}),
	)

	return &baseClient{
		HTTPClient:    &c,
		BaseURL:       baseURL,
		Authorization: authKey,
		Originator:    originator,
		Retry:         retry,
	}
}

// Request performs HTTP request and decodes response into result (for success)
// or into apierror.ErrorResponse (for errors).
// Method supports retry logic and uses timeout from httpClient.
// Status codes < 400 are considered successful, >= 400 - errors.
func (c *baseClient) Request(ctx context.Context, request *http.Request, result any) error {
	if request == nil {
		return fmt.Errorf("request is nil")
	}

	// Set authorization header if Auth is set
	if c.Authorization != nil {
		request.Header.Set("Authorization", c.Authorization.ToHeader())
	}

	if c.Originator != "" {
		request.Header.Set("X-Originator", c.Originator)
	}

	if traceID, ok := ctx.Value(reqctx.TraceIDContextKey).(string); ok {
		request.Header.Set(TraceIDHeader, traceID)
	}

	// Save request body for retry attempts
	var bodyBytes []byte
	if request.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(request.Body)
		if err != nil {
			return fmt.Errorf("failed to read request body: %w", err)
		}
		_ = request.Body.Close()
	}

	// Bind context to request
	request = request.WithContext(ctx)

	// Calculate total attempts: 1 initial + retry attempts
	// retry = 0 means no retries (1 total attempt)
	// retry = 1 means 1 retry (2 total attempts)
	totalAttempts := 1 + c.Retry
	timeout := time.Second
	var lastErr error
	var lastStatusCode int

	for attempt := range totalAttempts {
		// Restore request body for each attempt
		if len(bodyBytes) > 0 {
			request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}

		// Execute request
		response, err := c.HTTPClient.Do(request)
		if err != nil {
			lastErr = fmt.Errorf("http request failed (attempt %d/%d): %w", attempt+1, totalAttempts, err)
			lastStatusCode = http.StatusServiceUnavailable

			// If this is not the last attempt, wait before next one
			if attempt < totalAttempts-1 {
				time.Sleep(timeout)
				timeout += time.Second
			}
			continue
		}

		// Read response body
		responseBody, err := io.ReadAll(response.Body)
		_ = response.Body.Close() // Close immediately after reading
		if err != nil {
			lastErr = fmt.Errorf("failed to read response body (attempt %d/%d): %w", attempt+1, totalAttempts, err)
			lastStatusCode = response.StatusCode

			if attempt < totalAttempts-1 {
				time.Sleep(timeout)
				timeout += time.Second
			}
			continue
		}

		lastStatusCode = response.StatusCode

		// Check response code: < 400 - success, >= 400 - error
		if response.StatusCode < 400 {
			// Successful response - parse into result
			if result != nil {
				if err := json.Unmarshal(responseBody, result); err != nil {
					return fmt.Errorf("failed to unmarshal response: %w", err)
				}
			}
			return nil
		}

		// Status >= 400 - error, try to parse as ErrorResponse
		var apiErr apierror.ErrorResponse
		if err := json.Unmarshal(responseBody, &apiErr); err != nil || len(apiErr.Errors) == 0 {
			// Failed to parse as ErrorResponse - return as string
			lastErr = fmt.Errorf("API error (status %d): %s", response.StatusCode, string(responseBody))
		} else {
			// Successfully parsed ErrorResponse
			lastErr = &apiErr
		}

		// If this is not the last attempt and error code allows retry
		// (e.g., 5xx errors should be retried, 4xx - not)
		if attempt < totalAttempts-1 && response.StatusCode >= 500 {
			time.Sleep(timeout)
			timeout += time.Second
			continue
		}

		// For 4xx errors don't retry
		break
	}

	if lastErr != nil {
		return fmt.Errorf("request failed after %d attempts (last status: %d): %w", totalAttempts, lastStatusCode, lastErr)
	}

	return fmt.Errorf("request failed after %d attempts with status %d", totalAttempts, lastStatusCode)
}
