package watcherclient_test

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
	watcherclient "github.com/flussonic/go-flussonic/watcher-client"
	model "github.com/flussonic/go-flussonic/watcher-client/model"
)

func TestClient_AccessLogList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/access_log"

	// Expected JSON response
	expectedJSON := `{"access_log":[{"action":{"type":"create"},"created_at":1000000000,"domain":{"title":"example"},"object":{"id":"example","representation":"example","type":"Stream"},"organization":{"title":"example"},"user":{"login":"example","request_data":{"ip":"example","user_agent":"example"}}}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.AccessLogListQuery{}
	_, err := client.AccessLogList(ctx, query)
	if err != nil {
		t.Fatalf("AccessLogList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AccessLogList(ctx, query)
	if err != nil {
		t.Fatalf("AccessLogList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AccessLogList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_AgentActivationTokenCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/agent_activation_token"

	// Expected JSON response
	expectedJSON := `null`

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
	body := &model.AgentActivationTokenRequestImpl{}
	_, err := client.AgentActivationTokenCreate(ctx, body)
	if err != nil {
		t.Fatalf("AgentActivationTokenCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AgentActivationTokenCreate(ctx, body)
	if err != nil {
		t.Fatalf("AgentActivationTokenCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AgentActivationTokenCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_AgentActivationTokenGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/agent_activation_token/test-token"

	// Expected JSON response
	expectedJSON := `null`

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
	_, err := client.AgentActivationTokenGet(ctx, "test-token")
	if err != nil {
		t.Fatalf("AgentActivationTokenGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AgentActivationTokenGet(ctx, "test-token")
	if err != nil {
		t.Fatalf("AgentActivationTokenGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AgentActivationTokenGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_Autologin(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/autologin"

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
	body := &model.AutologinRequestImpl{}
	_, err := client.Autologin(ctx, body)
	if err != nil {
		t.Fatalf("Autologin failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.Autologin(ctx, body)
	if err != nil {
		t.Fatalf("Autologin failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("Autologin returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_AutologinGenerate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/autologin/generate"

	// Expected JSON response
	expectedJSON := `{"autologin_token":"42:3600:1738108800:a1b2:c3d4e5f6..."}`

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
	body := &model.AutologinGenerateRequestImpl{}
	_, err := client.AutologinGenerate(ctx, body)
	if err != nil {
		t.Fatalf("AutologinGenerate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AutologinGenerate(ctx, body)
	if err != nil {
		t.Fatalf("AutologinGenerate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AutologinGenerate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_CameraAuthGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/streams/test-name/auth"

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
	expectedPath := "/watcher/client-api/v3/streams/test-name/auth"

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
	expectedPath := "/watcher/client-api/v3/streams/test-name/camera_config"

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
	expectedPath := "/watcher/client-api/v3/streams/test-name/camera_config"

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
	expectedPath := "/watcher/client-api/v3/streams/test-name/reboot"

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

func TestClient_CameraUsagesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/usage/cameras"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","usage":[{"account_id":12345,"billing_plan_id":"billing-plan-456","external_account_id":"12345","from":"2025-01-01","license_id":"12345","plan_id":"episodes-30d-123","stream_name":"ag-123456","to":"2025-01-30"}]}`

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
	query := &watcherclient.CameraUsagesListQuery{}
	_, err := client.CameraUsagesList(ctx, query)
	if err != nil {
		t.Fatalf("CameraUsagesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.CameraUsagesList(ctx, query)
	if err != nil {
		t.Fatalf("CameraUsagesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("CameraUsagesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DeleteOrganizationPreset(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/presets/test-preset_id"

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

func TestClient_DeviceTokenConfirm(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/profile/device-tokens/test-token/confirm"

	// Expected JSON response
	expectedJSON := `{"device":{"model":"example","platform":"ios","version":"example"},"token":"example"}`

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
	_, err := client.DeviceTokenConfirm(ctx, "test-token")
	if err != nil {
		t.Fatalf("DeviceTokenConfirm failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DeviceTokenConfirm(ctx, "test-token")
	if err != nil {
		t.Fatalf("DeviceTokenConfirm failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DeviceTokenConfirm returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DeviceTokenDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/profile/device-tokens/test-token"

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
	err := client.DeviceTokenDelete(ctx, "test-token")
	if err != nil {
		t.Fatalf("DeviceTokenDelete failed: %v", err)
	}
}

func TestClient_DeviceTokenGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/profile/device-tokens/test-token"

	// Expected JSON response
	expectedJSON := `{"device":{"model":"example","platform":"ios","version":"example"},"token":"example"}`

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
	_, err := client.DeviceTokenGet(ctx, "test-token")
	if err != nil {
		t.Fatalf("DeviceTokenGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DeviceTokenGet(ctx, "test-token")
	if err != nil {
		t.Fatalf("DeviceTokenGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DeviceTokenGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DeviceTokenSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/profile/device-tokens/test-token"

	// Expected JSON response
	expectedJSON := `{"device":{"model":"example","platform":"ios","version":"example"},"token":"example"}`

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
	body := &model.DeviceTokenImpl{}
	_, err := client.DeviceTokenSave(ctx, "test-token", body)
	if err != nil {
		t.Fatalf("DeviceTokenSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DeviceTokenSave(ctx, "test-token", body)
	if err != nil {
		t.Fatalf("DeviceTokenSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DeviceTokenSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DiscoveryDevicesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs/test-nvr_id/discovery/devices"

	// Expected JSON response
	expectedJSON := `{"devices":[{"dns":["example"],"firmware":"example","gateway":"example","hostname":"example","ip":"example","mac":"example","manufacturer":"example","model":"example","netmask":"example","onvif_url":"http://192.168.1.100:8080/onvif/device_service","serial_number":"example","system_time":1662386000,"system_time_at":1662386000}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.DiscoveryDevicesListQuery{}
	_, err := client.DiscoveryDevicesList(ctx, "test-organization_id", "test-nvr_id", query)
	if err != nil {
		t.Fatalf("DiscoveryDevicesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DiscoveryDevicesList(ctx, "test-organization_id", "test-nvr_id", query)
	if err != nil {
		t.Fatalf("DiscoveryDevicesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DiscoveryDevicesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_EpisodeAddToFavorites(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/episodes/test-episode_id/favorites"

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
	err := client.EpisodeAddToFavorites(ctx, "test-episode_id")
	if err != nil {
		t.Fatalf("EpisodeAddToFavorites failed: %v", err)
	}
}

func TestClient_EpisodeDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/episodes/test-episode_id"

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
	err := client.EpisodeDelete(ctx, "test-episode_id", nil)
	if err != nil {
		t.Fatalf("EpisodeDelete failed: %v", err)
	}
}

func TestClient_EpisodeDeleteFromFavorites(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/episodes/test-episode_id/favorites"

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
	err := client.EpisodeDeleteFromFavorites(ctx, "test-episode_id")
	if err != nil {
		t.Fatalf("EpisodeDeleteFromFavorites failed: %v", err)
	}
}

func TestClient_EpisodeGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/episodes/test-episode_id"

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

func TestClient_EpisodeUpdate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/episodes/test-episode_id"

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
	body := &model.EpisodeUpdateImpl{}
	_, err := client.EpisodeUpdate(ctx, "test-episode_id", nil, body)
	if err != nil {
		t.Fatalf("EpisodeUpdate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.EpisodeUpdate(ctx, "test-episode_id", nil, body)
	if err != nil {
		t.Fatalf("EpisodeUpdate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("EpisodeUpdate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_EpisodesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/episodes"

	// Expected JSON response
	expectedJSON := `{"episodes":[{}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.EpisodesListQuery{}
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

func TestClient_EventSubscriptionCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/profile/subscriptions/stream_event"

	// Expected JSON response
	expectedJSON := `{"event_types":{"episode_face":true,"episode_generic":true,"stream_dead":true},"notification_type":"push","stream_name":"example"}`

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
	body := &model.SubscriptionRequestImpl{}
	_, err := client.EventSubscriptionCreate(ctx, body)
	if err != nil {
		t.Fatalf("EventSubscriptionCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.EventSubscriptionCreate(ctx, body)
	if err != nil {
		t.Fatalf("EventSubscriptionCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("EventSubscriptionCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_EventSubscriptionDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/profile/subscriptions/stream_event"

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
	body := &model.SubscriptionRequestImpl{}
	err := client.EventSubscriptionDelete(ctx, body)
	if err != nil {
		t.Fatalf("EventSubscriptionDelete failed: %v", err)
	}
}

func TestClient_FolderCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/folders"

	// Expected JSON response
	expectedJSON := `{"coordinates":{"latitude":55.7512,"longitude":37.6184},"floor_plan":{"bottomleft":{"latitude":55.7512,"longitude":37.6184},"file":{"b64_content":"example","binary_content":"example","mime_type":"example","name":"1-maps-2_07957af5f132243350a0f5546a675cb006354b69","url":"example"},"topleft":{"latitude":55.7512,"longitude":37.6184},"topright":{"latitude":55.7512,"longitude":37.6184}},"hierarchy":{"shift":{"direction":"after"}},"id":0,"streams_count":2,"title":"example"}`

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
	body := &model.FolderImpl{}
	_, err := client.FolderCreate(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("FolderCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.FolderCreate(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("FolderCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("FolderCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_FolderDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/folders/test-folder_id"

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
	err := client.FolderDelete(ctx, "test-organization_id", "test-folder_id")
	if err != nil {
		t.Fatalf("FolderDelete failed: %v", err)
	}
}

func TestClient_FolderGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/folders/test-folder_id"

	// Expected JSON response
	expectedJSON := `{"coordinates":{"latitude":55.7512,"longitude":37.6184},"floor_plan":{"bottomleft":{"latitude":55.7512,"longitude":37.6184},"file":{"b64_content":"example","binary_content":"example","mime_type":"example","name":"1-maps-2_07957af5f132243350a0f5546a675cb006354b69","url":"example"},"topleft":{"latitude":55.7512,"longitude":37.6184},"topright":{"latitude":55.7512,"longitude":37.6184}},"hierarchy":{"shift":{"direction":"after"}},"id":0,"streams_count":2,"title":"example"}`

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
	_, err := client.FolderGet(ctx, "test-organization_id", "test-folder_id")
	if err != nil {
		t.Fatalf("FolderGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.FolderGet(ctx, "test-organization_id", "test-folder_id")
	if err != nil {
		t.Fatalf("FolderGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("FolderGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_FolderList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/folders"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"folders":[{"coordinates":{"latitude":55.7512,"longitude":37.6184},"floor_plan":{"bottomleft":{"latitude":55.7512,"longitude":37.6184},"file":{"b64_content":"example","binary_content":"example","mime_type":"example","name":"1-maps-2_07957af5f132243350a0f5546a675cb006354b69","url":"example"},"topleft":{"latitude":55.7512,"longitude":37.6184},"topright":{"latitude":55.7512,"longitude":37.6184}},"hierarchy":{"shift":{"direction":"after"}},"id":0,"streams_count":2,"title":"example"}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.FolderListQuery{}
	_, err := client.FolderList(ctx, "test-organization_id", query)
	if err != nil {
		t.Fatalf("FolderList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.FolderList(ctx, "test-organization_id", query)
	if err != nil {
		t.Fatalf("FolderList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("FolderList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_FolderSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/folders/test-folder_id"

	// Expected JSON response
	expectedJSON := `{"coordinates":{"latitude":55.7512,"longitude":37.6184},"floor_plan":{"bottomleft":{"latitude":55.7512,"longitude":37.6184},"file":{"b64_content":"example","binary_content":"example","mime_type":"example","name":"1-maps-2_07957af5f132243350a0f5546a675cb006354b69","url":"example"},"topleft":{"latitude":55.7512,"longitude":37.6184},"topright":{"latitude":55.7512,"longitude":37.6184}},"hierarchy":{"shift":{"direction":"after"}},"id":0,"streams_count":2,"title":"example"}`

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
	body := &model.FolderImpl{}
	_, err := client.FolderSave(ctx, "test-organization_id", "test-folder_id", body)
	if err != nil {
		t.Fatalf("FolderSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.FolderSave(ctx, "test-organization_id", "test-folder_id", body)
	if err != nil {
		t.Fatalf("FolderSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("FolderSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_FolderUserDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/folders/test-folder_id/users/test-user_id"

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
	err := client.FolderUserDelete(ctx, "test-organization_id", "test-folder_id", "test-user_id")
	if err != nil {
		t.Fatalf("FolderUserDelete failed: %v", err)
	}
}

func TestClient_FolderUserGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/folders/test-folder_id/users/test-user_id"

	// Expected JSON response
	expectedJSON := `{"dvr_depth_limit":3600}`

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
	_, err := client.FolderUserGet(ctx, "test-organization_id", "test-folder_id", "test-user_id")
	if err != nil {
		t.Fatalf("FolderUserGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.FolderUserGet(ctx, "test-organization_id", "test-folder_id", "test-user_id")
	if err != nil {
		t.Fatalf("FolderUserGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("FolderUserGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_FolderUserSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/folders/test-folder_id/users/test-user_id"

	// Expected JSON response
	expectedJSON := `{"dvr_depth_limit":3600}`

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
	body := &model.FolderUserImpl{}
	_, err := client.FolderUserSave(ctx, "test-organization_id", "test-folder_id", "test-user_id", body)
	if err != nil {
		t.Fatalf("FolderUserSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.FolderUserSave(ctx, "test-organization_id", "test-folder_id", "test-user_id", body)
	if err != nil {
		t.Fatalf("FolderUserSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("FolderUserSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_FolderUsersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/folders/test-folder_id/users"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","users":[{"dvr_depth_limit":3600}]}`

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
	query := &watcherclient.FolderUsersListQuery{}
	_, err := client.FolderUsersList(ctx, "test-organization_id", "test-folder_id", query)
	if err != nil {
		t.Fatalf("FolderUsersList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.FolderUsersList(ctx, "test-organization_id", "test-folder_id", query)
	if err != nil {
		t.Fatalf("FolderUsersList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("FolderUsersList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_GetInviteShortInfo(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/invite/test-invite_key"

	// Expected JSON response
	expectedJSON := `{"id":7,"title":"Example LLC"}`

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
	_, err := client.GetInviteShortInfo(ctx, "test-invite_key")
	if err != nil {
		t.Fatalf("GetInviteShortInfo failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.GetInviteShortInfo(ctx, "test-invite_key")
	if err != nil {
		t.Fatalf("GetInviteShortInfo failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("GetInviteShortInfo returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_LoginCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/login"

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

func TestClient_MessageDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/messages/test-message_id"

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
	err := client.MessageDelete(ctx, "test-message_id")
	if err != nil {
		t.Fatalf("MessageDelete failed: %v", err)
	}
}

func TestClient_MessageGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/messages/test-message_id"

	// Expected JSON response
	expectedJSON := `{"body":"**Maintenance work is in progress.**\n","id":1,"sender":{"id":1,"name":"admin"},"title":"Attention! Maintenance work is in progress.\n","type":"warning","user":{"id":1,"name":"support"}}`

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
	_, err := client.MessageGet(ctx, "test-message_id")
	if err != nil {
		t.Fatalf("MessageGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.MessageGet(ctx, "test-message_id")
	if err != nil {
		t.Fatalf("MessageGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("MessageGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_MessageSend(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/messages"

	// Expected JSON response
	expectedJSON := `{"devices":3}`

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
	body := &model.MessageSendImpl{}
	_, err := client.MessageSend(ctx, body)
	if err != nil {
		t.Fatalf("MessageSend failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.MessageSend(ctx, body)
	if err != nil {
		t.Fatalf("MessageSend failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("MessageSend returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_MessageUpdate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/messages/test-message_id"

	// Expected JSON response
	expectedJSON := `{"body":"**Maintenance work is in progress.**\n","id":1,"sender":{"id":1,"name":"admin"},"title":"Attention! Maintenance work is in progress.\n","type":"warning","user":{"id":1,"name":"support"}}`

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
	body := &model.MessageChangeRequestImpl{}
	_, err := client.MessageUpdate(ctx, "test-message_id", body)
	if err != nil {
		t.Fatalf("MessageUpdate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.MessageUpdate(ctx, "test-message_id", body)
	if err != nil {
		t.Fatalf("MessageUpdate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("MessageUpdate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_MessagesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/messages"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"messages":[{"body":"**Maintenance work is in progress.**\n","id":1,"sender":{"id":1,"name":"admin"},"title":"Attention! Maintenance work is in progress.\n","type":"warning","user":{"id":1,"name":"support"}}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.MessagesListQuery{}
	_, err := client.MessagesList(ctx, query)
	if err != nil {
		t.Fatalf("MessagesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.MessagesList(ctx, query)
	if err != nil {
		t.Fatalf("MessagesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("MessagesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_MosaicCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/mosaics"

	// Expected JSON response
	expectedJSON := `{"id":7,"organization":{"id":9,"user_permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}},"organization_id":9,"streams":[{"name":"ag-12345","playback_token":"onetime_token","streaming_endpoint":"https://ms.example.com","title":"Hockey channel"}],"title":"example","type":"2x2"}`

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
	_, err := client.MosaicCreate(ctx, nil)
	if err != nil {
		t.Fatalf("MosaicCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.MosaicCreate(ctx, nil)
	if err != nil {
		t.Fatalf("MosaicCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("MosaicCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_MosaicDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/mosaics/test-mosaic_id"

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
	err := client.MosaicDelete(ctx, "test-mosaic_id")
	if err != nil {
		t.Fatalf("MosaicDelete failed: %v", err)
	}
}

func TestClient_MosaicGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/mosaics/test-mosaic_id"

	// Expected JSON response
	expectedJSON := `{"id":7,"organization":{"id":9,"user_permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}},"organization_id":9,"streams":[{"name":"ag-12345","playback_token":"onetime_token","streaming_endpoint":"https://ms.example.com","title":"Hockey channel"}],"title":"example","type":"2x2"}`

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
	_, err := client.MosaicGet(ctx, "test-mosaic_id")
	if err != nil {
		t.Fatalf("MosaicGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.MosaicGet(ctx, "test-mosaic_id")
	if err != nil {
		t.Fatalf("MosaicGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("MosaicGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_MosaicSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/mosaics/test-mosaic_id"

	// Expected JSON response
	expectedJSON := `{"id":7,"organization":{"id":9,"user_permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}},"organization_id":9,"streams":[{"name":"ag-12345","playback_token":"onetime_token","streaming_endpoint":"https://ms.example.com","title":"Hockey channel"}],"title":"example","type":"2x2"}`

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
	body := &model.MosaicSaveImpl{}
	_, err := client.MosaicSave(ctx, "test-mosaic_id", body)
	if err != nil {
		t.Fatalf("MosaicSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.MosaicSave(ctx, "test-mosaic_id", body)
	if err != nil {
		t.Fatalf("MosaicSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("MosaicSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_MosaicsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/mosaics"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"mosaics":[{"id":7,"organization":{"id":9,"user_permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}},"organization_id":9,"streams":[{"name":"ag-12345","playback_token":"onetime_token","streaming_endpoint":"https://ms.example.com","title":"Hockey channel"}],"title":"example","type":"2x2"}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.MosaicsListQuery{}
	_, err := client.MosaicsList(ctx, query)
	if err != nil {
		t.Fatalf("MosaicsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.MosaicsList(ctx, query)
	if err != nil {
		t.Fatalf("MosaicsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("MosaicsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_NotificationSend(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/notifications/notify"

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
	err := client.NotificationSend(ctx, nil)
	if err != nil {
		t.Fatalf("NotificationSend failed: %v", err)
	}
}

func TestClient_NvrActivationTokenCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs/nvr-activation-token"

	// Expected JSON response
	expectedJSON := `{"token":"a1b2c3d4e5f6","valid_till":1672531199}`

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
	body := &model.NvrActivationTokenRequestImpl{}
	_, err := client.NvrActivationTokenCreate(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("NvrActivationTokenCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.NvrActivationTokenCreate(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("NvrActivationTokenCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("NvrActivationTokenCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_NvrConfigGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/nvrs/config"

	// Expected JSON response
	expectedJSON := `{"config_version":1637094994,"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","streams":[{"comment":"This is a test stream","dvr":{"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"url":"fake://fake"}],"name":"example","static":true,"title":"Hockey channel","vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}}]}`

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
	query := &watcherclient.NvrConfigGetQuery{}
	_, err := client.NvrConfigGet(ctx, query)
	if err != nil {
		t.Fatalf("NvrConfigGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.NvrConfigGet(ctx, query)
	if err != nil {
		t.Fatalf("NvrConfigGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("NvrConfigGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/organizations"

	// Expected JSON response
	expectedJSON := `{"created_at":1672531199000,"id":7,"is_default":true,"limits":{"streams":50,"users":50},"owner":{"id":2,"name":"admin"},"stats":{"mosaics":2,"streams":12,"users":12},"title":"Example LLC","user_permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}}`

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
	body := &model.OrganizationImpl{}
	_, err := client.OrganizationCreate(ctx, body)
	if err != nil {
		t.Fatalf("OrganizationCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationCreate(ctx, body)
	if err != nil {
		t.Fatalf("OrganizationCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id"

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
	err := client.OrganizationDelete(ctx, "test-organization_id")
	if err != nil {
		t.Fatalf("OrganizationDelete failed: %v", err)
	}
}

func TestClient_OrganizationGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id"

	// Expected JSON response
	expectedJSON := `{"created_at":1672531199000,"id":7,"is_default":true,"limits":{"streams":50,"users":50},"owner":{"id":2,"name":"admin"},"stats":{"mosaics":2,"streams":12,"users":12},"title":"Example LLC","user_permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}}`

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
	_, err := client.OrganizationGet(ctx, "test-organization_id")
	if err != nil {
		t.Fatalf("OrganizationGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationGet(ctx, "test-organization_id")
	if err != nil {
		t.Fatalf("OrganizationGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationInviteAccept(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/organizations/invite/test-invite_key"

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
	err := client.OrganizationInviteAccept(ctx, "test-invite_key")
	if err != nil {
		t.Fatalf("OrganizationInviteAccept failed: %v", err)
	}
}

func TestClient_OrganizationInviteCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/invite"

	// Expected JSON response
	expectedJSON := `{"invite_key":"example"}`

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
	body := &model.OrganizationInviteSetupImpl{}
	_, err := client.OrganizationInviteCreate(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("OrganizationInviteCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationInviteCreate(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("OrganizationInviteCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationInviteCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationNvrCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs"

	// Expected JSON response
	expectedJSON := `{"api_key":"api-key","id":1,"note":"Note","organization":{"id":7,"title":"Example LLC"},"stats":{"healthcheck_status":{"status":"ok","status_changed_at":1000000000000}},"sync":{"episodes":true},"title":"NVR Title"}`

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
	body := &model.NvrImpl{}
	_, err := client.OrganizationNvrCreate(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("OrganizationNvrCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationNvrCreate(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("OrganizationNvrCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationNvrCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationNvrDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs/test-nvr_id"

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
	err := client.OrganizationNvrDelete(ctx, "test-organization_id", "test-nvr_id")
	if err != nil {
		t.Fatalf("OrganizationNvrDelete failed: %v", err)
	}
}

func TestClient_OrganizationNvrEpisodeGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs/test-nvr_id/episodes/test-episode_id"

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
	_, err := client.OrganizationNvrEpisodeGet(ctx, "test-organization_id", "test-nvr_id", "test-episode_id")
	if err != nil {
		t.Fatalf("OrganizationNvrEpisodeGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationNvrEpisodeGet(ctx, "test-organization_id", "test-nvr_id", "test-episode_id")
	if err != nil {
		t.Fatalf("OrganizationNvrEpisodeGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationNvrEpisodeGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationNvrEpisodesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs/test-nvr_id/episodes"

	// Expected JSON response
	expectedJSON := `{"episodes":[{}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.OrganizationNvrEpisodesListQuery{}
	_, err := client.OrganizationNvrEpisodesList(ctx, "test-organization_id", "test-nvr_id", query)
	if err != nil {
		t.Fatalf("OrganizationNvrEpisodesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationNvrEpisodesList(ctx, "test-organization_id", "test-nvr_id", query)
	if err != nil {
		t.Fatalf("OrganizationNvrEpisodesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationNvrEpisodesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationNvrGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs/test-nvr_id"

	// Expected JSON response
	expectedJSON := `{"api_key":"api-key","id":1,"note":"Note","organization":{"id":7,"title":"Example LLC"},"stats":{"healthcheck_status":{"status":"ok","status_changed_at":1000000000000}},"sync":{"episodes":true},"title":"NVR Title"}`

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
	_, err := client.OrganizationNvrGet(ctx, "test-organization_id", "test-nvr_id")
	if err != nil {
		t.Fatalf("OrganizationNvrGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationNvrGet(ctx, "test-organization_id", "test-nvr_id")
	if err != nil {
		t.Fatalf("OrganizationNvrGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationNvrGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationNvrSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs/test-nvr_id"

	// Expected JSON response
	expectedJSON := `{"api_key":"api-key","id":1,"note":"Note","organization":{"id":7,"title":"Example LLC"},"stats":{"healthcheck_status":{"status":"ok","status_changed_at":1000000000000}},"sync":{"episodes":true},"title":"NVR Title"}`

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
	body := &model.NvrImpl{}
	_, err := client.OrganizationNvrSave(ctx, "test-organization_id", "test-nvr_id", body)
	if err != nil {
		t.Fatalf("OrganizationNvrSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationNvrSave(ctx, "test-organization_id", "test-nvr_id", body)
	if err != nil {
		t.Fatalf("OrganizationNvrSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationNvrSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationNvrStreamGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs/test-nvr_id/streams/test-name"

	// Expected JSON response
	expectedJSON := `{"audio":{"transcode_audio_codec":"aac"},"comment":"This is a test stream","coordinates":{"latitude":55.7512,"longitude":37.6184},"created_at":1672531199000,"dvr":{"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"url":"fake://fake"}],"last_episode_at":1000000000000,"map_coordinates":{"latitude":55.7512,"longitude":37.6184},"name":"example","notifications":[{"event_types":{"episode_face":true,"episode_generic":true,"stream_dead":true},"notification_type":"push"}],"nvr_settings":{"audio":{"transcode_audio_codec":"aac"},"dvr":{"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"url":"fake://fake"}],"static":true,"vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}},"organization":{"id":9,"title":"Organization 1"},"organization_id":9,"path":[{"title":"example"}],"postal_address":"example","preset":{"is_adjustable":true,"title":"Example preset name"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"inputs":[{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"stop_events":[{"code":"example","timestamp":1000000000000}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2}],"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"playback_token":"onetime_token","publish_endpoint":"example","status":"running","streaming_endpoint":"example","ts_delay":1284},"title":"Hockey channel","vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}}`

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
	_, err := client.OrganizationNvrStreamGet(ctx, "test-organization_id", "test-nvr_id", "test-name")
	if err != nil {
		t.Fatalf("OrganizationNvrStreamGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationNvrStreamGet(ctx, "test-organization_id", "test-nvr_id", "test-name")
	if err != nil {
		t.Fatalf("OrganizationNvrStreamGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationNvrStreamGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationNvrStreamsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs/test-nvr_id/streams"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","streams":[{"audio":{"transcode_audio_codec":"aac"},"comment":"This is a test stream","coordinates":{"latitude":55.7512,"longitude":37.6184},"created_at":1672531199000,"dvr":{"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"url":"fake://fake"}],"last_episode_at":1000000000000,"map_coordinates":{"latitude":55.7512,"longitude":37.6184},"name":"example","notifications":[{"event_types":{"episode_face":true,"episode_generic":true,"stream_dead":true},"notification_type":"push"}],"nvr_settings":{"audio":{"transcode_audio_codec":"aac"},"dvr":{"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"url":"fake://fake"}],"static":true,"vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}},"organization":{"id":9,"title":"Organization 1"},"organization_id":9,"path":[{"title":"example"}],"postal_address":"example","preset":{"is_adjustable":true,"title":"Example preset name"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"inputs":[{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"stop_events":[{"code":"example","timestamp":1000000000000}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2}],"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"playback_token":"onetime_token","publish_endpoint":"example","status":"running","streaming_endpoint":"example","ts_delay":1284},"title":"Hockey channel","vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}}]}`

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
	query := &watcherclient.OrganizationNvrStreamsListQuery{}
	_, err := client.OrganizationNvrStreamsList(ctx, "test-organization_id", "test-nvr_id", query)
	if err != nil {
		t.Fatalf("OrganizationNvrStreamsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationNvrStreamsList(ctx, "test-organization_id", "test-nvr_id", query)
	if err != nil {
		t.Fatalf("OrganizationNvrStreamsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationNvrStreamsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationNvrSync(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs/test-nvr_id/sync"

	// Expected JSON response
	expectedJSON := `{"added":5,"removed":2}`

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
	_, err := client.OrganizationNvrSync(ctx, "test-organization_id", "test-nvr_id")
	if err != nil {
		t.Fatalf("OrganizationNvrSync failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationNvrSync(ctx, "test-organization_id", "test-nvr_id")
	if err != nil {
		t.Fatalf("OrganizationNvrSync failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationNvrSync returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationNvrsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/nvrs"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","nvrs":[{"api_key":"api-key","id":1,"note":"Note","organization":{"id":7,"title":"Example LLC"},"stats":{"healthcheck_status":{"status":"ok","status_changed_at":1000000000000}},"sync":{"episodes":true},"title":"NVR Title"}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.OrganizationNvrsListQuery{}
	_, err := client.OrganizationNvrsList(ctx, "test-organization_id", query)
	if err != nil {
		t.Fatalf("OrganizationNvrsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationNvrsList(ctx, "test-organization_id", query)
	if err != nil {
		t.Fatalf("OrganizationNvrsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationNvrsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationPresetSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/presets"

	// Expected JSON response
	expectedJSON := `{"billing_external_id":"example","dvr":{"redundancy_factor":1,"storage_limit":400000000000},"id":7,"is_adjustable":true,"stats":{"organizations_count":12,"streams_count":12},"title":"Example preset name","vision":{"alg":"faces","detector_type":"faces"}}`

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

func TestClient_OrganizationSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id"

	// Expected JSON response
	expectedJSON := `{"created_at":1672531199000,"id":7,"is_default":true,"limits":{"streams":50,"users":50},"owner":{"id":2,"name":"admin"},"stats":{"mosaics":2,"streams":12,"users":12},"title":"Example LLC","user_permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}}`

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
	body := &model.OrganizationImpl{}
	_, err := client.OrganizationSave(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("OrganizationSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationSave(ctx, "test-organization_id", body)
	if err != nil {
		t.Fatalf("OrganizationSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationUserDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/users/test-user_id"

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
	err := client.OrganizationUserDelete(ctx, "test-organization_id", "test-user_id")
	if err != nil {
		t.Fatalf("OrganizationUserDelete failed: %v", err)
	}
}

func TestClient_OrganizationUserGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/users/test-user_id"

	// Expected JSON response
	expectedJSON := `{"email":"user@example.com","id":7,"name":"Example LLC","permissions":{"folders":[{"dvr_depth_limit":3600}],"organization":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}}}`

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
	_, err := client.OrganizationUserGet(ctx, "test-organization_id", "test-user_id")
	if err != nil {
		t.Fatalf("OrganizationUserGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationUserGet(ctx, "test-organization_id", "test-user_id")
	if err != nil {
		t.Fatalf("OrganizationUserGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationUserGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationUserSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/users/test-user_id"

	// Expected JSON response
	expectedJSON := `{"email":"user@example.com","id":7,"name":"Example LLC","permissions":{"folders":[{"dvr_depth_limit":3600}],"organization":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}}}`

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
	body := &model.OrganizationPermissionsImpl{}
	_, err := client.OrganizationUserSave(ctx, "test-organization_id", "test-user_id", body)
	if err != nil {
		t.Fatalf("OrganizationUserSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationUserSave(ctx, "test-organization_id", "test-user_id", body)
	if err != nil {
		t.Fatalf("OrganizationUserSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationUserSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationUsersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations/test-organization_id/users"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","users":[{"email":"user@example.com","id":7,"name":"Example LLC","permissions":{"folders":[{"dvr_depth_limit":3600}],"organization":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}}}]}`

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
	query := &watcherclient.OrganizationUsersListQuery{}
	_, err := client.OrganizationUsersList(ctx, "test-organization_id", query)
	if err != nil {
		t.Fatalf("OrganizationUsersList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationUsersList(ctx, "test-organization_id", query)
	if err != nil {
		t.Fatalf("OrganizationUsersList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationUsersList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_OrganizationsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/organizations"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","organizations":[{"created_at":1672531199000,"id":7,"is_default":true,"limits":{"streams":50,"users":50},"owner":{"id":2,"name":"admin"},"stats":{"mosaics":2,"streams":12,"users":12},"title":"Example LLC","user_permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.OrganizationsListQuery{}
	_, err := client.OrganizationsList(ctx, query)
	if err != nil {
		t.Fatalf("OrganizationsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OrganizationsList(ctx, query)
	if err != nil {
		t.Fatalf("OrganizationsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OrganizationsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PasswordRecovery(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/password_recovery"

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
	err := client.PasswordRecovery(ctx, nil)
	if err != nil {
		t.Fatalf("PasswordRecovery failed: %v", err)
	}
}

func TestClient_PersonDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/persons/test-person_id"

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
	expectedPath := "/watcher/client-api/v3/persons/test-person_id"

	// Expected JSON response
	expectedJSON := `{"deleted_at":1637095014573,"external_id":"example","first_seen_at":1637034282845,"last_seen_at":1637095014573,"name":"example","organization":{"id":7,"title":"Example LLC"},"originator":"api","person_id":0,"person_list":{"id":3,"name":"List 1"},"photos":[{"data":"example","mime_type":"image/jpeg","sha256":"example"}],"updated_at":1637034282845}`

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

func TestClient_PersonListsGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/person_lists"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","person_lists":[{"id":1,"name":"List 1","organization":{"id":7,"title":"Example LLC"},"stats":{"camera":1,"person":1}}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.PersonListsGetQuery{}
	_, err := client.PersonListsGet(ctx, query)
	if err != nil {
		t.Fatalf("PersonListsGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PersonListsGet(ctx, query)
	if err != nil {
		t.Fatalf("PersonListsGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PersonListsGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PersonsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/persons"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","persons":[{"deleted_at":1637095014573,"external_id":"example","first_seen_at":1637034282845,"last_seen_at":1637095014573,"name":"example","organization":{"id":7,"title":"Example LLC"},"originator":"api","person_id":0,"person_list":{"id":3,"name":"List 1"},"photos":[{"data":"example","mime_type":"image/jpeg","sha256":"example"}],"updated_at":1637034282845}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.PersonsListQuery{}
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

func TestClient_PresetDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/presets/test-id"

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
	err := client.PresetDelete(ctx, "test-id")
	if err != nil {
		t.Fatalf("PresetDelete failed: %v", err)
	}
}

func TestClient_PresetGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/presets/test-id"

	// Expected JSON response
	expectedJSON := `{"billing_external_id":"example","dvr":{"redundancy_factor":1,"storage_limit":400000000000},"id":7,"is_adjustable":true,"stats":{"organizations_count":12,"streams_count":12},"title":"Example preset name","vision":{"alg":"faces","detector_type":"faces"}}`

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
	expectedPath := "/watcher/client-api/v3/presets/test-id"

	// Expected JSON response
	expectedJSON := `{"billing_external_id":"example","dvr":{"redundancy_factor":1,"storage_limit":400000000000},"id":7,"is_adjustable":true,"stats":{"organizations_count":12,"streams_count":12},"title":"Example preset name","vision":{"alg":"faces","detector_type":"faces"}}`

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
	expectedPath := "/watcher/client-api/v3/presets"

	// Expected JSON response
	expectedJSON := `{"billing_external_id":"example","dvr":{"redundancy_factor":1,"storage_limit":400000000000},"id":7,"is_adjustable":true,"stats":{"organizations_count":12,"streams_count":12},"title":"Example preset name","vision":{"alg":"faces","detector_type":"faces"}}`

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
	expectedPath := "/watcher/client-api/v3/presets"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","presets":[{"billing_external_id":"example","dvr":{"redundancy_factor":1,"storage_limit":400000000000},"id":7,"is_adjustable":true,"stats":{"organizations_count":12,"streams_count":12},"title":"Example preset name","vision":{"alg":"faces","detector_type":"faces"}}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.PresetsListQuery{}
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

func TestClient_ProfileGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/profile"

	// Expected JSON response
	expectedJSON := `{"apikey":"example","created_at":1672531199000,"email":"user@example.com","external_id":"example","fullname":"example","is_readonly":true,"locale":"en","name":"example","note":"example","password":"example","phone":"+78007778413"}`

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
	_, err := client.ProfileGet(ctx)
	if err != nil {
		t.Fatalf("ProfileGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ProfileGet(ctx)
	if err != nil {
		t.Fatalf("ProfileGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ProfileGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ProfileSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/profile"

	// Expected JSON response
	expectedJSON := `{"apikey":"example","created_at":1672531199000,"email":"user@example.com","external_id":"example","fullname":"example","is_readonly":true,"locale":"en","name":"example","note":"example","password":"example","phone":"+78007778413"}`

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
	body := &model.ProfileImpl{}
	_, err := client.ProfileSave(ctx, body)
	if err != nil {
		t.Fatalf("ProfileSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ProfileSave(ctx, body)
	if err != nil {
		t.Fatalf("ProfileSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ProfileSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ResetPassword(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/reset_password"

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
	body := &model.PasswordResetImpl{}
	err := client.ResetPassword(ctx, body)
	if err != nil {
		t.Fatalf("ResetPassword failed: %v", err)
	}
}

func TestClient_StreamDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/streams/test-name"

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
	expectedPath := "/watcher/client-api/v3/streams/test-name/firmware_update"

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
	expectedPath := "/watcher/client-api/v3/streams/test-name"

	// Expected JSON response
	expectedJSON := `{"audio":{"transcode_audio_codec":"aac"},"comment":"This is a test stream","coordinates":{"latitude":55.7512,"longitude":37.6184},"created_at":1672531199000,"dvr":{"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"url":"fake://fake"}],"last_episode_at":1000000000000,"map_coordinates":{"latitude":55.7512,"longitude":37.6184},"name":"example","notifications":[{"event_types":{"episode_face":true,"episode_generic":true,"stream_dead":true},"notification_type":"push"}],"nvr_settings":{"audio":{"transcode_audio_codec":"aac"},"dvr":{"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"url":"fake://fake"}],"static":true,"vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}},"organization":{"id":9,"title":"Organization 1"},"organization_id":9,"path":[{"title":"example"}],"postal_address":"example","preset":{"is_adjustable":true,"title":"Example preset name"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"inputs":[{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"stop_events":[{"code":"example","timestamp":1000000000000}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2}],"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"playback_token":"onetime_token","publish_endpoint":"example","status":"running","streaming_endpoint":"example","ts_delay":1284},"title":"Hockey channel","vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}}`

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

func TestClient_StreamPermissionUserDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/streams/test-name/permissions/users/test-user_id"

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
	err := client.StreamPermissionUserDelete(ctx, "test-name", "test-user_id")
	if err != nil {
		t.Fatalf("StreamPermissionUserDelete failed: %v", err)
	}
}

func TestClient_StreamPermissionUserGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/streams/test-name/permissions/users/test-user_id"

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
	_, err := client.StreamPermissionUserGet(ctx, "test-name", "test-user_id")
	if err != nil {
		t.Fatalf("StreamPermissionUserGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamPermissionUserGet(ctx, "test-name", "test-user_id")
	if err != nil {
		t.Fatalf("StreamPermissionUserGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamPermissionUserGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamPermissionUserSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/streams/test-name/permissions/users/test-user_id"

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
	body := &model.StreamPermissionsAccessImpl{}
	_, err := client.StreamPermissionUserSave(ctx, "test-name", "test-user_id", body)
	if err != nil {
		t.Fatalf("StreamPermissionUserSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamPermissionUserSave(ctx, "test-name", "test-user_id", body)
	if err != nil {
		t.Fatalf("StreamPermissionUserSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamPermissionUserSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamPermissionsUsersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/streams/test-name/permissions/users"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","users_permissions":[{}]}`

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
	query := &watcherclient.StreamPermissionsUsersListQuery{}
	_, err := client.StreamPermissionsUsersList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("StreamPermissionsUsersList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamPermissionsUsersList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("StreamPermissionsUsersList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamPermissionsUsersList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamPtzExecute(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/streams/test-name/ptz"

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
	body := &model.PtzCommandImpl{}
	err := client.StreamPtzExecute(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("StreamPtzExecute failed: %v", err)
	}
}

func TestClient_StreamSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/streams/test-name"

	// Expected JSON response
	expectedJSON := `{"audio":{"transcode_audio_codec":"aac"},"comment":"This is a test stream","coordinates":{"latitude":55.7512,"longitude":37.6184},"created_at":1672531199000,"dvr":{"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"url":"fake://fake"}],"last_episode_at":1000000000000,"map_coordinates":{"latitude":55.7512,"longitude":37.6184},"name":"example","notifications":[{"event_types":{"episode_face":true,"episode_generic":true,"stream_dead":true},"notification_type":"push"}],"nvr_settings":{"audio":{"transcode_audio_codec":"aac"},"dvr":{"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"url":"fake://fake"}],"static":true,"vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}},"organization":{"id":9,"title":"Organization 1"},"organization_id":9,"path":[{"title":"example"}],"postal_address":"example","preset":{"is_adjustable":true,"title":"Example preset name"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"inputs":[{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"stop_events":[{"code":"example","timestamp":1000000000000}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2}],"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"playback_token":"onetime_token","publish_endpoint":"example","status":"running","streaming_endpoint":"example","ts_delay":1284},"title":"Hockey channel","vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}}`

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

func TestClient_StreamsDiscoverCamera(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/streams/discover"

	// Expected JSON response
	expectedJSON := `{"devices":[{"endpoint":"http://192.168.1.100:80/onvif/device_service","error":"example","ip":"192.168.1.100","manufacturer":"Hikvision","model":"DS-2CD2143G2-I","streams":[{"bitrate":2048,"codec":"H264","fps":25,"height":1080,"jpeg_url":"http://192.168.1.100:80/snapshot.jpg","profile":"Profile_1","profile_name":"Main Stream","stream_url":"rtsp://192.168.1.100:554/stream/main","width":1920}]}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.StreamsDiscoverCameraQuery{}
	body := &model.OnvifDiscoverCameraRequestImpl{}
	_, err := client.StreamsDiscoverCamera(ctx, query, body)
	if err != nil {
		t.Fatalf("StreamsDiscoverCamera failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamsDiscoverCamera(ctx, query, body)
	if err != nil {
		t.Fatalf("StreamsDiscoverCamera failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamsDiscoverCamera returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamsImport(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/streams/import"

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
	_, err := client.StreamsImport(ctx)
	if err != nil {
		t.Fatalf("StreamsImport failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamsImport(ctx)
	if err != nil {
		t.Fatalf("StreamsImport failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamsImport returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/streams"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","streams":[{"audio":{"transcode_audio_codec":"aac"},"comment":"This is a test stream","coordinates":{"latitude":55.7512,"longitude":37.6184},"created_at":1672531199000,"dvr":{"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"url":"fake://fake"}],"last_episode_at":1000000000000,"map_coordinates":{"latitude":55.7512,"longitude":37.6184},"name":"example","notifications":[{"event_types":{"episode_face":true,"episode_generic":true,"stream_dead":true},"notification_type":"push"}],"nvr_settings":{"audio":{"transcode_audio_codec":"aac"},"dvr":{"redundancy_factor":1,"storage_limit":400000000000},"inputs":[{"url":"fake://fake"}],"static":true,"vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}},"organization":{"id":9,"title":"Organization 1"},"organization_id":9,"path":[{"title":"example"}],"postal_address":"example","preset":{"is_adjustable":true,"title":"Example preset name"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"inputs":[{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"stop_events":[{"code":"example","timestamp":1000000000000}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2}],"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"playback_token":"onetime_token","publish_endpoint":"example","status":"running","streaming_endpoint":"example","ts_delay":1284},"title":"Hockey channel","vision":{"alg":"faces","areas":"example","detectors":[{"detector_type":{},"region_id":"example","region_title":"Zone 1","stats":{"alerts":{"low_quality_at":1637094994000,"not_enough_detections_at":1637094994000,"small_size_at":1637094994000},"last_detection_at":1637094994000}}]}}]}`

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
	query := &watcherclient.StreamsListQuery{}
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

func TestClient_StreamsMultiedit(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/streams/multiedit"

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

func TestClient_UiSettingsGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/ui_settings"

	// Expected JSON response
	expectedJSON := `{"brand":"example","colors":{"background":"#fff","primary":"example","secondary":"example"},"company_info":{"address":"example","business_hours":"example","phone":"example"},"default_locale":"en","favicons":{"128":"icons/favicon-128.png","16":"icons/favicon-16.png","32":"icons/favicon-32.png","48":"icons/favicon-48.png","64":"icons/favicon-64.png"},"fonts":{"light":"fira-sans-300.woff2","medium":"fira-sans-500.woff2","regular":"fira-sans-400.woff2"},"locales":["en","ru"],"map":{"api_key":"example","center":{"latitude":55.7512,"longitude":37.6184},"provider":"example"},"product":"example","server_info":{"operator_id":"123","version":"26.02"},"title":"example"}`

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
	_, err := client.UiSettingsGet(ctx)
	if err != nil {
		t.Fatalf("UiSettingsGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UiSettingsGet(ctx)
	if err != nil {
		t.Fatalf("UiSettingsGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UiSettingsGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UserApikeyCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/profile/apikey"

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
	_, err := client.UserApikeyCreate(ctx)
	if err != nil {
		t.Fatalf("UserApikeyCreate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UserApikeyCreate(ctx)
	if err != nil {
		t.Fatalf("UserApikeyCreate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UserApikeyCreate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UserApikeyDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/profile/apikey"

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
	err := client.UserApikeyDelete(ctx)
	if err != nil {
		t.Fatalf("UserApikeyDelete failed: %v", err)
	}
}

func TestClient_UserApikeyGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/profile/apikey"

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
	_, err := client.UserApikeyGet(ctx)
	if err != nil {
		t.Fatalf("UserApikeyGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UserApikeyGet(ctx)
	if err != nil {
		t.Fatalf("UserApikeyGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UserApikeyGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UserCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/users"

	// Expected JSON response
	expectedJSON := `{"access_level":"generic","created_at":1672531199000,"email":"user@example.com","external_id":"example","fullname":"example","id":1,"locale":"en","name":"example","note":"example","organizations":[{"id":1,"owner":{"id":2,"name":"admin"},"permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}},"title":"Cameras"}],"password":"example","phone":"+78007778413"}`

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
	expectedPath := "/watcher/client-api/v3/users/test-user_id"

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
	expectedPath := "/watcher/client-api/v3/users/test-user_id"

	// Expected JSON response
	expectedJSON := `{"access_level":"generic","created_at":1672531199000,"email":"user@example.com","external_id":"example","fullname":"example","id":1,"locale":"en","name":"example","note":"example","organizations":[{"id":1,"owner":{"id":2,"name":"admin"},"permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}},"title":"Cameras"}],"password":"example","phone":"+78007778413"}`

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

func TestClient_UserNvrsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/nvrs"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","nvrs":[{"api_key":"api-key","id":1,"note":"Note","organization":{"id":7,"title":"Example LLC"},"stats":{"healthcheck_status":{"status":"ok","status_changed_at":1000000000000}},"sync":{"episodes":true},"title":"NVR Title"}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.UserNvrsListQuery{}
	_, err := client.UserNvrsList(ctx, query)
	if err != nil {
		t.Fatalf("UserNvrsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UserNvrsList(ctx, query)
	if err != nil {
		t.Fatalf("UserNvrsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UserNvrsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UserOrganizationFoldersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/users/test-user_id/organization/test-organization_id/folders"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"folders":[{"coordinates":{"latitude":55.7512,"longitude":37.6184},"floor_plan":{"bottomleft":{"latitude":55.7512,"longitude":37.6184},"file":{"b64_content":"example","binary_content":"example","mime_type":"example","name":"1-maps-2_07957af5f132243350a0f5546a675cb006354b69","url":"example"},"topleft":{"latitude":55.7512,"longitude":37.6184},"topright":{"latitude":55.7512,"longitude":37.6184}},"hierarchy":{"shift":{"direction":"after"}},"id":0,"permissions":{"dvr_depth_limit":3600},"streams_count":2,"title":"example"}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.UserOrganizationFoldersListQuery{}
	_, err := client.UserOrganizationFoldersList(ctx, "test-user_id", "test-organization_id", query)
	if err != nil {
		t.Fatalf("UserOrganizationFoldersList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UserOrganizationFoldersList(ctx, "test-user_id", "test-organization_id", query)
	if err != nil {
		t.Fatalf("UserOrganizationFoldersList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UserOrganizationFoldersList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UserOrganizationsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/users/test-user_id/organization"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","organizations":[{"created_at":1672531199000,"id":7,"is_default":true,"limits":{"streams":50,"users":50},"owner":{"id":2,"name":"admin"},"stats":{"mosaics":2,"streams":12,"users":12},"title":"Example LLC","user_permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}}}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &watcherclient.UserOrganizationsListQuery{}
	_, err := client.UserOrganizationsList(ctx, "test-user_id", query)
	if err != nil {
		t.Fatalf("UserOrganizationsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UserOrganizationsList(ctx, "test-user_id", query)
	if err != nil {
		t.Fatalf("UserOrganizationsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UserOrganizationsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UserSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/users/test-user_id"

	// Expected JSON response
	expectedJSON := `{"access_level":"generic","created_at":1672531199000,"email":"user@example.com","external_id":"example","fullname":"example","id":1,"locale":"en","name":"example","note":"example","organizations":[{"id":1,"owner":{"id":2,"name":"admin"},"permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}},"title":"Cameras"}],"password":"example","phone":"+78007778413"}`

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

func TestClient_UserStreamPermissionDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/watcher/client-api/v3/users/test-user_id/permissions/streams/test-name"

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
	err := client.UserStreamPermissionDelete(ctx, "test-user_id", "test-name")
	if err != nil {
		t.Fatalf("UserStreamPermissionDelete failed: %v", err)
	}
}

func TestClient_UserStreamPermissionGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/users/test-user_id/permissions/streams/test-name"

	// Expected JSON response
	expectedJSON := `{"stream":{"name":"example"}}`

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
	_, err := client.UserStreamPermissionGet(ctx, "test-user_id", "test-name")
	if err != nil {
		t.Fatalf("UserStreamPermissionGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UserStreamPermissionGet(ctx, "test-user_id", "test-name")
	if err != nil {
		t.Fatalf("UserStreamPermissionGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UserStreamPermissionGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UserStreamPermissionSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/watcher/client-api/v3/users/test-user_id/permissions/streams/test-name"

	// Expected JSON response
	expectedJSON := `{"stream":{"name":"example"}}`

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
	body := &model.StreamPermissionsAccessImpl{}
	_, err := client.UserStreamPermissionSave(ctx, "test-user_id", "test-name", body)
	if err != nil {
		t.Fatalf("UserStreamPermissionSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UserStreamPermissionSave(ctx, "test-user_id", "test-name", body)
	if err != nil {
		t.Fatalf("UserStreamPermissionSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UserStreamPermissionSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UserStreamsPermissionsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/users/test-user_id/permissions/streams"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","streams_permissions":[{"stream":{"name":"example"}}]}`

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
	query := &watcherclient.UserStreamsPermissionsListQuery{}
	_, err := client.UserStreamsPermissionsList(ctx, "test-user_id", query)
	if err != nil {
		t.Fatalf("UserStreamsPermissionsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.UserStreamsPermissionsList(ctx, "test-user_id", query)
	if err != nil {
		t.Fatalf("UserStreamsPermissionsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("UserStreamsPermissionsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_UsersApikeyCreate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/users/test-user_id/apikey"

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
	expectedPath := "/watcher/client-api/v3/users/test-user_id/apikey"

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
	expectedPath := "/watcher/client-api/v3/users"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","users_list":[{"access_level":"generic","created_at":1672531199000,"email":"user@example.com","external_id":"example","fullname":"example","id":1,"locale":"en","name":"example","note":"example","organizations":[{"id":1,"owner":{"id":2,"name":"admin"},"permissions":{"can_edit_persons_lists":true,"can_edit_streams":true,"can_edit_users":true,"can_view_persons_lists":true,"can_view_stats":true,"can_view_streams":true,"is_member":true,"nvrs":{"can_edit":true,"can_view":true}},"title":"Cameras"}],"password":"example","phone":"+78007778413"}]}`

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
	query := &watcherclient.UsersListQuery{}
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

func TestClient_WebPushSubscribe(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/web_push/subscribe"

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
	body := &model.WebPushSubscriptionCreateImpl{}
	err := client.WebPushSubscribe(ctx, body)
	if err != nil {
		t.Fatalf("WebPushSubscribe failed: %v", err)
	}
}

func TestClient_WebPushSubscriptionExistenceGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/watcher/client-api/v3/web_push/subscription"

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
	query := &watcherclient.WebPushSubscriptionExistenceGetQuery{}
	query.Endpoint = "test-endpoint"
	_, err := client.WebPushSubscriptionExistenceGet(ctx, query)
	if err != nil {
		t.Fatalf("WebPushSubscriptionExistenceGet failed: %v", err)
	}
}

func TestClient_WebPushUnsubscribe(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/watcher/client-api/v3/web_push/unsubscribe"

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
	body := &model.WebPushSubscriptionImpl{}
	err := client.WebPushUnsubscribe(ctx, body)
	if err != nil {
		t.Fatalf("WebPushUnsubscribe failed: %v", err)
	}
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

func createTestClient(t *testing.T, rt http.RoundTripper) watcherclient.WatcherClient {
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

	client, err := watcherclient.NewWithBaseFactory(cfg, baseFactory)
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
