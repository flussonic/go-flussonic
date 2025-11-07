package baseclient_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	_ "golang.org/x/sys/unix"

	"github.com/flussonic/go-flussonic/apierror"
	"github.com/flussonic/go-flussonic/internal/baseclient"
	"github.com/flussonic/go-flussonic/reqctx"
)

type mockAuthKey struct {
	header string
}

func (m *mockAuthKey) ToHeader() string {
	return m.header
}

// BaseClientTestSuite is the test suite for BaseClient
type BaseClientTestSuite struct {
	suite.Suite
}

func TestBaseClientTestSuite(t *testing.T) {
	suite.Run(t, new(BaseClientTestSuite))
}

// TestNew tests the baseclient.New function
func (s *BaseClientTestSuite) TestNew() {
	s.T().Parallel()
	s.Run("with custom http client", func() {
		customClient := &http.Client{Timeout: 10 * time.Second}
		authKey := &mockAuthKey{header: "Bearer token"}
		client := baseclient.New(customClient, "http://example.com", authKey, "test-originator", 3)

		s.Require().NotNil(client)
	})

	s.Run("with nil http client", func() {
		client := baseclient.New(nil, "http://example.com", nil, "", 0)

		s.Require().NotNil(client)
	})
}

// TestRequest_NilRequest tests Request with nil request
func (s *BaseClientTestSuite) TestRequest_NilRequest() {
	s.T().Parallel()
	client := baseclient.New(nil, "http://example.com", nil, "", 0)
	ctx := context.Background()

	err := client.Request(ctx, nil, nil)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "request is nil")
}

// TestRequest_SuccessfulRequest tests successful HTTP request
func (s *BaseClientTestSuite) TestRequest_SuccessfulRequest() {
	s.T().Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check headers
		s.Require().Equal("Bearer test-token", r.Header.Get("Authorization"))
		s.Require().Equal("test-service", r.Header.Get("X-Originator"))
		s.Require().Equal("trace-123", r.Header.Get(baseclient.TraceIDHeader))

		// Return success response
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	authKey := &mockAuthKey{header: "Bearer test-token"}
	client := baseclient.New(nil, server.URL, authKey, "test-service", 0)

	ctx := context.WithValue(context.Background(), reqctx.TraceIDContextKey, "trace-123")
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	var result map[string]string
	err = client.Request(ctx, req, &result)

	s.Require().NoError(err)
	s.Require().Equal("ok", result["status"])
}

// TestRequest_SuccessfulRequestWithBody tests successful POST request with body
func (s *BaseClientTestSuite) TestRequest_SuccessfulRequestWithBody() {
	s.T().Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read and verify body
		body, _ := io.ReadAll(r.Body)
		s.Require().Equal(`{"data":"test"}`, string(body))

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "created"})
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 0)
	ctx := context.Background()

	bodyStr := `{"data":"test"}`
	req, err := http.NewRequest("POST", server.URL+"/test", strings.NewReader(bodyStr))
	s.Require().NoError(err)

	var result map[string]string
	err = client.Request(ctx, req, &result)

	s.Require().NoError(err)
	s.Require().Equal("created", result["status"])
}

// TestRequest_SuccessfulRequestWithNilResult tests successful request with nil result
func (s *BaseClientTestSuite) TestRequest_SuccessfulRequestWithNilResult() {
	s.T().Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 0)
	ctx := context.Background()
	req, err := http.NewRequest("DELETE", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)
	s.Require().NoError(err)
}

// TestRequest_ApiError tests API error response (4xx)
func (s *BaseClientTestSuite) TestRequest_ApiError() {
	s.T().Parallel()
	message := "invalid request"
	code := "invalid_request"
	status := "400"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		errResp := apierror.ErrorResponse{
			Errors: []apierror.Error{
				{Title: &message, Code: &code, Status: &status},
			},
		}
		_ = json.NewEncoder(w).Encode(errResp)
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 0)
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)

	s.Require().Error(err)
	s.Require().Contains(err.Error(), "request failed after 1 attempts")
	s.Require().Contains(err.Error(), "invalid request")
}

// TestRequest_ApiErrorNonJSON tests API error with non-JSON response
func (s *BaseClientTestSuite) TestRequest_ApiErrorNonJSON() {
	s.T().Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("plain text error"))
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 0)
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)

	s.Require().Error(err)
	s.Require().Contains(err.Error(), "plain text error")
}

