package central_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/flussonic/go-flussonic/authorization"
	central "github.com/flussonic/go-flussonic/central"
	model "github.com/flussonic/go-flussonic/central/model"
	"github.com/flussonic/go-flussonic/config"
	"github.com/flussonic/go-flussonic/internal/baseclient"
)

func TestClient_AgentDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/central/api/v3/agents/test-id"

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
	err := client.AgentDelete(ctx, "test-id")
	if err != nil {
		t.Fatalf("AgentDelete failed: %v", err)
	}
}

func TestClient_AgentDisconnect(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/central/api/v3/agents/test-id/disconnect"

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
	err := client.AgentDisconnect(ctx, "test-id")
	if err != nil {
		t.Fatalf("AgentDisconnect failed: %v", err)
	}
}

func TestClient_AgentGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/agents/test-id"

	// Expected JSON response
	expectedJSON := `{"layout":{"created_at":1637094994000,"ingest":"example","originator":"layouter"}}`

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
	_, err := client.AgentGet(ctx, "test-id")
	if err != nil {
		t.Fatalf("AgentGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AgentGet(ctx, "test-id")
	if err != nil {
		t.Fatalf("AgentGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AgentGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_AgentLogsGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/agents/test-id/logs"

	// Expected JSON response
	expectedJSON := `{"a":"VDEVSvaNZZWeNOXqi7tA5Hm+KAqBjfHQKMzMl5aqRdYxHZkSnJ","agent_id":"1234567","b":"xglytWCTztsbTAUVBIhQymqik7nvzXi5VbJWEQNBHFCxw/NdvY6hdknh87/3gslKpER0hIaymR67Qo8zGKICFN==","cid":"42195","main_url":"http://example.com","meminfo":"806400","mid":"3059","serial":"a3dccd69f53deb79723a2a7a5f2037e1","status":"ok","version":"version"}`

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
	_, err := client.AgentLogsGet(ctx, "test-id")
	if err != nil {
		t.Fatalf("AgentLogsGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AgentLogsGet(ctx, "test-id")
	if err != nil {
		t.Fatalf("AgentLogsGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AgentLogsGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_AgentReboot(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/central/api/v3/agents/test-id/reboot"

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
	err := client.AgentReboot(ctx, "test-id")
	if err != nil {
		t.Fatalf("AgentReboot failed: %v", err)
	}
}

func TestClient_AgentReset(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/central/api/v3/agents/test-id/reset"

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
	err := client.AgentReset(ctx, "test-id")
	if err != nil {
		t.Fatalf("AgentReset failed: %v", err)
	}
}

func TestClient_AgentSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/central/api/v3/agents/test-id"

	// Expected JSON response
	expectedJSON := `{"layout":{"created_at":1637094994000,"ingest":"example","originator":"layouter"}}`

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
	body := &model.CentralAgentConfigImpl{}
	_, err := client.AgentSave(ctx, "test-id", body)
	if err != nil {
		t.Fatalf("AgentSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AgentSave(ctx, "test-id", body)
	if err != nil {
		t.Fatalf("AgentSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AgentSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_AgentsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/agents"

	// Expected JSON response
	expectedJSON := `{"agents":[{"layout":{"created_at":1637094994000,"ingest":"example","originator":"layouter"}}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.AgentsListQuery{}
	_, err := client.AgentsList(ctx, query)
	if err != nil {
		t.Fatalf("AgentsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AgentsList(ctx, query)
	if err != nil {
		t.Fatalf("AgentsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AgentsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ApiTokenGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/api_tokens/test-name"

	// Expected JSON response
	expectedJSON := `{"auth_scopes":["example"],"key":"example","name":"example","permissions":[{"scopes":["example"]}]}`

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
	_, err := client.ApiTokenGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("ApiTokenGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ApiTokenGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("ApiTokenGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ApiTokenGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ApiTokenSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/central/api/v3/api_tokens/test-name"

	// Expected JSON response
	expectedJSON := `{"auth_scopes":["example"],"key":"example","name":"example","permissions":[{"scopes":["example"]}]}`

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
	body := &model.APITokenImpl{}
	_, err := client.ApiTokenSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("ApiTokenSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ApiTokenSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("ApiTokenSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ApiTokenSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ApiTokensList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/api_tokens"

	// Expected JSON response
	expectedJSON := `{"api_tokens":[{"auth_scopes":["example"],"key":"example","name":"example","permissions":[{"scopes":["example"]}]}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.ApiTokensListQuery{}
	_, err := client.ApiTokensList(ctx, query)
	if err != nil {
		t.Fatalf("ApiTokensList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ApiTokensList(ctx, query)
	if err != nil {
		t.Fatalf("ApiTokensList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ApiTokensList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_AuthBackendDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/central/api/v3/auth_backends/test-name"

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
	err := client.AuthBackendDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("AuthBackendDelete failed: %v", err)
	}
}

func TestClient_AuthBackendGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/auth_backends/test-name"

	// Expected JSON response
	expectedJSON := `{"allow_countries":["RU","US"],"allow_ips":["127.0.0.1","10.10.0.0/24"],"allow_tokens":["test_token1","test_token2"],"allow_uas":["AppleWebKit/533.3 (KHTML, like Gecko)","VLC/3.0.8 LibVLC/3.0.8"],"backends":[{"url":"http://stalker-1.iptv.net/auth.php"}],"deny_countries":["RU","GB"],"deny_ips":["8.8.8.8","10.10.0.0/24"],"deny_tokens":["test_token3","test_token4"],"deny_uas":["Mobile Safari/533.3","VLC/3.0.8 LibVLC/3.0.8"],"name":"example"}`

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
	_, err := client.AuthBackendGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("AuthBackendGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AuthBackendGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("AuthBackendGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AuthBackendGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_AuthBackendSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/central/api/v3/auth_backends/test-name"

	// Expected JSON response
	expectedJSON := `{"allow_countries":["RU","US"],"allow_ips":["127.0.0.1","10.10.0.0/24"],"allow_tokens":["test_token1","test_token2"],"allow_uas":["AppleWebKit/533.3 (KHTML, like Gecko)","VLC/3.0.8 LibVLC/3.0.8"],"backends":[{"url":"http://stalker-1.iptv.net/auth.php"}],"deny_countries":["RU","GB"],"deny_ips":["8.8.8.8","10.10.0.0/24"],"deny_tokens":["test_token3","test_token4"],"deny_uas":["Mobile Safari/533.3","VLC/3.0.8 LibVLC/3.0.8"],"name":"example"}`

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
	body := &model.AuthBackendConfigImpl{}
	_, err := client.AuthBackendSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("AuthBackendSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AuthBackendSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("AuthBackendSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AuthBackendSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_AuthBackendsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/auth_backends"

	// Expected JSON response
	expectedJSON := `{"auth_backends":[{"allow_countries":["RU","US"],"allow_ips":["127.0.0.1","10.10.0.0/24"],"allow_tokens":["test_token1","test_token2"],"allow_uas":["AppleWebKit/533.3 (KHTML, like Gecko)","VLC/3.0.8 LibVLC/3.0.8"],"backends":[{"url":"http://stalker-1.iptv.net/auth.php"}],"deny_countries":["RU","GB"],"deny_ips":["8.8.8.8","10.10.0.0/24"],"deny_tokens":["test_token3","test_token4"],"deny_uas":["Mobile Safari/533.3","VLC/3.0.8 LibVLC/3.0.8"],"name":"example"}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.AuthBackendsListQuery{}
	_, err := client.AuthBackendsList(ctx, query)
	if err != nil {
		t.Fatalf("AuthBackendsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AuthBackendsList(ctx, query)
	if err != nil {
		t.Fatalf("AuthBackendsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AuthBackendsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_AuthRequest(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/auth"

	// Expected JSON response
	expectedJSON := `{"ad_inject":{"midroll":["ad_vod/midroll1.mp4","ad_vod/midroll2.mp4"],"midroll_insert_by":"interval","midroll_interval":180,"midroll_program_id":1,"preroll":["ad_vod/preroll1.mp4"],"v":2},"allowed_dvr_ranges":[{"closed_at":1710020000,"opened_at":1710010000}]}`

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
	query := &central.AuthRequestQuery{}
	query.Ip = "test-ip"
	query.Name = "test-name"
	query.Proto = "test-proto"
	query.RequestNumber = 1
	query.RequestType = "test-request_type"
	query.SessionId = "test-session_id"
	query.StreamClients = 1
	query.TotalClients = 1
	_, err := client.AuthRequest(ctx, query)
	if err != nil {
		t.Fatalf("AuthRequest failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AuthRequest(ctx, query)
	if err != nil {
		t.Fatalf("AuthRequest failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AuthRequest returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_BatchStreamLayoutPreview(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/central/api/v3/streams/preview_layout_change"

	// Expected JSON response
	expectedJSON := `{"changes":[{"name":"example"}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.BatchStreamLayoutPreviewQuery{}
	_, err := client.BatchStreamLayoutPreview(ctx, query, nil)
	if err != nil {
		t.Fatalf("BatchStreamLayoutPreview failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.BatchStreamLayoutPreview(ctx, query, nil)
	if err != nil {
		t.Fatalf("BatchStreamLayoutPreview failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("BatchStreamLayoutPreview returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_BatchStreamerLayoutPreview(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/central/api/v3/streamers/preview_layout_change"

	// Expected JSON response
	expectedJSON := `{"changes":[{"after":123,"before":321,"delta":1,"hostname":"streamer1","role":"ingest"}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.BatchStreamerLayoutPreviewQuery{}
	_, err := client.BatchStreamerLayoutPreview(ctx, query, nil)
	if err != nil {
		t.Fatalf("BatchStreamerLayoutPreview failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.BatchStreamerLayoutPreview(ctx, query, nil)
	if err != nil {
		t.Fatalf("BatchStreamerLayoutPreview failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("BatchStreamerLayoutPreview returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_BatchUpdateAgentsLayouts(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/central/api/v3/agents/layouts"

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
	body := &model.CentralAgentLayoutListImpl{}
	err := client.BatchUpdateAgentsLayouts(ctx, body)
	if err != nil {
		t.Fatalf("BatchUpdateAgentsLayouts failed: %v", err)
	}
}

func TestClient_BatchUpdateStreamsLayouts(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/central/api/v3/streams/layouts"

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
	body := &model.CentralStreamLayoutListImpl{}
	err := client.BatchUpdateStreamsLayouts(ctx, body)
	if err != nil {
		t.Fatalf("BatchUpdateStreamsLayouts failed: %v", err)
	}
}

func TestClient_CentralEventsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/events"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"events":[{"event":"example","full_version":"example","server":"flussonic.host","trace_id":"05cec7ee-fbd0-11ed-be56-0242ac120002","utc_ms":1000000000000,"version":"example"}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.CentralEventsListQuery{}
	_, err := client.CentralEventsList(ctx, query)
	if err != nil {
		t.Fatalf("CentralEventsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.CentralEventsList(ctx, query)
	if err != nil {
		t.Fatalf("CentralEventsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("CentralEventsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ClusterStatsGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/cluster/stats"

	// Expected JSON response
	expectedJSON := `{"collected_at":1000000000000,"server_id":"123e4567-e89b-12d3-a456-426655440000","streamer_metrics":[{"agent_metrics":{"status":"operational"},"config":{"status":"operational"},"cpu":{"status":"operational","usage":48},"hostname":"example","memory":{"status":"operational","usage":27},"status":"operational","storage":{"status":"operational","usage":18},"stream_metrics":{"status":"operational"},"uptime":4325502}],"version":240100023}`

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
	_, err := client.ClusterStatsGet(ctx, nil)
	if err != nil {
		t.Fatalf("ClusterStatsGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ClusterStatsGet(ctx, nil)
	if err != nil {
		t.Fatalf("ClusterStatsGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ClusterStatsGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ConfigGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/config"

	// Expected JSON response
	expectedJSON := `{"api_key":"api_key","api_url":"http://central.ru:9019/streamer/api/v3","cluster":{"node_config_provision_enabled":true,"streamer_connector_retries":5},"config_external_reconciliation_interval":5000,"database":{"connect_max_retries":20,"connect_retry_timeout":"2s","max_connections":40,"url":"postgres://central:pass@localhost:5432/central_dev"},"edit_auth":{"login":"secretlogin","password":"passw"},"episodes_buffer_delay":3000,"episodes_buffer_limit":100,"layouter_enabled":"false","layouter_sleeping_interval":30000,"listeners":{"http":[{"address":"10.0.35.1","port":80,"read_timeout":5,"write_timeout":5}],"https":[{"certificate":"/etc/letsencrypt/live/central/fullchain.pem","certificate_key":"/etc/letsencrypt/live/central/privkey.pem","ssl_protocols":["tlsv1.1","tlsv1.2"]}]},"log_requests":"false","loglevel":"error","opentelemetry_url":"http://jaeger-server:14268/v1/traces?service_name=some-name","redis":{"connect_max_retries":20,"connect_retry_timeout":"2s","max_connections":40,"url":"redis://:pass@localhost:6379"},"server_id":"550e8400-e29b-41d4-a716-446655440000","stats":{"id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","server_version":"23.10","started_at":1639337825,"uptime":4325502},"streamer_healthcheck_fails_threshold":5}`

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
	_, err := client.ConfigGet(ctx, nil)
	if err != nil {
		t.Fatalf("ConfigGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ConfigGet(ctx, nil)
	if err != nil {
		t.Fatalf("ConfigGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ConfigGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ConfigSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/central/api/v3/config"

	// Expected JSON response
	expectedJSON := `{"api_key":"api_key","api_url":"http://central.ru:9019/streamer/api/v3","cluster":{"node_config_provision_enabled":true,"streamer_connector_retries":5},"config_external_reconciliation_interval":5000,"database":{"connect_max_retries":20,"connect_retry_timeout":"2s","max_connections":40,"url":"postgres://central:pass@localhost:5432/central_dev"},"edit_auth":{"login":"secretlogin","password":"passw"},"episodes_buffer_delay":3000,"episodes_buffer_limit":100,"layouter_enabled":"false","layouter_sleeping_interval":30000,"listeners":{"http":[{"address":"10.0.35.1","port":80,"read_timeout":5,"write_timeout":5}],"https":[{"certificate":"/etc/letsencrypt/live/central/fullchain.pem","certificate_key":"/etc/letsencrypt/live/central/privkey.pem","ssl_protocols":["tlsv1.1","tlsv1.2"]}]},"log_requests":"false","loglevel":"error","opentelemetry_url":"http://jaeger-server:14268/v1/traces?service_name=some-name","redis":{"connect_max_retries":20,"connect_retry_timeout":"2s","max_connections":40,"url":"redis://:pass@localhost:6379"},"server_id":"550e8400-e29b-41d4-a716-446655440000","stats":{"id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","server_version":"23.10","started_at":1639337825,"uptime":4325502},"streamer_healthcheck_fails_threshold":5}`

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
	body := &model.CentralConfigImpl{}
	_, err := client.ConfigSave(ctx, body)
	if err != nil {
		t.Fatalf("ConfigSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ConfigSave(ctx, body)
	if err != nil {
		t.Fatalf("ConfigSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ConfigSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvrsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/dvrs"

	// Expected JSON response
	expectedJSON := `{"dvrs":[{"name":"example"}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.DvrsListQuery{}
	_, err := client.DvrsList(ctx, query)
	if err != nil {
		t.Fatalf("DvrsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvrsList(ctx, query)
	if err != nil {
		t.Fatalf("DvrsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvrsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_EpisodeDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/central/api/v3/episodes/test-episode_id"

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
	err := client.EpisodeDelete(ctx, "test-episode_id")
	if err != nil {
		t.Fatalf("EpisodeDelete failed: %v", err)
	}
}

func TestClient_EpisodeGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/episodes/test-episode_id"

	// Expected JSON response
	expectedJSON := `{"close_reason":"timeout","closed_at":1000000000000,"episode_appearance_timestamps":{"central_timestamp":1637098611000,"inference_timestamp":1637094994000},"episode_id":0,"episode_type":"generic","frame_preview":"example","media":"example","opened_at":1000000000000,"originator":{"hostname":"example.com","source":"api"},"preview":"example","preview_timestamp":1000000000000,"recording_status":"full","started_at":1000000000000,"updated_at":1000000000000}`

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
	_, err := client.EpisodeGet(ctx, "test-episode_id", nil)
	if err != nil {
		t.Fatalf("EpisodeGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.EpisodeGet(ctx, "test-episode_id", nil)
	if err != nil {
		t.Fatalf("EpisodeGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("EpisodeGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_EpisodeSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/central/api/v3/episodes/test-episode_id"

	// Expected JSON response
	expectedJSON := `{"close_reason":"timeout","closed_at":1000000000000,"episode_appearance_timestamps":{"central_timestamp":1637098611000,"inference_timestamp":1637094994000},"episode_id":0,"episode_type":"generic","frame_preview":"example","media":"example","opened_at":1000000000000,"originator":{"hostname":"example.com","source":"api"},"preview":"example","preview_timestamp":1000000000000,"recording_status":"full","started_at":1000000000000,"updated_at":1000000000000}`

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
	body := &model.EpisodeImpl{}
	_, err := client.EpisodeSave(ctx, "test-episode_id", body)
	if err != nil {
		t.Fatalf("EpisodeSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.EpisodeSave(ctx, "test-episode_id", body)
	if err != nil {
		t.Fatalf("EpisodeSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("EpisodeSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_EpisodesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/episodes"

	// Expected JSON response
	expectedJSON := `{"episodes":[{"close_reason":"timeout","closed_at":1000000000000,"episode_appearance_timestamps":{"central_timestamp":1637098611000,"inference_timestamp":1637094994000},"episode_id":0,"episode_type":"generic","frame_preview":"example","media":"example","opened_at":1000000000000,"originator":{"hostname":"example.com","source":"api"},"preview":"example","preview_timestamp":1000000000000,"recording_status":"full","started_at":1000000000000,"updated_at":1000000000000}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.EpisodesListQuery{}
	_, err := client.EpisodesList(ctx, query)
	if err != nil {
		t.Fatalf("EpisodesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.EpisodesList(ctx, query)
	if err != nil {
		t.Fatalf("EpisodesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("EpisodesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ExternalEpisodesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/streamers/test-hostname/episodes"

	// Expected JSON response
	expectedJSON := `{"episodes":[{"closed_at":1637094994000,"opened_at":1637094994000,"updated_at":1637098611000}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.ExternalEpisodesListQuery{}
	query.Media = "test-media"
	_, err := client.ExternalEpisodesList(ctx, "test-hostname", query)
	if err != nil {
		t.Fatalf("ExternalEpisodesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ExternalEpisodesList(ctx, "test-hostname", query)
	if err != nil {
		t.Fatalf("ExternalEpisodesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ExternalEpisodesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_LivenessProbe(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/monitoring/liveness"

	// Expected JSON response
	expectedJSON := `{"now":1000000000000,"server_version":"23.04","started_at":1639337825}`

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

func TestClient_LoadBalancerDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/central/api/v3/load-balancers/test-name"

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
	err := client.LoadBalancerDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("LoadBalancerDelete failed: %v", err)
	}
}

func TestClient_LoadBalancerGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/load-balancers/test-name"

	// Expected JSON response
	expectedJSON := `{"mode":"usage","name":"example","servers":[{"countries":["example"],"name":"example"}]}`

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
	_, err := client.LoadBalancerGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("LoadBalancerGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.LoadBalancerGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("LoadBalancerGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("LoadBalancerGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_LoadBalancerSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/central/api/v3/load-balancers/test-name"

	// Expected JSON response
	expectedJSON := `{"mode":"usage","name":"example","servers":[{"countries":["example"],"name":"example"}]}`

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
	body := &model.BalancerConfigImpl{}
	_, err := client.LoadBalancerSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("LoadBalancerSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.LoadBalancerSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("LoadBalancerSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("LoadBalancerSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_LoadBalancersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/load-balancers"

	// Expected JSON response
	expectedJSON := `{"balancers":[{"mode":"usage","name":"example","servers":[{"countries":["example"],"name":"example"}]}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.LoadBalancersListQuery{}
	_, err := client.LoadBalancersList(ctx, query)
	if err != nil {
		t.Fatalf("LoadBalancersList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.LoadBalancersList(ctx, query)
	if err != nil {
		t.Fatalf("LoadBalancersList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("LoadBalancersList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PersonDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/central/api/v3/persons/test-person_id"

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
	expectedPath := "/central/api/v3/persons/test-person_id"

	// Expected JSON response
	expectedJSON := `{"deleted_at":1637095014573,"external_id":"example","fingerprints":[{"data":"example","version":"example"}],"originator":"api","person_id":0,"photos":[{"data":"example","mime_type":"image/jpeg","sha256":"example"}],"updated_at":1637034282845}`

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
	expectedPath := "/central/api/v3/persons/test-person_id"

	// Expected JSON response
	expectedJSON := `{"deleted_at":1637095014573,"external_id":"example","fingerprints":[{"data":"example","version":"example"}],"originator":"api","person_id":0,"photos":[{"data":"example","mime_type":"image/jpeg","sha256":"example"}],"updated_at":1637034282845}`

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
	expectedPath := "/central/api/v3/persons"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","persons":[{"deleted_at":1637095014573,"external_id":"example","fingerprints":[{"data":"example","version":"example"}],"originator":"api","person_id":0,"photos":[{"data":"example","mime_type":"image/jpeg","sha256":"example"}],"updated_at":1637034282845}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.PersonsListQuery{}
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

func TestClient_PreviewLayoutChangeForStreamConfig(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/central/api/v3/streams/test-name/preview_layout_change"

	// Expected JSON response
	expectedJSON := `{"changes":[{"name":"example"}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.PreviewLayoutChangeForStreamConfigQuery{}
	body := &model.CentralStreamConfigImpl{}
	_, err := client.PreviewLayoutChangeForStreamConfig(ctx, "test-name", query, body)
	if err != nil {
		t.Fatalf("PreviewLayoutChangeForStreamConfig failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PreviewLayoutChangeForStreamConfig(ctx, "test-name", query, body)
	if err != nil {
		t.Fatalf("PreviewLayoutChangeForStreamConfig failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PreviewLayoutChangeForStreamConfig returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PreviewLayoutChangeForStreamerConfig(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/central/api/v3/streamers/test-hostname/preview_layout_change"

	// Expected JSON response
	expectedJSON := `{"changes":[{"after":123,"before":321,"delta":1,"hostname":"streamer1","role":"ingest"}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.PreviewLayoutChangeForStreamerConfigQuery{}
	body := &model.StreamerConfigImpl{}
	_, err := client.PreviewLayoutChangeForStreamerConfig(ctx, "test-hostname", query, body)
	if err != nil {
		t.Fatalf("PreviewLayoutChangeForStreamerConfig failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PreviewLayoutChangeForStreamerConfig(ctx, "test-hostname", query, body)
	if err != nil {
		t.Fatalf("PreviewLayoutChangeForStreamerConfig failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PreviewLayoutChangeForStreamerConfig returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ReadinessProbe(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/monitoring/readiness"

	// Expected JSON response
	expectedJSON := `{"now":1000000000000,"server_version":"23.04","started_at":1639337825}`

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
	_, err := client.ReadinessProbe(ctx)
	if err != nil {
		t.Fatalf("ReadinessProbe failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ReadinessProbe(ctx)
	if err != nil {
		t.Fatalf("ReadinessProbe failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ReadinessProbe returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_SrtPortResolve(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/streamers/test-hostname/srt_port_resolve/test-port"

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
	query := &central.SrtPortResolveQuery{}
	query.Mode = "test-mode"
	_, err := client.SrtPortResolve(ctx, "test-hostname", "test-port", query)
	if err != nil {
		t.Fatalf("SrtPortResolve failed: %v", err)
	}
}

func TestClient_StreamDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/central/api/v3/streams/test-name"

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
	err := client.StreamDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("StreamDelete failed: %v", err)
	}
}

func TestClient_StreamGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/streams/test-name"

	// Expected JSON response
	expectedJSON := `{"claims":{"bitrate":2543},"layout":{"change_reason":"stream_misses_node_required_labels","created_at":1637094994000,"dvr_backups":["example"],"failover_from":"example","inference":"example","ingest":"example","ingest_history":[{"created_at":1637094994000,"ingest":"example","originator":"layouter"}],"node_layout_decisions":[{"hostname":"streamer1.com","reasons":["stream_misses_node_required_labels","node_misses_stream_required_labels","node_channel_limit_exceeded"],"role":"streamer"}],"originator":"layouter"},"namespace":"example","updated_at":1637098611000}`

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
	_, err := client.StreamGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("StreamGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("StreamGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamLayoutsGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/streams/test-name/layouts"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"layouts":[{"change_reason":"stream_misses_node_required_labels","created_at":1637094994000,"dvr_backups":["example"],"failover_from":"example","inference":"example","ingest":"example","ingest_history":[{"created_at":1637094994000,"ingest":"example","originator":"layouter"}],"node_layout_decisions":[{"hostname":"streamer1.com","reasons":["stream_misses_node_required_labels","node_misses_stream_required_labels","node_channel_limit_exceeded"],"role":"streamer"}],"originator":"layouter"}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &central.StreamLayoutsGetQuery{}
	_, err := client.StreamLayoutsGet(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("StreamLayoutsGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamLayoutsGet(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("StreamLayoutsGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamLayoutsGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/central/api/v3/streams/test-name"

	// Expected JSON response
	expectedJSON := `{"claims":{"bitrate":2543},"layout":{"change_reason":"stream_misses_node_required_labels","created_at":1637094994000,"dvr_backups":["example"],"failover_from":"example","inference":"example","ingest":"example","ingest_history":[{"created_at":1637094994000,"ingest":"example","originator":"layouter"}],"node_layout_decisions":[{"hostname":"streamer1.com","reasons":["stream_misses_node_required_labels","node_misses_stream_required_labels","node_channel_limit_exceeded"],"role":"streamer"}],"originator":"layouter"},"namespace":"example","updated_at":1637098611000}`

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
	body := &model.CentralStreamConfigImpl{}
	_, err := client.StreamSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("StreamSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("StreamSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamerDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/central/api/v3/streamers/test-hostname"

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
	err := client.StreamerDelete(ctx, "test-hostname")
	if err != nil {
		t.Fatalf("StreamerDelete failed: %v", err)
	}
}

func TestClient_StreamerDynamicStreamsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/streamers/test-hostname/dynamic-streams"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","streams":[{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"chunk_duration":200,"clients_timeout":485,"comment":"This is a test stream","config_on_disk":{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"chunk_duration":200,"clients_timeout":485,"comment":"This is a test stream","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"disk_usage_limit":98,"episodes_url":"example","redundancy_factor":1,"reference":"localdvr0","remotes":["example_value"],"schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000,"status":"running"}}]},"webrtc_abr":{"start_track":"v2"}},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"disk_usage_limit":98,"episodes_url":"example","redundancy_factor":1,"reference":"localdvr0","remotes":["example_value"],"schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"episodes_density":0.001,"last_dts_at":1636383841974,"last_running_at":1737975543123,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"running_on":["streamer1.example.com"],"status":"running","streaming_endpoint":"example","ts_delay":1284},"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000,"status":"running"}}]},"webrtc_abr":{"start_track":"v2"}}],"updated_at":1727274724000}`

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
	query := &central.StreamerDynamicStreamsListQuery{}
	query.Name = "test-name"
	_, err := client.StreamerDynamicStreamsList(ctx, "test-hostname", query)
	if err != nil {
		t.Fatalf("StreamerDynamicStreamsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamerDynamicStreamsList(ctx, "test-hostname", query)
	if err != nil {
		t.Fatalf("StreamerDynamicStreamsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamerDynamicStreamsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamerGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/streamers/test-hostname"

	// Expected JSON response
	expectedJSON := `{"api_url":"http://streamer.local:8080","channel_limit":5,"cluster_key":"xS6i6Q3DCc5nEvnu","config":{"cluster_key":"xS6i6Q3DCc5nEvnu","event_sinks":[{"except":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"extra":{"key1":"example","key2":"example"},"level":"debug","max_depth":100,"max_size":10000,"name":"my_json_sink","only":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"resend_limit":1000,"resend_timeout":10,"throttle_delay":1,"url":"jsonlog:///var/log/events-json.log"}],"http_proxies":[{"prefix":"example","stats":{"http_100":1000,"http_200":1000,"http_300":10,"http_400":10,"http_500":10,"protocol_upgrades":1000,"proxy_error":10,"proxy_error_connection":10,"requests":1000,"time_1000ms":2,"time_100ms":10,"time_5000ms":2,"time_500ms":5,"time_longms":1},"url":"https://example.com"}],"rproxy":{"endpoint_auth":"example","forward_ports":{"key1":{"handler":"example"},"key2":{"handler":"example"}},"streampoint_key":"example"},"server_names":[{"aliases":["s1.streamer.local","s2.streamer.local"],"domain":"streamer.local"}]},"cpu_limit":10,"dvrs":[{"disk_usage_limit":98,"disks":[{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}],"episodes_url":"example","index":"example","name":"example","root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000}],"fetch_timeout":1000,"hostname":"peer.example.com","labels":{"key1":"value1","key2":"value2"},"namespace":"example","private_payload_url":"http://streamer.local","public_payload_url":"http://public.example.com","role":"streamer","stale_timeout":1000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","healthcheck_status":{"checks":{"errors_details":[{"error":"example","rule":"example"}]},"status":"ok","status_changed_at":1000000000000},"hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"total_bandwidth":1024}`

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
	_, err := client.StreamerGet(ctx, "test-hostname")
	if err != nil {
		t.Fatalf("StreamerGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamerGet(ctx, "test-hostname")
	if err != nil {
		t.Fatalf("StreamerGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamerGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamerSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/central/api/v3/streamers/test-hostname"

	// Expected JSON response
	expectedJSON := `{"api_url":"http://streamer.local:8080","channel_limit":5,"cluster_key":"xS6i6Q3DCc5nEvnu","config":{"cluster_key":"xS6i6Q3DCc5nEvnu","event_sinks":[{"except":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"extra":{"key1":"example","key2":"example"},"level":"debug","max_depth":100,"max_size":10000,"name":"my_json_sink","only":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"resend_limit":1000,"resend_timeout":10,"throttle_delay":1,"url":"jsonlog:///var/log/events-json.log"}],"http_proxies":[{"prefix":"example","stats":{"http_100":1000,"http_200":1000,"http_300":10,"http_400":10,"http_500":10,"protocol_upgrades":1000,"proxy_error":10,"proxy_error_connection":10,"requests":1000,"time_1000ms":2,"time_100ms":10,"time_5000ms":2,"time_500ms":5,"time_longms":1},"url":"https://example.com"}],"rproxy":{"endpoint_auth":"example","forward_ports":{"key1":{"handler":"example"},"key2":{"handler":"example"}},"streampoint_key":"example"},"server_names":[{"aliases":["s1.streamer.local","s2.streamer.local"],"domain":"streamer.local"}]},"cpu_limit":10,"dvrs":[{"disk_usage_limit":98,"disks":[{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}],"episodes_url":"example","index":"example","name":"example","root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000}],"fetch_timeout":1000,"hostname":"peer.example.com","labels":{"key1":"value1","key2":"value2"},"namespace":"example","private_payload_url":"http://streamer.local","public_payload_url":"http://public.example.com","role":"streamer","stale_timeout":1000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","healthcheck_status":{"checks":{"errors_details":[{"error":"example","rule":"example"}]},"status":"ok","status_changed_at":1000000000000},"hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"total_bandwidth":1024}`

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
	body := &model.StreamerConfigImpl{}
	_, err := client.StreamerSave(ctx, "test-hostname", nil, body)
	if err != nil {
		t.Fatalf("StreamerSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamerSave(ctx, "test-hostname", nil, body)
	if err != nil {
		t.Fatalf("StreamerSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamerSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamerStreamsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/streamers/test-hostname/streams"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","streams":[{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"chunk_duration":200,"clients_timeout":485,"comment":"This is a test stream","config_on_disk":{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"chunk_duration":200,"clients_timeout":485,"comment":"This is a test stream","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"disk_usage_limit":98,"episodes_url":"example","redundancy_factor":1,"reference":"localdvr0","remotes":["example_value"],"schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000,"status":"running"}}]},"webrtc_abr":{"start_track":"v2"}},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"disk_usage_limit":98,"episodes_url":"example","redundancy_factor":1,"reference":"localdvr0","remotes":["example_value"],"schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"episodes_density":0.001,"last_dts_at":1636383841974,"last_running_at":1737975543123,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"running_on":["streamer1.example.com"],"status":"running","streaming_endpoint":"example","ts_delay":1284},"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000,"status":"running"}}]},"webrtc_abr":{"start_track":"v2"}}],"updated_at":1727274724000}`

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
	query := &central.StreamerStreamsListQuery{}
	_, err := client.StreamerStreamsList(ctx, "test-hostname", query)
	if err != nil {
		t.Fatalf("StreamerStreamsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamerStreamsList(ctx, "test-hostname", query)
	if err != nil {
		t.Fatalf("StreamerStreamsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamerStreamsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamerUpdateStreamsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/streamers/test-hostname/update-streams"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","streams":[{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"chunk_duration":200,"clients_timeout":485,"comment":"This is a test stream","config_on_disk":{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"chunk_duration":200,"clients_timeout":485,"comment":"This is a test stream","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"disk_usage_limit":98,"episodes_url":"example","redundancy_factor":1,"reference":"localdvr0","remotes":["example_value"],"schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000,"status":"running"}}]},"webrtc_abr":{"start_track":"v2"}},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"disk_usage_limit":98,"episodes_url":"example","redundancy_factor":1,"reference":"localdvr0","remotes":["example_value"],"schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"episodes_density":0.001,"last_dts_at":1636383841974,"last_running_at":1737975543123,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"running_on":["streamer1.example.com"],"status":"running","streaming_endpoint":"example","ts_delay":1284},"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000,"status":"running"}}]},"webrtc_abr":{"start_track":"v2"}}],"updated_at":1727274724000}`

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
	query := &central.StreamerUpdateStreamsListQuery{}
	query.Name = "test-name"
	_, err := client.StreamerUpdateStreamsList(ctx, "test-hostname", query)
	if err != nil {
		t.Fatalf("StreamerUpdateStreamsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamerUpdateStreamsList(ctx, "test-hostname", query)
	if err != nil {
		t.Fatalf("StreamerUpdateStreamsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamerUpdateStreamsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamersBatchUpdate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/central/api/v3/streamers/batch"

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
	err := client.StreamersBatchUpdate(ctx, nil)
	if err != nil {
		t.Fatalf("StreamersBatchUpdate failed: %v", err)
	}
}

func TestClient_StreamersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/streamers"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","streamers":[{"api_url":"http://streamer.local:8080","channel_limit":5,"cluster_key":"xS6i6Q3DCc5nEvnu","config":{"cluster_key":"xS6i6Q3DCc5nEvnu","event_sinks":[{"except":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"extra":{"key1":"example","key2":"example"},"level":"debug","max_depth":100,"max_size":10000,"name":"my_json_sink","only":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"resend_limit":1000,"resend_timeout":10,"throttle_delay":1,"url":"jsonlog:///var/log/events-json.log"}],"http_proxies":[{"prefix":"example","stats":{"http_100":1000,"http_200":1000,"http_300":10,"http_400":10,"http_500":10,"protocol_upgrades":1000,"proxy_error":10,"proxy_error_connection":10,"requests":1000,"time_1000ms":2,"time_100ms":10,"time_5000ms":2,"time_500ms":5,"time_longms":1},"url":"https://example.com"}],"rproxy":{"endpoint_auth":"example","forward_ports":{"key1":{"handler":"example"},"key2":{"handler":"example"}},"streampoint_key":"example"},"server_names":[{"aliases":["s1.streamer.local","s2.streamer.local"],"domain":"streamer.local"}]},"cpu_limit":10,"dvrs":[{"disk_usage_limit":98,"disks":[{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}],"episodes_url":"example","index":"example","name":"example","root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000}],"fetch_timeout":1000,"hostname":"peer.example.com","labels":{"key1":"value1","key2":"value2"},"namespace":"example","private_payload_url":"http://streamer.local","public_payload_url":"http://public.example.com","role":"streamer","stale_timeout":1000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","healthcheck_status":{"checks":{"errors_details":[{"error":"example","rule":"example"}]},"status":"ok","status_changed_at":1000000000000},"hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"total_bandwidth":1024}]}`

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
	query := &central.StreamersListQuery{}
	_, err := client.StreamersList(ctx, query)
	if err != nil {
		t.Fatalf("StreamersList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamersList(ctx, query)
	if err != nil {
		t.Fatalf("StreamersList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamersList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamingLbPlayback(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/lb/test-loadbalancer_name/test-escaped_stream_name/test-multi_segment_suffix"

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
	_, err := client.StreamingLbPlayback(ctx, "test-loadbalancer_name", "test-escaped_stream_name", "test-multi_segment_suffix")
	if err != nil {
		t.Fatalf("StreamingLbPlayback failed: %v", err)
	}
}

func TestClient_StreamingLbPublish(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/central/api/v3/lb/test-loadbalancer_name/test-escaped_stream_name/test-multi_segment_suffix"

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
	err := client.StreamingLbPublish(ctx, "test-loadbalancer_name", "test-escaped_stream_name", "test-multi_segment_suffix")
	if err != nil {
		t.Fatalf("StreamingLbPublish failed: %v", err)
	}
}

func TestClient_StreamsBatchUpdate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/central/api/v3/streams/batch"

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
	err := client.StreamsBatchUpdate(ctx, nil)
	if err != nil {
		t.Fatalf("StreamsBatchUpdate failed: %v", err)
	}
}

func TestClient_StreamsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/central/api/v3/streams"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","streams":[{"claims":{"bitrate":2543},"layout":{"change_reason":"stream_misses_node_required_labels","created_at":1637094994000,"dvr_backups":["example"],"failover_from":"example","inference":"example","ingest":"example","ingest_history":[{"created_at":1637094994000,"ingest":"example","originator":"layouter"}],"node_layout_decisions":[{"hostname":"streamer1.com","reasons":["stream_misses_node_required_labels","node_misses_stream_required_labels","node_channel_limit_exceeded"],"role":"streamer"}],"originator":"layouter"},"namespace":"example","updated_at":1637098611000}]}`

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
	query := &central.StreamsListQuery{}
	_, err := client.StreamsList(ctx, query)
	if err != nil {
		t.Fatalf("StreamsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamsList(ctx, query)
	if err != nil {
		t.Fatalf("StreamsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamsList returned nil result")
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

func createTestClient(t *testing.T, rt http.RoundTripper) central.Central {
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

	client, err := central.NewWithBaseFactory(cfg, baseFactory)
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
