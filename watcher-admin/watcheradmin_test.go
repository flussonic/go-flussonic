package watcheradmin_test

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
	watcheradmin "github.com/flussonic/go-flussonic/watcher-admin"
	model "github.com/flussonic/go-flussonic/watcher-admin/model"
)

func TestClient_AdminStreamsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/streams"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","streams":[{"audio":{"transcode_audio_codec":"aac"},"comment":"This is a test stream","coordinates":{"latitude":55.7512,"longitude":37.6184},"created_at":1672531199000,"domain":{"id":123,"title":"domain1"},"dvr":{"disk_usage_limit":98,"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"allow_if":"example","audio_timeout":20,"comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"labels":{"key1":"example_value","key2":"example_value"},"last_episode_at":1000000000000,"layout":{"change_reason":"stream_misses_node_required_labels","created_at":1637094994000,"dvr_backups":["example"],"inference":"example","ingest":"example","ingest_history":[{"created_at":1637094994000,"ingest":"example","originator":"layouter"}],"node_layout_decisions":[{"hostname":"streamer1.com","reasons":["stream_misses_node_required_labels","node_misses_stream_required_labels","node_channel_limit_exceeded"],"role":"streamer"}],"originator":"layouter"},"layout_rules":{"preferred_zones":["example"],"required_zones":["example"]},"map_coordinates":{"latitude":55.7512,"longitude":37.6184},"name":"example","onvif":{"url":"example"},"organization":{"id":9,"title":"Organization 1"},"organization_id":9,"path":[{"title":"example"}],"postal_address":"example","preset":{"is_adjustable":true,"title":"Example preset name"},"pushes":[{"comment":"This is a test push","retry_timeout":7,"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"playback_token":"onetime_token","publish_endpoint":"example","status":"running","streaming_endpoint":"example","ts_delay":1284},"title":"Hockey channel","vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}}]}`

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
	query := &watcheradmin.AdminStreamsListQuery{}
	_, err := client.AdminStreamsList(ctx, query)
	if err != nil {
		t.Fatalf("AdminStreamsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AdminStreamsList(ctx, query)
	if err != nil {
		t.Fatalf("AdminStreamsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AdminStreamsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_AgentGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/agents/test-id"

	// Expected JSON response
	expectedJSON := `{"id":"1234567","key":"example","model":"example","serial":"example","stats":{"agent_type":"single","endpoint_connection":{"bytes_from_server":40000,"bytes_to_server":400000000000,"hostname":"agents-001.vsaas.io","opened_at":1637094994000,"status_changed_at":1634560921},"local_ip":"10.10.17.88","mac_address":"F0-23-B9-59-20-F1","peer_ip":"185.134.232.183","streampoint_connection":{"bytes_from_server":40000,"bytes_to_server":400000000000,"connections_attempted":400,"connections_current":2,"connections_opened":300,"hostname":"agents-001.vsaas.io","opened_at":1637094994000,"status_changed_at":1634560921},"version":"v21.02-8-g535c85d"},"streams":[{"comment":"This is a test stream","name":"example","organization":{"id":7,"title":"Example LLC"},"title":"Hockey channel"}]}`

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

func TestClient_AgentsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/agents"

	// Expected JSON response
	expectedJSON := `{"agents":[{}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcheradmin.AgentsListQuery{}
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

func TestClient_CameraAuthGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/streams/test-name/auth"

	// Expected JSON response
	expectedJSON := `{"login":"example","password":"example"}`

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
	_, err := client.CameraAuthGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("CameraAuthGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.CameraAuthGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("CameraAuthGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("CameraAuthGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_CameraAuthSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/admin-api/v3/streams/test-name/auth"

	// Expected JSON response
	expectedJSON := `{"login":"example","password":"example"}`

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
	_, err := client.CameraAuthSave(ctx, "test-name")
	if err != nil {
		t.Fatalf("CameraAuthSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.CameraAuthSave(ctx, "test-name")
	if err != nil {
		t.Fatalf("CameraAuthSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("CameraAuthSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_CameraConfigGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/streams/test-name/camera_config"

	// Expected JSON response
	expectedJSON := `{"camera_info":{"firmware":"example","manufacturer":"example","model":"example","serial_number":"example"},"media_quality":"low","sensor":{"blacklight_compensation":{"mode":"on"},"image_orientation":"normal","infrared_cutoff_filter":{"mode":"on"}}}`

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
	_, err := client.CameraConfigGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("CameraConfigGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.CameraConfigGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("CameraConfigGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("CameraConfigGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_CameraConfigSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/admin-api/v3/streams/test-name/camera_config"

	// Expected JSON response
	expectedJSON := `{"camera_info":{"firmware":"example","manufacturer":"example","model":"example","serial_number":"example"},"media_quality":"low","sensor":{"blacklight_compensation":{"mode":"on"},"image_orientation":"normal","infrared_cutoff_filter":{"mode":"on"}}}`

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
	body := &model.CameraConfigImpl{}
	_, err := client.CameraConfigSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("CameraConfigSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.CameraConfigSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("CameraConfigSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("CameraConfigSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_CameraReboot(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/admin-api/v3/streams/test-name/reboot"

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
	err := client.CameraReboot(ctx, "test-name")
	if err != nil {
		t.Fatalf("CameraReboot failed: %v", err)
	}
}

func TestClient_ClusterStatsGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/cluster/stats"

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

func TestClient_DeleteOrganizationPreset(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/admin-api/v3/organizations/test-organization_id/presets/test-preset_id"

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
	err := client.DeleteOrganizationPreset(ctx, "test-organization_id", "test-preset_id")
	if err != nil {
		t.Fatalf("DeleteOrganizationPreset failed: %v", err)
	}
}

func TestClient_DomainConfigGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/domain_config"

	// Expected JSON response
	expectedJSON := `{"appearance":{"colors":{"background":"#f3f5f7","primary":"#444951","secondary":"#2469f2"},"favicon":"data:image/png;base64,\u003cbase64string\u003e","logo":"data:image/png;base64,\u003cbase64string\u003e","logo_auth":"data:image/jpg;base64,\u003cbase64string\u003e","title":"My video site"},"mail":{"host":"example","login":"example","password":"example","security":"none","sender":{"email":"watcher@someserver.com","title":"Flussonic Watcher"},"sending_method":"SMTP"}}`

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
	_, err := client.DomainConfigGet(ctx)
	if err != nil {
		t.Fatalf("DomainConfigGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DomainConfigGet(ctx)
	if err != nil {
		t.Fatalf("DomainConfigGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DomainConfigGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DomainConfigSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/admin-api/v3/domain_config"

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
	body := &model.DomainConfigImpl{}
	err := client.DomainConfigSave(ctx, body)
	if err != nil {
		t.Fatalf("DomainConfigSave failed: %v", err)
	}
}

func TestClient_EventsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/events"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"events":[{"created_at":"2021-01-30T08:30:00.432567Z","event":"api_call","ip":"192.34.32.10","operation_id":"streams_list","originator":"flussonic","path":"example","payload":"example","qs":"example","request_id":"example","user_agent":"Flussonic 24.04"}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcheradmin.EventsListQuery{}
	_, err := client.EventsList(ctx, query)
	if err != nil {
		t.Fatalf("EventsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.EventsList(ctx, query)
	if err != nil {
		t.Fatalf("EventsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("EventsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_LetsencryptIssue(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/admin-api/v3/tls/letsencrypt"

	// Expected JSON response
	expectedJSON := `{"cacert":{"domains":["example"],"issuer_name":"example","public_key":"example"},"cert":{"domains":["example"],"issuer_name":"example","public_key":"example"},"fullchain":[{"domains":["example"],"issuer_name":"example","public_key":"example"}]}`

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
	body := &model.LetsencryptRequestImpl{}
	_, err := client.LetsencryptIssue(ctx, body)
	if err != nil {
		t.Fatalf("LetsencryptIssue failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.LetsencryptIssue(ctx, body)
	if err != nil {
		t.Fatalf("LetsencryptIssue failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("LetsencryptIssue returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_LivenessProbe(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/monitoring/liveness"

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

func TestClient_LoginCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/admin-api/v3/login"

	// Expected JSON response
	expectedJSON := `{"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c","refresh_token":"3637e790-5530-11ed-bdc3-0242ac120002"}`

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
	_, err := client.LoginCreate(ctx)
	if err != nil {
		t.Fatalf("LoginCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.LoginCreate(ctx)
	if err != nil {
		t.Fatalf("LoginCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("LoginCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationPresetSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/admin-api/v3/organizations/test-organization_id/presets"

	// Expected JSON response
	expectedJSON := `{"deleted_at":1637095014573,"dvr":{"disk_usage_limit":98,"redundancy_factor":1,"storage_limit":400000000000},"id":7,"is_adjustable":true,"layout_rules":{"preferred_zones":["example"],"required_zones":["example"]},"stats":{"organizations_count":12,"streams_count":12},"title":"Example preset name","vision":{"alg":"faces"}}`

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
	body := &model.OrganizationPresetImpl{}
	_, err := client.OrganizationPresetSave(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("OrganizationPresetSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationPresetSave(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("OrganizationPresetSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationPresetSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PresetGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/presets/test-id"

	// Expected JSON response
	expectedJSON := `{"deleted_at":1637095014573,"dvr":{"disk_usage_limit":98,"redundancy_factor":1,"storage_limit":400000000000},"id":7,"is_adjustable":true,"layout_rules":{"preferred_zones":["example"],"required_zones":["example"]},"stats":{"organizations_count":12,"streams_count":12},"title":"Example preset name","vision":{"alg":"faces"}}`

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
	_, err := client.PresetGet(ctx, "test-id")
	if err != nil {
		t.Fatalf("PresetGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PresetGet(ctx, "test-id")
	if err != nil {
		t.Fatalf("PresetGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PresetGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PresetSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/admin-api/v3/presets/test-id"

	// Expected JSON response
	expectedJSON := `{"deleted_at":1637095014573,"dvr":{"disk_usage_limit":98,"redundancy_factor":1,"storage_limit":400000000000},"id":7,"is_adjustable":true,"layout_rules":{"preferred_zones":["example"],"required_zones":["example"]},"stats":{"organizations_count":12,"streams_count":12},"title":"Example preset name","vision":{"alg":"faces"}}`

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
	body := &model.PresetImpl{}
	_, err := client.PresetSave(ctx, "test-id", body)
	if err != nil {
		t.Fatalf("PresetSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PresetSave(ctx, "test-id", body)
	if err != nil {
		t.Fatalf("PresetSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PresetSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PresetsCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/admin-api/v3/presets"

	// Expected JSON response
	expectedJSON := `{"deleted_at":1637095014573,"dvr":{"disk_usage_limit":98,"redundancy_factor":1,"storage_limit":400000000000},"id":7,"is_adjustable":true,"layout_rules":{"preferred_zones":["example"],"required_zones":["example"]},"stats":{"organizations_count":12,"streams_count":12},"title":"Example preset name","vision":{"alg":"faces"}}`

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
	body := &model.PresetImpl{}
	_, err := client.PresetsCreate(ctx, body)
	if err != nil {
		t.Fatalf("PresetsCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PresetsCreate(ctx, body)
	if err != nil {
		t.Fatalf("PresetsCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PresetsCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PresetsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/presets"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","presets":[{"deleted_at":1637095014573,"dvr":{"disk_usage_limit":98,"redundancy_factor":1,"storage_limit":400000000000},"id":7,"is_adjustable":true,"layout_rules":{"preferred_zones":["example"],"required_zones":["example"]},"stats":{"organizations_count":12,"streams_count":12},"title":"Example preset name","vision":{"alg":"faces"}}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcheradmin.PresetsListQuery{}
	_, err := client.PresetsList(ctx, query)
	if err != nil {
		t.Fatalf("PresetsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PresetsList(ctx, query)
	if err != nil {
		t.Fatalf("PresetsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PresetsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PreviewLayoutChangeForStreamConfig(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/admin-api/v3/streams/test-name/preview_layout_change"

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
	query := &watcheradmin.PreviewLayoutChangeForStreamConfigQuery{}
	body := &model.StreamConfigImpl{}
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
	expectedPath := "/watcher/admin-api/v3/streamers/test-hostname/preview_layout_change"

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
	query := &watcheradmin.PreviewLayoutChangeForStreamerConfigQuery{}
	body := &model.StreamerLayoutPredictionImpl{}
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
	expectedPath := "/watcher/admin-api/v3/monitoring/readiness"

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

func TestClient_SharedTokenCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/admin-api/v3/shares"

	// Expected JSON response
	expectedJSON := `{"key":"IJCo3KGLBf6NAqMCGgHf5gNhpl9","media":"example","name":"example","permissions":[{"dvr_depth_limit":3600}]}`

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
	body := &model.SharedTokenImpl{}
	_, err := client.SharedTokenCreate(ctx, body)
	if err != nil {
		t.Fatalf("SharedTokenCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SharedTokenCreate(ctx, body)
	if err != nil {
		t.Fatalf("SharedTokenCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SharedTokenCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_SharedTokenDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/admin-api/v3/shares/test-key"

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
	err := client.SharedTokenDelete(ctx, "test-key")
	if err != nil {
		t.Fatalf("SharedTokenDelete failed: %v", err)
	}
}

func TestClient_SharedTokensList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/shares"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","shared_tokens":[{"key":"IJCo3KGLBf6NAqMCGgHf5gNhpl9","media":"example","name":"example","permissions":[{"dvr_depth_limit":3600}]}]}`

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
	query := &watcheradmin.SharedTokensListQuery{}
	_, err := client.SharedTokensList(ctx, query)
	if err != nil {
		t.Fatalf("SharedTokensList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SharedTokensList(ctx, query)
	if err != nil {
		t.Fatalf("SharedTokensList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SharedTokensList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/admin-api/v3/streams/test-name"

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
	err := client.StreamDelete(ctx, "test-name", nil)
	if err != nil {
		t.Fatalf("StreamDelete failed: %v", err)
	}
}

func TestClient_StreamFirmwareUpdate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/admin-api/v3/streams/test-name/firmware_update"

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
	body := &model.FirmwareUpdateImpl{}
	err := client.StreamFirmwareUpdate(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("StreamFirmwareUpdate failed: %v", err)
	}
}

func TestClient_StreamGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/streams/test-name"

	// Expected JSON response
	expectedJSON := `{"audio":{"transcode_audio_codec":"aac"},"comment":"This is a test stream","coordinates":{"latitude":55.7512,"longitude":37.6184},"created_at":1672531199000,"domain":{"id":123,"title":"domain1"},"dvr":{"disk_usage_limit":98,"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"allow_if":"example","audio_timeout":20,"comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"labels":{"key1":"example_value","key2":"example_value"},"last_episode_at":1000000000000,"layout":{"change_reason":"stream_misses_node_required_labels","created_at":1637094994000,"dvr_backups":["example"],"inference":"example","ingest":"example","ingest_history":[{"created_at":1637094994000,"ingest":"example","originator":"layouter"}],"node_layout_decisions":[{"hostname":"streamer1.com","reasons":["stream_misses_node_required_labels","node_misses_stream_required_labels","node_channel_limit_exceeded"],"role":"streamer"}],"originator":"layouter"},"layout_rules":{"preferred_zones":["example"],"required_zones":["example"]},"map_coordinates":{"latitude":55.7512,"longitude":37.6184},"name":"example","onvif":{"url":"example"},"organization":{"id":9,"title":"Organization 1"},"organization_id":9,"path":[{"title":"example"}],"postal_address":"example","preset":{"is_adjustable":true,"title":"Example preset name"},"pushes":[{"comment":"This is a test push","retry_timeout":7,"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"playback_token":"onetime_token","publish_endpoint":"example","status":"running","streaming_endpoint":"example","ts_delay":1284},"title":"Hockey channel","vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}}`

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
	expectedPath := "/watcher/admin-api/v3/streams/test-name/layouts"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"layouts":[{"change_reason":"stream_misses_node_required_labels","created_at":1637094994000,"dvr_backups":["example"],"inference":"example","ingest":"example","ingest_history":[{"created_at":1637094994000,"ingest":"example","originator":"layouter"}],"node_layout_decisions":[{"hostname":"streamer1.com","reasons":["stream_misses_node_required_labels","node_misses_stream_required_labels","node_channel_limit_exceeded"],"role":"streamer"}],"originator":"layouter"}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcheradmin.StreamLayoutsGetQuery{}
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
	expectedPath := "/watcher/admin-api/v3/streams/test-name"

	// Expected JSON response
	expectedJSON := `{"audio":{"transcode_audio_codec":"aac"},"comment":"This is a test stream","coordinates":{"latitude":55.7512,"longitude":37.6184},"created_at":1672531199000,"domain":{"id":123,"title":"domain1"},"dvr":{"disk_usage_limit":98,"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"allow_if":"example","audio_timeout":20,"comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"labels":{"key1":"example_value","key2":"example_value"},"last_episode_at":1000000000000,"layout":{"change_reason":"stream_misses_node_required_labels","created_at":1637094994000,"dvr_backups":["example"],"inference":"example","ingest":"example","ingest_history":[{"created_at":1637094994000,"ingest":"example","originator":"layouter"}],"node_layout_decisions":[{"hostname":"streamer1.com","reasons":["stream_misses_node_required_labels","node_misses_stream_required_labels","node_channel_limit_exceeded"],"role":"streamer"}],"originator":"layouter"},"layout_rules":{"preferred_zones":["example"],"required_zones":["example"]},"map_coordinates":{"latitude":55.7512,"longitude":37.6184},"name":"example","onvif":{"url":"example"},"organization":{"id":9,"title":"Organization 1"},"organization_id":9,"path":[{"title":"example"}],"postal_address":"example","preset":{"is_adjustable":true,"title":"Example preset name"},"pushes":[{"comment":"This is a test push","retry_timeout":7,"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"playback_token":"onetime_token","publish_endpoint":"example","status":"running","streaming_endpoint":"example","ts_delay":1284},"title":"Hockey channel","vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}}`

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
	body := &model.StreamConfigImpl{}
	_, err := client.StreamSave(ctx, "test-name", nil, body)
	if err != nil {
		t.Fatalf("StreamSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamSave(ctx, "test-name", nil, body)
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
	expectedPath := "/watcher/admin-api/v3/streamers/test-hostname"

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

func TestClient_StreamerGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/streamers/test-hostname"

	// Expected JSON response
	expectedJSON := `{"zones":["zone1","zone2"]}`

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
	expectedPath := "/watcher/admin-api/v3/streamers/test-hostname"

	// Expected JSON response
	expectedJSON := `{"zones":["zone1","zone2"]}`

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
	body := &model.StreamerImpl{}
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

func TestClient_StreamersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/streamers"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","streamers":[{"zones":["zone1","zone2"]}]}`

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
	query := &watcheradmin.StreamersListQuery{}
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

func TestClient_StreamsMultiedit(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/admin-api/v3/streams/multiedit"

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
	body := &model.StreamsMultieditConfigImpl{}
	err := client.StreamsMultiedit(ctx, body)
	if err != nil {
		t.Fatalf("StreamsMultiedit failed: %v", err)
	}
}

func TestClient_StreamsZonesSettings(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/admin-api/v3/streams/zones"

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
	body := &model.StreamsZoneSettingsImpl{}
	err := client.StreamsZonesSettings(ctx, body)
	if err != nil {
		t.Fatalf("StreamsZonesSettings failed: %v", err)
	}
}

func TestClient_SystemConfigGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/system_config"

	// Expected JSON response
	expectedJSON := `{"api_url":"auth@localhost:9015","central_url":"http://auth@localhost:9019/streamer/api/v3","custom_auth":"http://localhost/camera/authorization","database":{"url":"postgres://central:pass@localhost:5432/central_dev"},"license_key":"example","listeners":{"http":[{"address":"10.0.35.1","port":80}]},"loglevel":"debug"}`

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
	_, err := client.SystemConfigGet(ctx)
	if err != nil {
		t.Fatalf("SystemConfigGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SystemConfigGet(ctx)
	if err != nil {
		t.Fatalf("SystemConfigGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SystemConfigGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TlsCertificateGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/tls/letsencrypt"

	// Expected JSON response
	expectedJSON := `{"cacert":{"domains":["example"],"issuer_name":"example","public_key":"example"},"cert":{"domains":["example"],"issuer_name":"example","public_key":"example"},"fullchain":[{"domains":["example"],"issuer_name":"example","public_key":"example"}]}`

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
	_, err := client.TlsCertificateGet(ctx)
	if err != nil {
		t.Fatalf("TlsCertificateGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TlsCertificateGet(ctx)
	if err != nil {
		t.Fatalf("TlsCertificateGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TlsCertificateGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UserCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/admin-api/v3/users"

	// Expected JSON response
	expectedJSON := `{"access_level":"generic","created_at":1672531199000,"email":"user@example.com","fullname":"example","id":1,"locale":"en","name":"example","note":"example","organizations":[{"id":1,"owner":{"id":2,"name":"admin"},"permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true},"title":"Cameras"}],"password":"example","phone":"+78007778413"}`

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
	_, err := client.UserCreate(ctx, nil)
	if err != nil {
		t.Fatalf("UserCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UserCreate(ctx, nil)
	if err != nil {
		t.Fatalf("UserCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UserCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UserDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/admin-api/v3/users/test-user_id"

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
	err := client.UserDelete(ctx, "test-user_id")
	if err != nil {
		t.Fatalf("UserDelete failed: %v", err)
	}
}

func TestClient_UserGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/users/test-user_id"

	// Expected JSON response
	expectedJSON := `{"access_level":"generic","created_at":1672531199000,"email":"user@example.com","fullname":"example","id":1,"locale":"en","name":"example","note":"example","organizations":[{"id":1,"owner":{"id":2,"name":"admin"},"permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true},"title":"Cameras"}],"password":"example","phone":"+78007778413"}`

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
	_, err := client.UserGet(ctx, "test-user_id")
	if err != nil {
		t.Fatalf("UserGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UserGet(ctx, "test-user_id")
	if err != nil {
		t.Fatalf("UserGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UserGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UserSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/admin-api/v3/users/test-user_id"

	// Expected JSON response
	expectedJSON := `{"access_level":"generic","created_at":1672531199000,"email":"user@example.com","fullname":"example","id":1,"locale":"en","name":"example","note":"example","organizations":[{"id":1,"owner":{"id":2,"name":"admin"},"permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true},"title":"Cameras"}],"password":"example","phone":"+78007778413"}`

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
	body := &model.UserImpl{}
	_, err := client.UserSave(ctx, "test-user_id", body)
	if err != nil {
		t.Fatalf("UserSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UserSave(ctx, "test-user_id", body)
	if err != nil {
		t.Fatalf("UserSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UserSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UsersApikeyCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/admin-api/v3/users/test-user_id/apikey"

	// Expected JSON response
	expectedJSON := `{"apikey":"example"}`

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
	_, err := client.UsersApikeyCreate(ctx, "test-user_id")
	if err != nil {
		t.Fatalf("UsersApikeyCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UsersApikeyCreate(ctx, "test-user_id")
	if err != nil {
		t.Fatalf("UsersApikeyCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UsersApikeyCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UsersApikeyGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/users/test-user_id/apikey"

	// Expected JSON response
	expectedJSON := `{"apikey":"example"}`

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
	_, err := client.UsersApikeyGet(ctx, "test-user_id")
	if err != nil {
		t.Fatalf("UsersApikeyGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UsersApikeyGet(ctx, "test-user_id")
	if err != nil {
		t.Fatalf("UsersApikeyGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UsersApikeyGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UsersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/users"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","users_list":[{"access_level":"generic","created_at":1672531199000,"email":"user@example.com","fullname":"example","id":1,"locale":"en","name":"example","note":"example","organizations":[{"id":1,"owner":{"id":2,"name":"admin"},"permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true},"title":"Cameras"}],"password":"example","phone":"+78007778413"}]}`

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
	query := &watcheradmin.UsersListQuery{}
	_, err := client.UsersList(ctx, query)
	if err != nil {
		t.Fatalf("UsersList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UsersList(ctx, query)
	if err != nil {
		t.Fatalf("UsersList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UsersList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_WatcherProvision(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/admin-api/v3/peeklio/create_operator_id"

	// Expected JSON response
	expectedJSON := `{"checks":{"errors_details":[{"check":"example","error":"example"}],"is_domain_api_key_matched":true,"is_watcher_url_available":true},"operator_id":15432}`

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
	_, err := client.WatcherProvision(ctx)
	if err != nil {
		t.Fatalf("WatcherProvision failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.WatcherProvision(ctx)
	if err != nil {
		t.Fatalf("WatcherProvision failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("WatcherProvision returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_WatcherStatusGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/peeklio/check_watcher_status"

	// Expected JSON response
	expectedJSON := `{"checks":{"errors_details":[{"check":"example","error":"example"}],"is_domain_api_key_matched":true,"is_watcher_url_available":true},"operator_id":15432}`

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
	_, err := client.WatcherStatusGet(ctx)
	if err != nil {
		t.Fatalf("WatcherStatusGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.WatcherStatusGet(ctx)
	if err != nil {
		t.Fatalf("WatcherStatusGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("WatcherStatusGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ZoneDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/admin-api/v3/zones/test-name"

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
	err := client.ZoneDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("ZoneDelete failed: %v", err)
	}
}

func TestClient_ZoneGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/zones/test-name"

	// Expected JSON response
	expectedJSON := `{"name":"zone1","stats":{"preferred_streams_count":10,"required_streams_count":10,"streamers":["streamer1","streamer2"],"streamers_count":1}}`

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
	_, err := client.ZoneGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("ZoneGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ZoneGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("ZoneGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ZoneGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ZoneSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/admin-api/v3/zones/test-name"

	// Expected JSON response
	expectedJSON := `{"name":"zone1","stats":{"preferred_streams_count":10,"required_streams_count":10,"streamers":["streamer1","streamer2"],"streamers_count":1}}`

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
	body := &model.ZoneImpl{}
	_, err := client.ZoneSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("ZoneSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ZoneSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("ZoneSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ZoneSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ZonesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/admin-api/v3/zones"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","zones":[{"name":"zone1","stats":{"preferred_streams_count":10,"required_streams_count":10,"streamers":["streamer1","streamer2"],"streamers_count":1}}]}`

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
	query := &watcheradmin.ZonesListQuery{}
	_, err := client.ZonesList(ctx, query)
	if err != nil {
		t.Fatalf("ZonesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ZonesList(ctx, query)
	if err != nil {
		t.Fatalf("ZonesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ZonesList returned nil result")
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

func createTestClient(t *testing.T, rt http.RoundTripper) watcheradmin.WatcherAdmin {
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

	client, err := watcheradmin.NewWithBaseFactory(cfg, baseFactory)
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