// TestRequest_ApiErrorEmptyErrors tests API error with empty errors array
func (s *BaseClientTestSuite) TestRequest_ApiErrorEmptyErrors() {
	s.T().Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(apierror.ErrorResponse{Errors: []apierror.Error{}})
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 0)
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)

	s.Require().Error(err)
	s.Require().Contains(err.Error(), "API error")
}

// TestRequest_ServerError5xx tests 5xx server error without retry
func (s *BaseClientTestSuite) TestRequest_ServerError5xx() {
	s.T().Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("server error"))
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 0)
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)

	s.Require().Error(err)
	s.Require().Contains(err.Error(), "server error")
}

// TestRequest_ServerError5xxWithRetry tests 5xx server error with retry (3 retries = 4 total attempts)
func (s *BaseClientTestSuite) TestRequest_ServerError5xxWithRetry() {
	s.T().Parallel()
	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("service unavailable"))
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 3) // 3 retries = 4 total attempts
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)

	s.Require().Error(err)
	s.Require().Equal(int32(4), attempts.Load(), "should make 4 attempts (1 initial + 3 retries)")
	s.Require().Contains(err.Error(), "request failed after 4 attempts")
}

// TestRequest_ServerError5xxRecovery tests 5xx error with successful retry
func (s *BaseClientTestSuite) TestRequest_ServerError5xxRecovery() {
	s.T().Parallel()
	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count := attempts.Add(1)
		if count < 2 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 3) // 3 retries = up to 4 attempts
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	var result map[string]string
	err = client.Request(ctx, req, &result)

	s.Require().NoError(err)
	s.Require().Equal(int32(2), attempts.Load())
	s.Require().Equal("ok", result["status"])
}

// TestRequest_ClientError4xxNoRetry tests 4xx error doesn't retry
func (s *BaseClientTestSuite) TestRequest_ClientError4xxNoRetry() {
	s.T().Parallel()
	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not found"))
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 3) // 3 retries configured
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)

	s.Require().Error(err)
	s.Require().Equal(int32(1), attempts.Load(), "4xx errors should not retry")
}

// TestRequest_NetworkError tests network error with retry
func (s *BaseClientTestSuite) TestRequest_NetworkError() {
	s.T().Parallel()
	// Create a client that will fail with network error
	client := baseclient.New(nil, "http://invalid-host-that-does-not-exist.local:9999", nil, "", 1) // 1 retry = 2 attempts
	ctx := context.Background()
	req, err := http.NewRequest("GET", "http://invalid-host-that-does-not-exist.local:9999/test", http.NoBody)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)

	s.Require().Error(err)
	s.Require().Contains(err.Error(), "request failed after 2 attempts")
	s.Require().Contains(err.Error(), "http request failed")
}

// TestRequest_BodyReadError tests error when reading request body
func (s *BaseClientTestSuite) TestRequest_BodyReadError() {
	s.T().Parallel()
	client := baseclient.New(nil, "http://example.com", nil, "", 0)
	ctx := context.Background()

	// Create a reader that always fails
	errReader := &errorReader{err: fmt.Errorf("read error")}
	req, err := http.NewRequest("POST", "http://example.com/test", errReader)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)

	s.Require().Error(err)
	s.Require().Contains(err.Error(), "failed to read request body")
}

// TestRequest_ResponseBodyReadError tests error when reading response body
func (s *BaseClientTestSuite) TestRequest_ResponseBodyReadError() {
	s.T().Parallel()
	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		// Set Content-Length but don't write the body
		w.Header().Set("Content-Length", "1000")
		// Close connection prematurely
		hj, ok := w.(http.Hijacker)
		if ok {
			conn, _, _ := hj.Hijack()
			_ = conn.Close()
		}
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 1) // 1 retry = 2 attempts
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)

	s.Require().Error(err)
	s.Require().GreaterOrEqual(attempts.Load(), int32(2), "should have retried")
}

