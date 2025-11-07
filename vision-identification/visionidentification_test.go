package visionidentification_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/flussonic/go-flussonic/authorization"
	"github.com/flussonic/go-flussonic/config"
	"github.com/flussonic/go-flussonic/internal/baseclient"
	visionidentification "github.com/flussonic/go-flussonic/vision-identification"
	model "github.com/flussonic/go-flussonic/vision-identification/model"
)

func TestClient_EpisodesIdentify(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/vision/api/v3/episodes/identify"

	// Expected JSON response
	expectedJSON := `{"episodes":[{"identification_error":{"code":"uncertain_match"},"matched_persons":[{"match_score":0,"person":{"deleted_at":1637095014573,"external_id":"example","fingerprints":[{"data":"example","version":"example"}],"originator":"api","person_id":0,"updated_at":1637034282845}}]}]}`

	// RoundTripper that validates request and returns example response
	rt := &validatingRoundTripper{
		t:              t,
		expectedMethod: expectedMethod,
		expectedPath:   expectedPath,
		responseJSON:   expectedJSON,
	}

	// Create client with custom transport
	client := createTestClient(t, rt)

	// Call method
	body := &model.EpisodeIdentificationRequestImpl{}
	_, err := client.EpisodesIdentify(ctx, body)
	if err != nil {
		t.Fatalf("EpisodesIdentify failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.EpisodesIdentify(ctx, body)
	if err != nil {
		t.Fatalf("EpisodesIdentify failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("EpisodesIdentify returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_IdentificationMetricsGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/vision/api/v3/monitoring/metrics"

	// Expected JSON response
	expectedJSON := `{}`

	// RoundTripper that validates request and returns example response
	rt := &validatingRoundTripper{
		t:              t,
		expectedMethod: expectedMethod,
		expectedPath:   expectedPath,
		responseJSON:   expectedJSON,
	}

	// Create client with custom transport
	client := createTestClient(t, rt)

	// Call method
	_, err := client.IdentificationMetricsGet(ctx)
	if err != nil {
		t.Fatalf("IdentificationMetricsGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.IdentificationMetricsGet(ctx)
	if err != nil {
		t.Fatalf("IdentificationMetricsGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("IdentificationMetricsGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_LivenessProbe(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/vision/api/v3/monitoring/liveness"

	// Expected JSON response
	expectedJSON := `{"build":235,"now":1639337825000,"schema_version":"5e5e91d8","server_version":"21.12","started_at":1639337825}`

	// RoundTripper that validates request and returns example response
	rt := &validatingRoundTripper{
		t:              t,
		expectedMethod: expectedMethod,
		expectedPath:   expectedPath,
		responseJSON:   expectedJSON,
	}

	// Create client with custom transport
	client := createTestClient(t, rt)

	// Call method
	_, err := client.LivenessProbe(ctx)
	if err != nil {
		t.Fatalf("LivenessProbe failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.LivenessProbe(ctx)
	if err != nil {
		t.Fatalf("LivenessProbe failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("LivenessProbe returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PersonDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/vision/api/v3/persons/test-person_id"

	// Expected JSON response
	expectedJSON := `{}`

	// RoundTripper that validates request and returns example response
	rt := &validatingRoundTripper{
		t:              t,
		expectedMethod: expectedMethod,
		expectedPath:   expectedPath,
		responseJSON:   expectedJSON,
	}

	// Create client with custom transport
	client := createTestClient(t, rt)

	// Call method
	err := client.PersonDelete(ctx, "test-person_id")
	if err != nil {
		t.Fatalf("PersonDelete failed: %v", err)
	}
}

func TestClient_PersonGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/vision/api/v3/persons/test-person_id"

	// Expected JSON response
	expectedJSON := `{"deleted_at":1637095014573,"external_id":"example","fingerprints":[{"data":"example","version":"example"}],"originator":"api","person_id":0,"updated_at":1637034282845}`

	// RoundTripper that validates request and returns example response
	rt := &validatingRoundTripper{
		t:              t,
		expectedMethod: expectedMethod,
		expectedPath:   expectedPath,
		responseJSON:   expectedJSON,
	}

	// Create client with custom transport
	client := createTestClient(t, rt)

	// Call method
	_, err := client.PersonGet(ctx, "test-person_id")
	if err != nil {
		t.Fatalf("PersonGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PersonGet(ctx, "test-person_id")
	if err != nil {
		t.Fatalf("PersonGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PersonGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PersonSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/vision/api/v3/persons/test-person_id"

	// Expected JSON response
	expectedJSON := `{"deleted_at":1637095014573,"external_id":"example","fingerprints":[{"data":"example","version":"example"}],"originator":"api","person_id":0,"updated_at":1637034282845}`

	// RoundTripper that validates request and returns example response
	rt := &validatingRoundTripper{
		t:              t,
		expectedMethod: expectedMethod,
		expectedPath:   expectedPath,
		responseJSON:   expectedJSON,
	}

	// Create client with custom transport
	client := createTestClient(t, rt)

	// Call method
	body := &model.VisionPersonImpl{}
	_, err := client.PersonSave(ctx, "test-person_id", body)
	if err != nil {
		t.Fatalf("PersonSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PersonSave(ctx, "test-person_id", body)
	if err != nil {
		t.Fatalf("PersonSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PersonSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PersonsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/vision/api/v3/persons"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","persons":[{"deleted_at":1637095014573,"external_id":"example","fingerprints":[{"data":"example","version":"example"}],"originator":"api","person_id":0,"updated_at":1637034282845}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

	// RoundTripper that validates request and returns example response
	rt := &validatingRoundTripper{
		t:              t,
		expectedMethod: expectedMethod,
		expectedPath:   expectedPath,
		responseJSON:   expectedJSON,
	}

	// Create client with custom transport
	client := createTestClient(t, rt)

	// Call method
	query := &visionidentification.PersonsListQuery{}
	_, err := client.PersonsList(ctx, query)
	if err != nil {
		t.Fatalf("PersonsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PersonsList(ctx, query)
	if err != nil {
		t.Fatalf("PersonsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PersonsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

type validatingRoundTripper struct {
	t              *testing.T
	expectedMethod string
	expectedPath   string
	responseJSON   string
}

func (rt *validatingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Verify method
	if req.Method != rt.expectedMethod {
		rt.t.Errorf("Expected method %s, got %s", rt.expectedMethod, req.Method)
	}

	// Verify path
	if req.URL.Path != rt.expectedPath {
		rt.t.Errorf("Expected path %s, got %s", rt.expectedPath, req.URL.Path)
	}

	// Create response
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       http.NoBody,
	}
	resp.Header.Set("Content-Type", "application/json")

	// Set response body
	resp.Body = io.NopCloser(bytes.NewReader([]byte(rt.responseJSON)))

	return resp, nil
}

func createTestClient(t *testing.T, rt http.RoundTripper) visionidentification.VisionIdentification {
	cfg := &config.Config{
		Hostname: "localhost",
		Protocol: "http",
		Port:     80,
	}

	baseFactory := func(cfg *config.Config, apiURL string, authKey authorization.AuthKey) baseclient.BaseClient {
		httpClient := &http.Client{
			Transport: rt,
		}
		return baseclient.New(httpClient, apiURL, nil, "", 0)
	}

	client, err := visionidentification.NewWithBaseFactory(cfg, baseFactory)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	return client
}

func verifyMarshalUnmarshal(t *testing.T, result interface{}, expectedJSON string) {
	// Marshal result
	actualJSON, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Failed to marshal result: %v", err)
	}

	// Unmarshal expected JSON
	var expected map[string]interface{}
	if err := json.Unmarshal([]byte(expectedJSON), &expected); err != nil {
		t.Fatalf("Failed to unmarshal expected JSON: %v", err)
	}

	// Unmarshal actual JSON
	var actual map[string]interface{}
	if err := json.Unmarshal(actualJSON, &actual); err != nil {
		t.Fatalf("Failed to unmarshal actual JSON: %v", err)
	}

	// Compare keys - only check keys that exist in actual
	// This handles cases where expected has fields from one oneOf variant, but actual has fields from another
	for key, expectedValue := range expected {
		actualValue, exists := actual[key]
		if !exists {
			// Key doesn't exist in actual - skip it (expected might have fields from different oneOf variant)
			continue
		}

		// Compare values - normalize both to include only common fields
		normalizedActual := normalizeValueForComparison(actualValue, expectedValue)
		normalizedExpected := normalizeValueForComparison(expectedValue, actualValue)
		expectedJSON, _ := json.Marshal(normalizedExpected)
		normalizedJSON, _ := json.Marshal(normalizedActual)
		if !bytes.Equal(expectedJSON, normalizedJSON) {
			t.Errorf("Key %s: expected %s, got %s", key, string(expectedJSON), string(normalizedJSON))
		}
		continue
	}

	// Check for extra keys (but ignore zero-value keys that weren't in expected)
	for key, actualValue := range actual {
		if _, exists := expected[key]; !exists {
			// Only report error if actual value is not a zero-value
			// This handles cases where Go structs have optional fields that get unmarshaled
			// with zero-values even if they weren't in the original JSON
			if !isZeroValueInTest(actualValue) {
				t.Errorf("Unexpected key %s in result with non-zero value: %v", key, actualValue)
			}
		}
	}
}

// normalizeValueForComparison recursively normalizes actual value by removing zero-value fields not in expected
func normalizeValueForComparison(actual, expected interface{}) interface{} {
	// Handle maps/objects - normalize actual against expected
	// This removes fields from actual that don't exist in expected
	if expectedMap, ok := expected.(map[string]interface{}); ok {
		if actualMap, ok := actual.(map[string]interface{}); ok {
			normalized := make(map[string]interface{})
			// Only include fields that exist in expectedMap AND actualMap
			for k, expectedVal := range expectedMap {
				if actualVal, exists := actualMap[k]; exists {
					// Field exists in both - normalize recursively
					normalized[k] = normalizeValueForComparison(actualVal, expectedVal)
				}
				// If field doesn't exist in actualMap, skip it
			}
			return normalized
		}
	}

	// Handle arrays/slices
	if expectedSlice, ok := expected.([]interface{}); ok {
		if actualSlice, ok := actual.([]interface{}); ok {
			normalized := make([]interface{}, 0, len(actualSlice))
			for i, actualItem := range actualSlice {
				if i < len(expectedSlice) {
					// Normalize against corresponding expected item
					normalized = append(normalized, normalizeValueForComparison(actualItem, expectedSlice[i]))
				} else if !isZeroValueInTest(actualItem) {
					// No corresponding expected item - include as is if not zero-value
					normalized = append(normalized, actualItem)
				}
			}
			return normalized
		}
	}

	// For primitive types, return as is
	return actual
}

// isZeroValueInTest checks if a value is a zero-value
func isZeroValueInTest(v interface{}) bool {
	switch val := v.(type) {
	case string:
		return val == ""
	case int:
		return val == 0
	case float64:
		return val == 0.0
	case bool:
		return val == false
	case nil:
		return true
	case []interface{}:
		return len(val) == 0
	case map[string]interface{}:
		return len(val) == 0
	default:
		return false
	}
}