// TestRequest_ResponseBodyReadErrorWithRetry tests response body read error that retries
func (s *BaseClientTestSuite) TestRequest_ResponseBodyReadErrorWithRetry() {
	s.T().Parallel()
	var attempts atomic.Int32

	// Create a custom transport that simulates body read failures
	transport := &brokenBodyTransport{
		attempts: &attempts,
	}

	client := baseclient.New(&http.Client{Transport: transport}, "http://example.com", nil, "", 2) // 2 retries = 3 attempts
	ctx := context.Background()
	req, err := http.NewRequest("GET", "http://example.com/test", http.NoBody)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)

	s.Require().Error(err)
	s.Require().Equal(int32(3), attempts.Load())
	s.Require().Contains(err.Error(), "failed to read response body")
}

// TestRequest_UnmarshalError tests error when unmarshaling response
func (s *BaseClientTestSuite) TestRequest_UnmarshalError() {
	s.T().Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 0)
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	var result map[string]string
	err = client.Request(ctx, req, &result)

	s.Require().Error(err)
	s.Require().Contains(err.Error(), "failed to unmarshal response")
}

// TestRequest_RetryWithBody tests retry with request body preservation
func (s *BaseClientTestSuite) TestRequest_RetryWithBody() {
	s.T().Parallel()
	var attempts atomic.Int32
	var receivedBodies []string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count := attempts.Add(1)
		body, _ := io.ReadAll(r.Body)
		receivedBodies = append(receivedBodies, string(body))

		if count < 2 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 3) // 3 retries = up to 4 attempts
	ctx := context.Background()

	bodyStr := `{"data":"test"}`
	req, err := http.NewRequest("POST", server.URL+"/test", strings.NewReader(bodyStr))
	s.Require().NoError(err)

	var result map[string]string
	err = client.Request(ctx, req, &result)

	s.Require().NoError(err)
	s.Require().Equal(int32(2), attempts.Load())
	s.Require().Len(receivedBodies, 2)
	for i, body := range receivedBodies {
		s.Require().Equal(bodyStr, body, "body should be preserved for attempt %d", i+1)
	}
}

// TestRequest_NoAuthAndNoOriginator tests request without auth and originator
func (s *BaseClientTestSuite) TestRequest_NoAuthAndNoOriginator() {
	s.T().Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Require().Empty(r.Header.Get("Authorization"))
		s.Require().Empty(r.Header.Get("X-Originator"))
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 0)
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	var result map[string]string
	err = client.Request(ctx, req, &result)

	s.Require().NoError(err)
}

// TestRequest_NoTraceIDInContext tests request without trace ID in context
func (s *BaseClientTestSuite) TestRequest_NoTraceIDInContext() {
	s.T().Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Require().Empty(r.Header.Get(baseclient.TraceIDHeader))
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 0)
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	var result map[string]string
	err = client.Request(ctx, req, &result)

	s.Require().NoError(err)
}

// TestRequest_ZeroRetry tests that zero retry means no retries (1 attempt)
func (s *BaseClientTestSuite) TestRequest_ZeroRetry() {
	s.T().Parallel()
	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 0) // 0 retries = 1 attempt
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	var result map[string]string
	err = client.Request(ctx, req, &result)

	s.Require().NoError(err)
	s.Require().Equal(int32(1), attempts.Load(), "zero retry should result in 1 attempt")
}

// TestRequest_OneRetry tests that retry = 1 means 1 retry (2 total attempts)
func (s *BaseClientTestSuite) TestRequest_OneRetry() {
	s.T().Parallel()
	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("service unavailable"))
	}))
	defer server.Close()

	client := baseclient.New(nil, server.URL, nil, "", 1) // 1 retry = 2 attempts
	ctx := context.Background()
	req, err := http.NewRequest("GET", server.URL+"/test", http.NoBody)
	s.Require().NoError(err)

	err = client.Request(ctx, req, nil)

	s.Require().Error(err)
	s.Require().Equal(int32(2), attempts.Load(), "1 retry should result in 2 attempts")
	s.Require().Contains(err.Error(), "request failed after 2 attempts")
}

// errorReader is a helper type that always returns an error when Read is called
type errorReader struct {
	err error
}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, e.err
}

func (e *errorReader) Close() error {
	return nil
}

// brokenBodyTransport is a custom http.RoundTripper that returns responses with broken body readers
type brokenBodyTransport struct {
	attempts *atomic.Int32
}

func (t *brokenBodyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.attempts.Add(1)

	// Return a response with a body that fails to read
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       &errorReader{err: fmt.Errorf("body read error")},
		Header:     make(http.Header),
	}, nil
}
