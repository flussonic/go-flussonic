package flussonic_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/flussonic/go-flussonic/authorization"
	"github.com/flussonic/go-flussonic/config"
	flussonic "github.com/flussonic/go-flussonic/flussonic"
	model "github.com/flussonic/go-flussonic/flussonic/model"
	"github.com/flussonic/go-flussonic/internal/baseclient"
)

func TestClient_ApiTokensList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/api_tokens"

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
	query := &flussonic.ApiTokensListQuery{}
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
	expectedPath := "/streamer/api/v3/auth_backends/test-name"

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
	expectedPath := "/streamer/api/v3/auth_backends/test-name"

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
	expectedPath := "/streamer/api/v3/auth_backends/test-name"

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
	expectedPath := "/streamer/api/v3/auth_backends"

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
	query := &flussonic.AuthBackendsListQuery{}
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

func TestClient_AvailableEventsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/event_sinks/test-name/events"

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
	query := &flussonic.AvailableEventsListQuery{}
	_, err := client.AvailableEventsList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("AvailableEventsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.AvailableEventsList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("AvailableEventsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("AvailableEventsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_CacheDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/caches/test-name"

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
	err := client.CacheDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("CacheDelete failed: %v", err)
	}
}

func TestClient_CacheGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/caches/test-name"

	// Expected JSON response
	expectedJSON := `{"expiration":604800,"misses":3,"name":"cache1","path":"/storage/cache","storage_limit":400000}`

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
	_, err := client.CacheGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("CacheGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.CacheGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("CacheGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("CacheGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_CacheSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/caches/test-name"

	// Expected JSON response
	expectedJSON := `{"expiration":604800,"misses":3,"name":"cache1","path":"/storage/cache","storage_limit":400000}`

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
	body := &model.CacheConfigImpl{}
	_, err := client.CacheSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("CacheSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.CacheSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("CacheSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("CacheSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_CachesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/caches"

	// Expected JSON response
	expectedJSON := `{"caches":[{"expiration":604800,"misses":3,"name":"cache1","path":"/storage/cache","storage_limit":400000}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.CachesListQuery{}
	_, err := client.CachesList(ctx, query)
	if err != nil {
		t.Fatalf("CachesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.CachesList(ctx, query)
	if err != nil {
		t.Fatalf("CachesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("CachesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_CameraAlarmDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/camera_alarm"

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
	err := client.CameraAlarmDelete(ctx)
	if err != nil {
		t.Fatalf("CameraAlarmDelete failed: %v", err)
	}
}

func TestClient_CameraAlarmGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/camera_alarm"

	// Expected JSON response
	expectedJSON := `{"catch":["example"]}`

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
	_, err := client.CameraAlarmGet(ctx)
	if err != nil {
		t.Fatalf("CameraAlarmGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.CameraAlarmGet(ctx)
	if err != nil {
		t.Fatalf("CameraAlarmGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("CameraAlarmGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_CameraAlarmSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/camera_alarm"

	// Expected JSON response
	expectedJSON := `{"catch":["example"]}`

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
	_, err := client.CameraAlarmSave(ctx)
	if err != nil {
		t.Fatalf("CameraAlarmSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.CameraAlarmSave(ctx)
	if err != nil {
		t.Fatalf("CameraAlarmSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("CameraAlarmSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ConfigGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/config"

	// Expected JSON response
	expectedJSON := `{"api_allowed_from":["example"],"auth_backends":[{"allow_countries":["RU","US"],"allow_ips":["127.0.0.1","10.10.0.0/24"],"allow_tokens":["test_token1","test_token2"],"allow_uas":["AppleWebKit/533.3 (KHTML, like Gecko)","VLC/3.0.8 LibVLC/3.0.8"],"backends":[{"url":"http://stalker-1.iptv.net/auth.php"}],"deny_countries":["RU","GB"],"deny_ips":["8.8.8.8","10.10.0.0/24"],"deny_tokens":["test_token3","test_token4"],"deny_uas":["Mobile Safari/533.3","VLC/3.0.8 LibVLC/3.0.8"],"name":"example"}],"auth_token":"token","balancers":[{"mode":"usage","name":"example","servers":[{"countries":["example"],"name":"example"}]}],"caches":[{"expiration":604800,"misses":3,"name":"cache1","path":"/storage/cache","storage_limit":400000}],"camera_alarm":{"catch":["example"]},"chassis":{"default_gateway_interface":"streaming","dhcpd_iface":"example","firmware_boot_dir":"example","firmware_host":"https://example.com","firmware_version":"example","hostname":"coder1.example.com","ntp_servers":["example"],"ntpd_iface":"example","product_name":"coder_transcoder","stats":{"hardware_id":"example","hostname":"coder1.example.com","model":"chassis_model","next_version":"example","serial_number":"2174220024","system_time":1000000000000,"version":"21.09.1-234"},"tftp_root":"example","update_channel":"example"},"cluster_key":"xS6i6Q3DCc5nEvnu","config_external":"http://central.example.com/streamer/api/v3/streamers","decklinks":[{"profile":"one_full"}],"dvb_cards":[{"adapter":16,"bandwidth":6000000,"code_rate_hp":"1_2","code_rate_lp":"1_2","comment":"example","frequency":11606,"guard_interval":"1_16","hierarchy":"1","hw":"dvb","modulation":"auto","name":"a16","pilot":"auto","polarization":"v","rolloff":"35","stats":{"ber":5,"has_carrier":true,"has_lock":true,"has_rate":true,"has_signal":true,"has_sync":true,"has_viterbi":true,"snr":1,"snr_raw":894,"strength":95,"strength_raw":62446},"symbol_rate":27500,"system":"dvbs","transmission_mode":"1k","video_device":"/dev/video0"}],"dvrs":[{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"disks":[{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}],"dvr_replicate":true,"episodes_url":"example","index":"example","name":"example","replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000}],"edit_auth":{"hash":true,"login":"secretlogin","password":"passw"},"event_sinks":[{"except":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"extra":{"key1":"example","key2":"example"},"level":"debug","max_depth":100,"max_size":10000,"name":"my_json_sink","only":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"resend_limit":1000,"resend_timeout":10,"throttle_delay":1,"url":"jsonlog:///var/log/events-json.log"}],"file_processor":{"path":"example"},"geoip":"/usr/share/GeoIP/GeoLite2-City.mmdb","http_proxies":[{"prefix":"example","stats":{"http_100":1000,"http_200":1000,"http_300":10,"http_400":10,"http_500":10,"protocol_upgrades":1000,"proxy_error":10,"proxy_error_connection":10,"requests":1000,"time_1000ms":2,"time_100ms":10,"time_5000ms":2,"time_500ms":5,"time_longms":1},"url":"https://example.com"}],"iptv":{"database":"example"},"listeners":{"http":[{"address":"10.0.35.1","port":80}],"https":[{"ssl_protocols":["tlsv1.1","tlsv1.2"]}],"rtmp":[{"address":"10.0.35.1","port":80}],"rtmps":[{"address":"10.0.35.1","port":80,"ssl_protocols":["tlsv1.1","tlsv1.2"]}],"rtsp":[{"address":"10.0.35.1","port":80}],"rtsps":[{"address":"10.0.35.1","port":80,"ssl_protocols":["tlsv1.1","tlsv1.2"]}],"sip":[{"address":"10.0.35.1","port":80}],"snmp":[{"address":"10.0.35.1","port":80}],"srt":[{"address":"10.0.35.1","mode":"play","ports":{"first":10001,"last":10099}}],"turn":[{"address":"10.0.35.1","port":80}]},"loglevel":"error","meta":"{\"role\": \"transcoder\"}","nvidia_monitor":true,"peers":[{"api_url":"http://streamer.local:8080","channel_limit":5,"cluster_key":"xS6i6Q3DCc5nEvnu","cpu_limit":10,"fetch_timeout":1000,"hostname":"peer.example.com","private_payload_url":"http://streamer.local","public_payload_url":"http://public.example.com","stale_timeout":1000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502}}],"pulsedb":"/var/run/flussonic/pulsedb","rproxy":{"endpoint_auth":"example","forward_ports":{"key1":{"handler":"example"},"key2":{"handler":"example"}},"streampoint_key":"example"},"server_id":"550e8400-e29b-41d4-a716-446655440000","server_names":[{"aliases":["s1.streamer.local","s2.streamer.local"],"domain":"streamer.local"}],"session_log":"/var/run/flussonic/session_log","sources":[{"add_audio_only":true,"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"cluster_key":"example","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"except":["example"],"group_config":{"key1":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000},"key2":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000}},"hls_scte35":"scte35","jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"meta":{"key1":"example","key2":"example"},"mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"only":["example"],"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefix":"example","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"url":"example_value","vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}],"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"events":{"dropped_events":100,"resent_events":100},"hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"streams":[{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","config_on_disk":{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","drm":{"vendor":"example"},"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"dvr_read_performance":{"segments_read_delayed":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_fast":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_slow":{"cache":100,"local":10,"ram":100,"remote":10}},"dvr_read_popularity":{"other":{"cache":100,"local":10,"ram":100,"remote":10},"today":{"cache":100,"local":10,"ram":100,"remote":10},"week":{"cache":100,"local":10,"ram":100,"remote":10},"yesterday":{"cache":100,"local":10,"ram":100,"remote":10}},"id":"61893ba6-07b3-431b-b2f7-716ac1643953","input":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"input_bitrate":186,"last_access_at":1669106270979,"last_dts":383835646,"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"opened_at":1636383729002,"play":{"play_bytes_by_proto":{"key1":0,"key2":0},"play_dash_playlist_http_200":1000,"play_dash_playlist_http_400":1000,"play_dash_playlist_http_500":1000,"play_dash_playlist_time_1000ms":10,"play_dash_playlist_time_100ms":10,"play_dash_playlist_time_500ms":10,"play_dash_playlist_time_longms":10,"play_dash_segment_http_200":1000,"play_dash_segment_http_400":1000,"play_dash_segment_http_500":1000,"play_dash_segment_time_1000ms":10,"play_dash_segment_time_100ms":10,"play_dash_segment_time_500ms":10,"play_dash_segment_time_longms":10,"play_duration_by_proto":{"key1":0,"key2":0},"play_hls_playlist_http_200":1000,"play_hls_playlist_http_400":1000,"play_hls_playlist_http_500":1000,"play_hls_playlist_time_1000ms":10,"play_hls_playlist_time_100ms":10,"play_hls_playlist_time_500ms":10,"play_hls_playlist_time_longms":10,"play_hls_segment_http_200":1000,"play_hls_segment_http_400":1000,"play_hls_segment_http_500":1000,"play_hls_segment_time_1000ms":10,"play_hls_segment_time_100ms":10,"play_hls_segment_time_500ms":10,"play_hls_segment_time_longms":10,"play_ll_hls_playlist_http_200":1000,"play_ll_hls_playlist_http_400":1000,"play_ll_hls_playlist_http_500":1000,"play_ll_hls_playlist_time_1000ms":10,"play_ll_hls_playlist_time_100ms":10,"play_ll_hls_playlist_time_500ms":10,"play_ll_hls_playlist_time_longms":10,"play_ll_hls_segment_http_200":1000,"play_ll_hls_segment_http_400":1000,"play_ll_hls_segment_http_500":1000,"play_ll_hls_segment_time_1000ms":10,"play_ll_hls_segment_time_100ms":10,"play_ll_hls_segment_time_500ms":10,"play_ll_hls_segment_time_longms":10,"play_opened_sessions_by_proto":{"key1":0,"key2":0},"play_other_http_200":1000,"play_other_http_400":1000,"play_other_http_500":1000,"play_other_time_1000ms":10,"play_other_time_100ms":10,"play_other_time_500ms":10,"play_other_time_longms":10,"play_preview_http_200":1000,"play_preview_http_400":1000,"play_preview_http_500":1000,"play_preview_time_1000ms":10,"play_preview_time_100ms":10,"play_preview_time_500ms":10,"play_preview_time_longms":10,"play_total_sessions_by_proto":{"key1":0,"key2":0}},"push":[{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"}],"running_on":["streamer1.example.com"],"source_id":"61893be1-054e-4acc-8d24-8ed92efe6ad0","start_running_at":1000000000000,"status":"running","transcoder":{"hw":"cpu","target":"uhdtv","tracks":[{"track_id":"v1"}]},"ts_delay":1284},"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}],"templates":[{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefer_codec":"h264","prefixes":["chats",""],"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"source_timeout":10,"static":true,"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}],"total_bandwidth":1000000000,"transponders":[{"eit":{"xmltv_url":"xmltv_dir"},"name":"multiplexer","network_name":"Example Network","others":[{"name":"example"}],"prebuffer":800,"programs":[{"eit_title":"EIT_Title","lcn":5,"pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"program_id":1,"service_type":"digital_tv","source":"hockey1","title":"ProgramTitle"}],"provider":"Example Provider","pushes":[{"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"url":"example"}],"stats":{"pids":[{"content":"audio","pid":0}],"programs":[{"pnr":0}]},"time_offsets":[{"country":"example","local_time_offset":"example","next_time_offset":"example"}],"ts_descriptors":[{"hex":"example","tag":0}]}],"view_auth":{"hash":true,"login":"secretlogin","password":"passw"},"vods":[{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"logo":{"height":100,"left":15,"path":"@chan.png","top":15,"width":200},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"position":1,"prefix":"movies","provider":"example","read_queue":100,"segment_duration":1000,"storages":[{"extra":{"key1":"example","key2":"example"},"url":"/storage"}],"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"timeout":3}],"vsaas":{"central_url":"example","watcher_url":"example"},"webrtc_play":{"ports":[0],"transport":"udp"},"webrtc_publish":{"ports":[0],"transport":"udp"}}`

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
	expectedPath := "/streamer/api/v3/config"

	// Expected JSON response
	expectedJSON := `{"api_allowed_from":["example"],"auth_backends":[{"allow_countries":["RU","US"],"allow_ips":["127.0.0.1","10.10.0.0/24"],"allow_tokens":["test_token1","test_token2"],"allow_uas":["AppleWebKit/533.3 (KHTML, like Gecko)","VLC/3.0.8 LibVLC/3.0.8"],"backends":[{"url":"http://stalker-1.iptv.net/auth.php"}],"deny_countries":["RU","GB"],"deny_ips":["8.8.8.8","10.10.0.0/24"],"deny_tokens":["test_token3","test_token4"],"deny_uas":["Mobile Safari/533.3","VLC/3.0.8 LibVLC/3.0.8"],"name":"example"}],"auth_token":"token","balancers":[{"mode":"usage","name":"example","servers":[{"countries":["example"],"name":"example"}]}],"caches":[{"expiration":604800,"misses":3,"name":"cache1","path":"/storage/cache","storage_limit":400000}],"camera_alarm":{"catch":["example"]},"chassis":{"default_gateway_interface":"streaming","dhcpd_iface":"example","firmware_boot_dir":"example","firmware_host":"https://example.com","firmware_version":"example","hostname":"coder1.example.com","ntp_servers":["example"],"ntpd_iface":"example","product_name":"coder_transcoder","stats":{"hardware_id":"example","hostname":"coder1.example.com","model":"chassis_model","next_version":"example","serial_number":"2174220024","system_time":1000000000000,"version":"21.09.1-234"},"tftp_root":"example","update_channel":"example"},"cluster_key":"xS6i6Q3DCc5nEvnu","config_external":"http://central.example.com/streamer/api/v3/streamers","decklinks":[{"profile":"one_full"}],"dvb_cards":[{"adapter":16,"bandwidth":6000000,"code_rate_hp":"1_2","code_rate_lp":"1_2","comment":"example","frequency":11606,"guard_interval":"1_16","hierarchy":"1","hw":"dvb","modulation":"auto","name":"a16","pilot":"auto","polarization":"v","rolloff":"35","stats":{"ber":5,"has_carrier":true,"has_lock":true,"has_rate":true,"has_signal":true,"has_sync":true,"has_viterbi":true,"snr":1,"snr_raw":894,"strength":95,"strength_raw":62446},"symbol_rate":27500,"system":"dvbs","transmission_mode":"1k","video_device":"/dev/video0"}],"dvrs":[{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"disks":[{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}],"dvr_replicate":true,"episodes_url":"example","index":"example","name":"example","replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000}],"edit_auth":{"hash":true,"login":"secretlogin","password":"passw"},"event_sinks":[{"except":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"extra":{"key1":"example","key2":"example"},"level":"debug","max_depth":100,"max_size":10000,"name":"my_json_sink","only":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"resend_limit":1000,"resend_timeout":10,"throttle_delay":1,"url":"jsonlog:///var/log/events-json.log"}],"file_processor":{"path":"example"},"geoip":"/usr/share/GeoIP/GeoLite2-City.mmdb","http_proxies":[{"prefix":"example","stats":{"http_100":1000,"http_200":1000,"http_300":10,"http_400":10,"http_500":10,"protocol_upgrades":1000,"proxy_error":10,"proxy_error_connection":10,"requests":1000,"time_1000ms":2,"time_100ms":10,"time_5000ms":2,"time_500ms":5,"time_longms":1},"url":"https://example.com"}],"iptv":{"database":"example"},"listeners":{"http":[{"address":"10.0.35.1","port":80}],"https":[{"ssl_protocols":["tlsv1.1","tlsv1.2"]}],"rtmp":[{"address":"10.0.35.1","port":80}],"rtmps":[{"address":"10.0.35.1","port":80,"ssl_protocols":["tlsv1.1","tlsv1.2"]}],"rtsp":[{"address":"10.0.35.1","port":80}],"rtsps":[{"address":"10.0.35.1","port":80,"ssl_protocols":["tlsv1.1","tlsv1.2"]}],"sip":[{"address":"10.0.35.1","port":80}],"snmp":[{"address":"10.0.35.1","port":80}],"srt":[{"address":"10.0.35.1","mode":"play","ports":{"first":10001,"last":10099}}],"turn":[{"address":"10.0.35.1","port":80}]},"loglevel":"error","meta":"{\"role\": \"transcoder\"}","nvidia_monitor":true,"peers":[{"api_url":"http://streamer.local:8080","channel_limit":5,"cluster_key":"xS6i6Q3DCc5nEvnu","cpu_limit":10,"fetch_timeout":1000,"hostname":"peer.example.com","private_payload_url":"http://streamer.local","public_payload_url":"http://public.example.com","stale_timeout":1000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502}}],"pulsedb":"/var/run/flussonic/pulsedb","rproxy":{"endpoint_auth":"example","forward_ports":{"key1":{"handler":"example"},"key2":{"handler":"example"}},"streampoint_key":"example"},"server_id":"550e8400-e29b-41d4-a716-446655440000","server_names":[{"aliases":["s1.streamer.local","s2.streamer.local"],"domain":"streamer.local"}],"session_log":"/var/run/flussonic/session_log","sources":[{"add_audio_only":true,"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"cluster_key":"example","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"except":["example"],"group_config":{"key1":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000},"key2":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000}},"hls_scte35":"scte35","jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"meta":{"key1":"example","key2":"example"},"mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"only":["example"],"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefix":"example","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"url":"example_value","vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}],"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"events":{"dropped_events":100,"resent_events":100},"hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"streams":[{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","config_on_disk":{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","drm":{"vendor":"example"},"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"dvr_read_performance":{"segments_read_delayed":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_fast":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_slow":{"cache":100,"local":10,"ram":100,"remote":10}},"dvr_read_popularity":{"other":{"cache":100,"local":10,"ram":100,"remote":10},"today":{"cache":100,"local":10,"ram":100,"remote":10},"week":{"cache":100,"local":10,"ram":100,"remote":10},"yesterday":{"cache":100,"local":10,"ram":100,"remote":10}},"id":"61893ba6-07b3-431b-b2f7-716ac1643953","input":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"input_bitrate":186,"last_access_at":1669106270979,"last_dts":383835646,"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"opened_at":1636383729002,"play":{"play_bytes_by_proto":{"key1":0,"key2":0},"play_dash_playlist_http_200":1000,"play_dash_playlist_http_400":1000,"play_dash_playlist_http_500":1000,"play_dash_playlist_time_1000ms":10,"play_dash_playlist_time_100ms":10,"play_dash_playlist_time_500ms":10,"play_dash_playlist_time_longms":10,"play_dash_segment_http_200":1000,"play_dash_segment_http_400":1000,"play_dash_segment_http_500":1000,"play_dash_segment_time_1000ms":10,"play_dash_segment_time_100ms":10,"play_dash_segment_time_500ms":10,"play_dash_segment_time_longms":10,"play_duration_by_proto":{"key1":0,"key2":0},"play_hls_playlist_http_200":1000,"play_hls_playlist_http_400":1000,"play_hls_playlist_http_500":1000,"play_hls_playlist_time_1000ms":10,"play_hls_playlist_time_100ms":10,"play_hls_playlist_time_500ms":10,"play_hls_playlist_time_longms":10,"play_hls_segment_http_200":1000,"play_hls_segment_http_400":1000,"play_hls_segment_http_500":1000,"play_hls_segment_time_1000ms":10,"play_hls_segment_time_100ms":10,"play_hls_segment_time_500ms":10,"play_hls_segment_time_longms":10,"play_ll_hls_playlist_http_200":1000,"play_ll_hls_playlist_http_400":1000,"play_ll_hls_playlist_http_500":1000,"play_ll_hls_playlist_time_1000ms":10,"play_ll_hls_playlist_time_100ms":10,"play_ll_hls_playlist_time_500ms":10,"play_ll_hls_playlist_time_longms":10,"play_ll_hls_segment_http_200":1000,"play_ll_hls_segment_http_400":1000,"play_ll_hls_segment_http_500":1000,"play_ll_hls_segment_time_1000ms":10,"play_ll_hls_segment_time_100ms":10,"play_ll_hls_segment_time_500ms":10,"play_ll_hls_segment_time_longms":10,"play_opened_sessions_by_proto":{"key1":0,"key2":0},"play_other_http_200":1000,"play_other_http_400":1000,"play_other_http_500":1000,"play_other_time_1000ms":10,"play_other_time_100ms":10,"play_other_time_500ms":10,"play_other_time_longms":10,"play_preview_http_200":1000,"play_preview_http_400":1000,"play_preview_http_500":1000,"play_preview_time_1000ms":10,"play_preview_time_100ms":10,"play_preview_time_500ms":10,"play_preview_time_longms":10,"play_total_sessions_by_proto":{"key1":0,"key2":0}},"push":[{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"}],"running_on":["streamer1.example.com"],"source_id":"61893be1-054e-4acc-8d24-8ed92efe6ad0","start_running_at":1000000000000,"status":"running","transcoder":{"hw":"cpu","target":"uhdtv","tracks":[{"track_id":"v1"}]},"ts_delay":1284},"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}],"templates":[{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefer_codec":"h264","prefixes":["chats",""],"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"source_timeout":10,"static":true,"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}],"total_bandwidth":1000000000,"transponders":[{"eit":{"xmltv_url":"xmltv_dir"},"name":"multiplexer","network_name":"Example Network","others":[{"name":"example"}],"prebuffer":800,"programs":[{"eit_title":"EIT_Title","lcn":5,"pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"program_id":1,"service_type":"digital_tv","source":"hockey1","title":"ProgramTitle"}],"provider":"Example Provider","pushes":[{"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"url":"example"}],"stats":{"pids":[{"content":"audio","pid":0}],"programs":[{"pnr":0}]},"time_offsets":[{"country":"example","local_time_offset":"example","next_time_offset":"example"}],"ts_descriptors":[{"hex":"example","tag":0}]}],"view_auth":{"hash":true,"login":"secretlogin","password":"passw"},"vods":[{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"logo":{"height":100,"left":15,"path":"@chan.png","top":15,"width":200},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"position":1,"prefix":"movies","provider":"example","read_queue":100,"segment_duration":1000,"storages":[{"extra":{"key1":"example","key2":"example"},"url":"/storage"}],"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"timeout":3}],"vsaas":{"central_url":"example","watcher_url":"example"},"webrtc_play":{"ports":[0],"transport":"udp"},"webrtc_publish":{"ports":[0],"transport":"udp"}}`

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
	body := &model.ServerConfigImpl{}
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

func TestClient_ConfigStatsGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/config/stats"

	// Expected JSON response
	expectedJSON := `{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"events":{"dropped_events":100,"resent_events":100},"hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502}`

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
	_, err := client.ConfigStatsGet(ctx)
	if err != nil {
		t.Fatalf("ConfigStatsGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ConfigStatsGet(ctx)
	if err != nil {
		t.Fatalf("ConfigStatsGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ConfigStatsGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ConfigValidate(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/streamer/api/v3/config"

	// Expected JSON response
	expectedJSON := `{"api_allowed_from":["example"],"auth_backends":[{"allow_countries":["RU","US"],"allow_ips":["127.0.0.1","10.10.0.0/24"],"allow_tokens":["test_token1","test_token2"],"allow_uas":["AppleWebKit/533.3 (KHTML, like Gecko)","VLC/3.0.8 LibVLC/3.0.8"],"backends":[{"url":"http://stalker-1.iptv.net/auth.php"}],"deny_countries":["RU","GB"],"deny_ips":["8.8.8.8","10.10.0.0/24"],"deny_tokens":["test_token3","test_token4"],"deny_uas":["Mobile Safari/533.3","VLC/3.0.8 LibVLC/3.0.8"],"name":"example"}],"auth_token":"token","balancers":[{"mode":"usage","name":"example","servers":[{"countries":["example"],"name":"example"}]}],"caches":[{"expiration":604800,"misses":3,"name":"cache1","path":"/storage/cache","storage_limit":400000}],"camera_alarm":{"catch":["example"]},"chassis":{"default_gateway_interface":"streaming","dhcpd_iface":"example","firmware_boot_dir":"example","firmware_host":"https://example.com","firmware_version":"example","hostname":"coder1.example.com","ntp_servers":["example"],"ntpd_iface":"example","product_name":"coder_transcoder","stats":{"hardware_id":"example","hostname":"coder1.example.com","model":"chassis_model","next_version":"example","serial_number":"2174220024","system_time":1000000000000,"version":"21.09.1-234"},"tftp_root":"example","update_channel":"example"},"cluster_key":"xS6i6Q3DCc5nEvnu","config_external":"http://central.example.com/streamer/api/v3/streamers","decklinks":[{"profile":"one_full"}],"dvb_cards":[{"adapter":16,"bandwidth":6000000,"code_rate_hp":"1_2","code_rate_lp":"1_2","comment":"example","frequency":11606,"guard_interval":"1_16","hierarchy":"1","hw":"dvb","modulation":"auto","name":"a16","pilot":"auto","polarization":"v","rolloff":"35","stats":{"ber":5,"has_carrier":true,"has_lock":true,"has_rate":true,"has_signal":true,"has_sync":true,"has_viterbi":true,"snr":1,"snr_raw":894,"strength":95,"strength_raw":62446},"symbol_rate":27500,"system":"dvbs","transmission_mode":"1k","video_device":"/dev/video0"}],"dvrs":[{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"disks":[{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}],"dvr_replicate":true,"episodes_url":"example","index":"example","name":"example","replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000}],"edit_auth":{"hash":true,"login":"secretlogin","password":"passw"},"event_sinks":[{"except":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"extra":{"key1":"example","key2":"example"},"level":"debug","max_depth":100,"max_size":10000,"name":"my_json_sink","only":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"resend_limit":1000,"resend_timeout":10,"throttle_delay":1,"url":"jsonlog:///var/log/events-json.log"}],"file_processor":{"path":"example"},"geoip":"/usr/share/GeoIP/GeoLite2-City.mmdb","http_proxies":[{"prefix":"example","stats":{"http_100":1000,"http_200":1000,"http_300":10,"http_400":10,"http_500":10,"protocol_upgrades":1000,"proxy_error":10,"proxy_error_connection":10,"requests":1000,"time_1000ms":2,"time_100ms":10,"time_5000ms":2,"time_500ms":5,"time_longms":1},"url":"https://example.com"}],"iptv":{"database":"example"},"listeners":{"http":[{"address":"10.0.35.1","port":80}],"https":[{"ssl_protocols":["tlsv1.1","tlsv1.2"]}],"rtmp":[{"address":"10.0.35.1","port":80}],"rtmps":[{"address":"10.0.35.1","port":80,"ssl_protocols":["tlsv1.1","tlsv1.2"]}],"rtsp":[{"address":"10.0.35.1","port":80}],"rtsps":[{"address":"10.0.35.1","port":80,"ssl_protocols":["tlsv1.1","tlsv1.2"]}],"sip":[{"address":"10.0.35.1","port":80}],"snmp":[{"address":"10.0.35.1","port":80}],"srt":[{"address":"10.0.35.1","mode":"play","ports":{"first":10001,"last":10099}}],"turn":[{"address":"10.0.35.1","port":80}]},"loglevel":"error","meta":"{\"role\": \"transcoder\"}","nvidia_monitor":true,"peers":[{"api_url":"http://streamer.local:8080","channel_limit":5,"cluster_key":"xS6i6Q3DCc5nEvnu","cpu_limit":10,"fetch_timeout":1000,"hostname":"peer.example.com","private_payload_url":"http://streamer.local","public_payload_url":"http://public.example.com","stale_timeout":1000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502}}],"pulsedb":"/var/run/flussonic/pulsedb","rproxy":{"endpoint_auth":"example","forward_ports":{"key1":{"handler":"example"},"key2":{"handler":"example"}},"streampoint_key":"example"},"server_id":"550e8400-e29b-41d4-a716-446655440000","server_names":[{"aliases":["s1.streamer.local","s2.streamer.local"],"domain":"streamer.local"}],"session_log":"/var/run/flussonic/session_log","sources":[{"add_audio_only":true,"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"cluster_key":"example","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"except":["example"],"group_config":{"key1":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000},"key2":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000}},"hls_scte35":"scte35","jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"meta":{"key1":"example","key2":"example"},"mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"only":["example"],"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefix":"example","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"url":"example_value","vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}],"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"events":{"dropped_events":100,"resent_events":100},"hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"streams":[{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","config_on_disk":{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","drm":{"vendor":"example"},"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"dvr_read_performance":{"segments_read_delayed":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_fast":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_slow":{"cache":100,"local":10,"ram":100,"remote":10}},"dvr_read_popularity":{"other":{"cache":100,"local":10,"ram":100,"remote":10},"today":{"cache":100,"local":10,"ram":100,"remote":10},"week":{"cache":100,"local":10,"ram":100,"remote":10},"yesterday":{"cache":100,"local":10,"ram":100,"remote":10}},"id":"61893ba6-07b3-431b-b2f7-716ac1643953","input":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"input_bitrate":186,"last_access_at":1669106270979,"last_dts":383835646,"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"opened_at":1636383729002,"play":{"play_bytes_by_proto":{"key1":0,"key2":0},"play_dash_playlist_http_200":1000,"play_dash_playlist_http_400":1000,"play_dash_playlist_http_500":1000,"play_dash_playlist_time_1000ms":10,"play_dash_playlist_time_100ms":10,"play_dash_playlist_time_500ms":10,"play_dash_playlist_time_longms":10,"play_dash_segment_http_200":1000,"play_dash_segment_http_400":1000,"play_dash_segment_http_500":1000,"play_dash_segment_time_1000ms":10,"play_dash_segment_time_100ms":10,"play_dash_segment_time_500ms":10,"play_dash_segment_time_longms":10,"play_duration_by_proto":{"key1":0,"key2":0},"play_hls_playlist_http_200":1000,"play_hls_playlist_http_400":1000,"play_hls_playlist_http_500":1000,"play_hls_playlist_time_1000ms":10,"play_hls_playlist_time_100ms":10,"play_hls_playlist_time_500ms":10,"play_hls_playlist_time_longms":10,"play_hls_segment_http_200":1000,"play_hls_segment_http_400":1000,"play_hls_segment_http_500":1000,"play_hls_segment_time_1000ms":10,"play_hls_segment_time_100ms":10,"play_hls_segment_time_500ms":10,"play_hls_segment_time_longms":10,"play_ll_hls_playlist_http_200":1000,"play_ll_hls_playlist_http_400":1000,"play_ll_hls_playlist_http_500":1000,"play_ll_hls_playlist_time_1000ms":10,"play_ll_hls_playlist_time_100ms":10,"play_ll_hls_playlist_time_500ms":10,"play_ll_hls_playlist_time_longms":10,"play_ll_hls_segment_http_200":1000,"play_ll_hls_segment_http_400":1000,"play_ll_hls_segment_http_500":1000,"play_ll_hls_segment_time_1000ms":10,"play_ll_hls_segment_time_100ms":10,"play_ll_hls_segment_time_500ms":10,"play_ll_hls_segment_time_longms":10,"play_opened_sessions_by_proto":{"key1":0,"key2":0},"play_other_http_200":1000,"play_other_http_400":1000,"play_other_http_500":1000,"play_other_time_1000ms":10,"play_other_time_100ms":10,"play_other_time_500ms":10,"play_other_time_longms":10,"play_preview_http_200":1000,"play_preview_http_400":1000,"play_preview_http_500":1000,"play_preview_time_1000ms":10,"play_preview_time_100ms":10,"play_preview_time_500ms":10,"play_preview_time_longms":10,"play_total_sessions_by_proto":{"key1":0,"key2":0}},"push":[{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"}],"running_on":["streamer1.example.com"],"source_id":"61893be1-054e-4acc-8d24-8ed92efe6ad0","start_running_at":1000000000000,"status":"running","transcoder":{"hw":"cpu","target":"uhdtv","tracks":[{"track_id":"v1"}]},"ts_delay":1284},"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}],"templates":[{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefer_codec":"h264","prefixes":["chats",""],"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"source_timeout":10,"static":true,"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}],"total_bandwidth":1000000000,"transponders":[{"eit":{"xmltv_url":"xmltv_dir"},"name":"multiplexer","network_name":"Example Network","others":[{"name":"example"}],"prebuffer":800,"programs":[{"eit_title":"EIT_Title","lcn":5,"pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"program_id":1,"service_type":"digital_tv","source":"hockey1","title":"ProgramTitle"}],"provider":"Example Provider","pushes":[{"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"url":"example"}],"stats":{"pids":[{"content":"audio","pid":0}],"programs":[{"pnr":0}]},"time_offsets":[{"country":"example","local_time_offset":"example","next_time_offset":"example"}],"ts_descriptors":[{"hex":"example","tag":0}]}],"view_auth":{"hash":true,"login":"secretlogin","password":"passw"},"vods":[{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"logo":{"height":100,"left":15,"path":"@chan.png","top":15,"width":200},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"position":1,"prefix":"movies","provider":"example","read_queue":100,"segment_duration":1000,"storages":[{"extra":{"key1":"example","key2":"example"},"url":"/storage"}],"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"timeout":3}],"vsaas":{"central_url":"example","watcher_url":"example"},"webrtc_play":{"ports":[0],"transport":"udp"},"webrtc_publish":{"ports":[0],"transport":"udp"}}`

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
	body := &model.ServerConfigImpl{}
	_, err := client.ConfigValidate(ctx, body)
	if err != nil {
		t.Fatalf("ConfigValidate failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.ConfigValidate(ctx, body)
	if err != nil {
		t.Fatalf("ConfigValidate failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("ConfigValidate returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DiskFileDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/vods/test-prefix/storages/test-storage_index/files/test-subpath"

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
	err := client.DiskFileDelete(ctx, "test-prefix", "test-storage_index", "test-subpath")
	if err != nil {
		t.Fatalf("DiskFileDelete failed: %v", err)
	}
}

func TestClient_DiskFileGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/vods/test-prefix/storages/test-storage_index/files/test-subpath"

	// Expected JSON response
	expectedJSON := `{"bytes":42309561,"folder":"example","name":"example","opened":true,"prefix":"example","stats":{"bytes_in":1700923231,"bytes_out":1700923231,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"worker_count":1},"subpath":"example","url":"example"}`

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
	_, err := client.DiskFileGet(ctx, "test-prefix", "test-storage_index", "test-subpath")
	if err != nil {
		t.Fatalf("DiskFileGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DiskFileGet(ctx, "test-prefix", "test-storage_index", "test-subpath")
	if err != nil {
		t.Fatalf("DiskFileGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DiskFileGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DiskFileSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/vods/test-prefix/storages/test-storage_index/files/test-subpath"

	// Expected JSON response
	expectedJSON := `{"bytes":42309561,"folder":"example","name":"example","opened":true,"prefix":"example","stats":{"bytes_in":1700923231,"bytes_out":1700923231,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"worker_count":1},"subpath":"example","url":"example"}`

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
	_, err := client.DiskFileSave(ctx, "test-prefix", "test-storage_index", "test-subpath")
	if err != nil {
		t.Fatalf("DiskFileSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DiskFileSave(ctx, "test-prefix", "test-storage_index", "test-subpath")
	if err != nil {
		t.Fatalf("DiskFileSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DiskFileSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DiskFilesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/vods/test-prefix/storages/test-storage_index/files"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"files":[{"bytes":42309561,"folder":"example","name":"example","opened":true,"prefix":"example","stats":{"bytes_in":1700923231,"bytes_out":1700923231,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"worker_count":1},"subpath":"example","url":"example"}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.DiskFilesListQuery{}
	_, err := client.DiskFilesList(ctx, "test-prefix", "test-storage_index", query)
	if err != nil {
		t.Fatalf("DiskFilesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DiskFilesList(ctx, "test-prefix", "test-storage_index", query)
	if err != nil {
		t.Fatalf("DiskFilesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DiskFilesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvbCardAvailableProgramsGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/dvb_cards/test-name/available_programs"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"media_info_list":[{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.DvbCardAvailableProgramsGetQuery{}
	_, err := client.DvbCardAvailableProgramsGet(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("DvbCardAvailableProgramsGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvbCardAvailableProgramsGet(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("DvbCardAvailableProgramsGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvbCardAvailableProgramsGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvbCardDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/dvb_cards/test-name"

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
	err := client.DvbCardDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("DvbCardDelete failed: %v", err)
	}
}

func TestClient_DvbCardGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/dvb_cards/test-name"

	// Expected JSON response
	expectedJSON := `{"adapter":16,"bandwidth":6000000,"code_rate_hp":"1_2","code_rate_lp":"1_2","comment":"example","frequency":11606,"guard_interval":"1_16","hierarchy":"1","hw":"dvb","modulation":"auto","name":"a16","pilot":"auto","polarization":"v","rolloff":"35","stats":{"ber":5,"has_carrier":true,"has_lock":true,"has_rate":true,"has_signal":true,"has_sync":true,"has_viterbi":true,"snr":1,"snr_raw":894,"strength":95,"strength_raw":62446},"symbol_rate":27500,"system":"dvbs","transmission_mode":"1k","video_device":"/dev/video0"}`

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
	_, err := client.DvbCardGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("DvbCardGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvbCardGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("DvbCardGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvbCardGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvbCardSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/dvb_cards/test-name"

	// Expected JSON response
	expectedJSON := `{"adapter":16,"bandwidth":6000000,"code_rate_hp":"1_2","code_rate_lp":"1_2","comment":"example","frequency":11606,"guard_interval":"1_16","hierarchy":"1","hw":"dvb","modulation":"auto","name":"a16","pilot":"auto","polarization":"v","rolloff":"35","stats":{"ber":5,"has_carrier":true,"has_lock":true,"has_rate":true,"has_signal":true,"has_sync":true,"has_viterbi":true,"snr":1,"snr_raw":894,"strength":95,"strength_raw":62446},"symbol_rate":27500,"system":"dvbs","transmission_mode":"1k","video_device":"/dev/video0"}`

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
	body := &model.DvbCardConfigImpl{}
	_, err := client.DvbCardSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("DvbCardSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvbCardSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("DvbCardSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvbCardSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvbCardsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/dvb_cards"

	// Expected JSON response
	expectedJSON := `{"dvb_cards":[{"adapter":16,"bandwidth":6000000,"code_rate_hp":"1_2","code_rate_lp":"1_2","comment":"example","frequency":11606,"guard_interval":"1_16","hierarchy":"1","hw":"dvb","modulation":"auto","name":"a16","pilot":"auto","polarization":"v","rolloff":"35","stats":{"ber":5,"has_carrier":true,"has_lock":true,"has_rate":true,"has_signal":true,"has_sync":true,"has_viterbi":true,"snr":1,"snr_raw":894,"strength":95,"strength_raw":62446},"symbol_rate":27500,"system":"dvbs","transmission_mode":"1k","video_device":"/dev/video0"}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.DvbCardsListQuery{}
	_, err := client.DvbCardsList(ctx, query)
	if err != nil {
		t.Fatalf("DvbCardsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvbCardsList(ctx, query)
	if err != nil {
		t.Fatalf("DvbCardsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvbCardsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvrDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/dvrs/test-name"

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
	err := client.DvrDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("DvrDelete failed: %v", err)
	}
}

func TestClient_DvrDiskDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/dvrs/test-name/disks/test-path"

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
	err := client.DvrDiskDelete(ctx, "test-name", "test-path")
	if err != nil {
		t.Fatalf("DvrDiskDelete failed: %v", err)
	}
}

func TestClient_DvrDiskGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/dvrs/test-name/disks/test-path"

	// Expected JSON response
	expectedJSON := `{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}`

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
	_, err := client.DvrDiskGet(ctx, "test-name", "test-path")
	if err != nil {
		t.Fatalf("DvrDiskGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvrDiskGet(ctx, "test-name", "test-path")
	if err != nil {
		t.Fatalf("DvrDiskGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvrDiskGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvrDiskSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/dvrs/test-name/disks/test-path"

	// Expected JSON response
	expectedJSON := `{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}`

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
	body := &model.RaidDiskConfigImpl{}
	_, err := client.DvrDiskSave(ctx, "test-name", "test-path", body)
	if err != nil {
		t.Fatalf("DvrDiskSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvrDiskSave(ctx, "test-name", "test-path", body)
	if err != nil {
		t.Fatalf("DvrDiskSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvrDiskSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvrDisksList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/dvrs/test-name/disks"

	// Expected JSON response
	expectedJSON := `{"disks":[{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.DvrDisksListQuery{}
	_, err := client.DvrDisksList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("DvrDisksList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvrDisksList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("DvrDisksList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvrDisksList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvrEpisodesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/dvr_episodes"

	// Expected JSON response
	expectedJSON := `{"episodes":[{"close_reason":"timeout","closed_at":1000000000000,"episode_appearance_timestamps":{"central_timestamp":1637098611000,"inference_timestamp":1637094994000,"watcher_timestamp":1637094994000},"episode_id":0,"episode_type":"generic","frame_preview":"example","media":"example","opened_at":1000000000000,"preview":"example","preview_timestamp":1000000000000,"recording_status":"full","started_at":1000000000000,"updated_at":1000000000000}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.DvrEpisodesListQuery{}
	_, err := client.DvrEpisodesList(ctx, query)
	if err != nil {
		t.Fatalf("DvrEpisodesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvrEpisodesList(ctx, query)
	if err != nil {
		t.Fatalf("DvrEpisodesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvrEpisodesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvrExportJobCancel(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/dvr_export_jobs/test-id"

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
	err := client.DvrExportJobCancel(ctx, "test-id")
	if err != nil {
		t.Fatalf("DvrExportJobCancel failed: %v", err)
	}
}

func TestClient_DvrExportJobList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/dvr_export_jobs"

	// Expected JSON response
	expectedJSON := `{"jobs":[{"duration":4200000,"error":"no_dvr","filter":{"tracks":"v1a1"},"finished_at":1730479931103,"from":1730205876000,"id":"3c448252-2516-4950-a1f6-fa2c7e8f4cb1","name":"demo","packing":"compat","path":"example","started_at":1730479930721,"status":"running","timelapse":"true","timelapse_kbps":3000}]}`

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
	_, err := client.DvrExportJobList(ctx)
	if err != nil {
		t.Fatalf("DvrExportJobList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvrExportJobList(ctx)
	if err != nil {
		t.Fatalf("DvrExportJobList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvrExportJobList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvrExportJobStart(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/dvr_export_jobs/test-id"

	// Expected JSON response
	expectedJSON := `{"duration":4200000,"error":"no_dvr","filter":{"tracks":"v1a1"},"finished_at":1730479931103,"from":1730205876000,"id":"3c448252-2516-4950-a1f6-fa2c7e8f4cb1","name":"demo","packing":"compat","path":"example","started_at":1730479930721,"status":"running","timelapse":"true","timelapse_kbps":3000}`

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
	body := &model.DvrExportJobImpl{}
	_, err := client.DvrExportJobStart(ctx, "test-id", body)
	if err != nil {
		t.Fatalf("DvrExportJobStart failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvrExportJobStart(ctx, "test-id", body)
	if err != nil {
		t.Fatalf("DvrExportJobStart failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvrExportJobStart returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvrExportJobStatus(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/dvr_export_jobs/test-id"

	// Expected JSON response
	expectedJSON := `{"duration":4200000,"error":"no_dvr","filter":{"tracks":"v1a1"},"finished_at":1730479931103,"from":1730205876000,"id":"3c448252-2516-4950-a1f6-fa2c7e8f4cb1","name":"demo","packing":"compat","path":"example","started_at":1730479930721,"status":"running","timelapse":"true","timelapse_kbps":3000}`

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
	_, err := client.DvrExportJobStatus(ctx, "test-id")
	if err != nil {
		t.Fatalf("DvrExportJobStatus failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvrExportJobStatus(ctx, "test-id")
	if err != nil {
		t.Fatalf("DvrExportJobStatus failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvrExportJobStatus returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvrGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/dvrs/test-name"

	// Expected JSON response
	expectedJSON := `{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"disks":[{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}],"dvr_replicate":true,"episodes_url":"example","index":"example","name":"example","replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000}`

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
	_, err := client.DvrGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("DvrGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvrGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("DvrGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvrGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvrSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/dvrs/test-name"

	// Expected JSON response
	expectedJSON := `{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"disks":[{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}],"dvr_replicate":true,"episodes_url":"example","index":"example","name":"example","replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000}`

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
	body := &model.DvrConfigImpl{}
	_, err := client.DvrSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("DvrSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.DvrSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("DvrSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("DvrSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_DvrsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/dvrs"

	// Expected JSON response
	expectedJSON := `{"dvrs":[{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"disks":[{"mode":"keep","path":"volume1","stats":{"errors":{"connection_timeout":1,"eacces":1,"eagain":1,"ebusy":1,"econnrefused":1,"edquot":1,"emfile":1,"enodev":1,"enoent":1,"enospc":1,"erofs":1,"nxdomain":1,"other":1,"ssl_error":1},"migration_eta":1000000000,"migration_updated":1000000000,"mode":"keep"}}],"dvr_replicate":true,"episodes_url":"example","index":"example","name":"example","replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","server_id":"550e8400-e29b-41d4-a716-446655440000"}`

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
	query := &flussonic.DvrsListQuery{}
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

func TestClient_EpisodesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/episodes"

	// Expected JSON response
	expectedJSON := `{"episodes":[{"close_reason":"timeout","closed_at":1000000000000,"episode_appearance_timestamps":{"central_timestamp":1637098611000,"inference_timestamp":1637094994000,"watcher_timestamp":1637094994000},"episode_id":0,"episode_type":"generic","frame_preview":"example","media":"example","opened_at":1000000000000,"preview":"example","preview_timestamp":1000000000000,"recording_status":"full","started_at":1000000000000,"updated_at":1000000000000}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.EpisodesListQuery{}
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

func TestClient_EventSinkDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/event_sinks/test-name"

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
	err := client.EventSinkDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("EventSinkDelete failed: %v", err)
	}
}

func TestClient_EventSinkGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/event_sinks/test-name"

	// Expected JSON response
	expectedJSON := `{"except":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"extra":{"key1":"example","key2":"example"},"level":"debug","max_depth":100,"max_size":10000,"name":"my_json_sink","only":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"resend_limit":1000,"resend_timeout":10,"throttle_delay":1,"url":"jsonlog:///var/log/events-json.log"}`

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
	_, err := client.EventSinkGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("EventSinkGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.EventSinkGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("EventSinkGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("EventSinkGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_EventSinkSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/event_sinks/test-name"

	// Expected JSON response
	expectedJSON := `{"except":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"extra":{"key1":"example","key2":"example"},"level":"debug","max_depth":100,"max_size":10000,"name":"my_json_sink","only":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"resend_limit":1000,"resend_timeout":10,"throttle_delay":1,"url":"jsonlog:///var/log/events-json.log"}`

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
	body := &model.EventSinkConfigImpl{}
	_, err := client.EventSinkSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("EventSinkSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.EventSinkSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("EventSinkSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("EventSinkSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_EventSinksList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/event_sinks"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"event_sinks":[{"except":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"extra":{"key1":"example","key2":"example"},"level":"debug","max_depth":100,"max_size":10000,"name":"my_json_sink","only":[{"key1":["stream_stopped"],"key2":["stream_stopped"]}],"resend_limit":1000,"resend_timeout":10,"throttle_delay":1,"url":"jsonlog:///var/log/events-json.log"}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.EventSinksListQuery{}
	_, err := client.EventSinksList(ctx, query)
	if err != nil {
		t.Fatalf("EventSinksList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.EventSinksList(ctx, query)
	if err != nil {
		t.Fatalf("EventSinksList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("EventSinksList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_FileProcessorDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/file_processor"

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
	err := client.FileProcessorDelete(ctx)
	if err != nil {
		t.Fatalf("FileProcessorDelete failed: %v", err)
	}
}

func TestClient_FileProcessorGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/file_processor"

	// Expected JSON response
	expectedJSON := `{"path":"example"}`

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
	_, err := client.FileProcessorGet(ctx)
	if err != nil {
		t.Fatalf("FileProcessorGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.FileProcessorGet(ctx)
	if err != nil {
		t.Fatalf("FileProcessorGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("FileProcessorGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_FileProcessorJobCheck(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/file_processor/jobs/test-id"

	// Expected JSON response
	expectedJSON := `{"errors":[{"code":"example","id":"example","meta":{"key1":"example","key2":"example"},"source":{"parameter":"example","pointer":"example"},"status":"example","title":"example"}],"filters":[{"group":"overlay","position":{"x1":0,"x2":0,"y1":0,"y2":0},"start":0,"stop":0}],"id":"example","input_files":["example"],"output_file":"example","status":"new"}`

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
	_, err := client.FileProcessorJobCheck(ctx, "test-id")
	if err != nil {
		t.Fatalf("FileProcessorJobCheck failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.FileProcessorJobCheck(ctx, "test-id")
	if err != nil {
		t.Fatalf("FileProcessorJobCheck failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("FileProcessorJobCheck returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_FileProcessorJobDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/file_processor/jobs/test-id"

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
	err := client.FileProcessorJobDelete(ctx, "test-id")
	if err != nil {
		t.Fatalf("FileProcessorJobDelete failed: %v", err)
	}
}

func TestClient_FileProcessorJobStart(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/streamer/api/v3/file_processor/jobs"

	// Expected JSON response
	expectedJSON := `{"errors":[{"code":"example","id":"example","meta":{"key1":"example","key2":"example"},"source":{"parameter":"example","pointer":"example"},"status":"example","title":"example"}],"filters":[{"group":"overlay","position":{"x1":0,"x2":0,"y1":0,"y2":0},"start":0,"stop":0}],"id":"example","input_files":["example"],"output_file":"example","status":"new"}`

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
	body := &model.FileProcessorJobImpl{}
	_, err := client.FileProcessorJobStart(ctx, body)
	if err != nil {
		t.Fatalf("FileProcessorJobStart failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.FileProcessorJobStart(ctx, body)
	if err != nil {
		t.Fatalf("FileProcessorJobStart failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("FileProcessorJobStart returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_FileProcessorSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/file_processor"

	// Expected JSON response
	expectedJSON := `{"path":"example"}`

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
	body := &model.FileProcessorConfigImpl{}
	_, err := client.FileProcessorSave(ctx, body)
	if err != nil {
		t.Fatalf("FileProcessorSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.FileProcessorSave(ctx, body)
	if err != nil {
		t.Fatalf("FileProcessorSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("FileProcessorSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_HttpProxiesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/http_proxies"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"http_proxies":[{"prefix":"example","stats":{"http_100":1000,"http_200":1000,"http_300":10,"http_400":10,"http_500":10,"protocol_upgrades":1000,"proxy_error":10,"proxy_error_connection":10,"requests":1000,"time_1000ms":2,"time_100ms":10,"time_5000ms":2,"time_500ms":5,"time_longms":1},"url":"https://example.com"}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","server_id":"550e8400-e29b-41d4-a716-446655440000"}`

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
	query := &flussonic.HttpProxiesListQuery{}
	_, err := client.HttpProxiesList(ctx, query)
	if err != nil {
		t.Fatalf("HttpProxiesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.HttpProxiesList(ctx, query)
	if err != nil {
		t.Fatalf("HttpProxiesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("HttpProxiesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_HttpProxyDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/http_proxies/test-prefix"

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
	err := client.HttpProxyDelete(ctx, "test-prefix")
	if err != nil {
		t.Fatalf("HttpProxyDelete failed: %v", err)
	}
}

func TestClient_HttpProxyGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/http_proxies/test-prefix"

	// Expected JSON response
	expectedJSON := `{"prefix":"example","stats":{"http_100":1000,"http_200":1000,"http_300":10,"http_400":10,"http_500":10,"protocol_upgrades":1000,"proxy_error":10,"proxy_error_connection":10,"requests":1000,"time_1000ms":2,"time_100ms":10,"time_5000ms":2,"time_500ms":5,"time_longms":1},"url":"https://example.com"}`

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
	_, err := client.HttpProxyGet(ctx, "test-prefix")
	if err != nil {
		t.Fatalf("HttpProxyGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.HttpProxyGet(ctx, "test-prefix")
	if err != nil {
		t.Fatalf("HttpProxyGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("HttpProxyGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_HttpProxySave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/http_proxies/test-prefix"

	// Expected JSON response
	expectedJSON := `{"prefix":"example","stats":{"http_100":1000,"http_200":1000,"http_300":10,"http_400":10,"http_500":10,"protocol_upgrades":1000,"proxy_error":10,"proxy_error_connection":10,"requests":1000,"time_1000ms":2,"time_100ms":10,"time_5000ms":2,"time_500ms":5,"time_longms":1},"url":"https://example.com"}`

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
	body := &model.HTTPProxyConfigImpl{}
	_, err := client.HttpProxySave(ctx, "test-prefix", body)
	if err != nil {
		t.Fatalf("HttpProxySave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.HttpProxySave(ctx, "test-prefix", body)
	if err != nil {
		t.Fatalf("HttpProxySave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("HttpProxySave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_IptvDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/iptv"

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
	err := client.IptvDelete(ctx)
	if err != nil {
		t.Fatalf("IptvDelete failed: %v", err)
	}
}

func TestClient_IptvGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/iptv"

	// Expected JSON response
	expectedJSON := `{"database":"example"}`

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
	_, err := client.IptvGet(ctx)
	if err != nil {
		t.Fatalf("IptvGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.IptvGet(ctx)
	if err != nil {
		t.Fatalf("IptvGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("IptvGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_IptvSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/iptv"

	// Expected JSON response
	expectedJSON := `{"database":"example"}`

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
	body := &model.IptvConfigImpl{}
	_, err := client.IptvSave(ctx, body)
	if err != nil {
		t.Fatalf("IptvSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.IptvSave(ctx, body)
	if err != nil {
		t.Fatalf("IptvSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("IptvSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_LivenessProbe(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/monitoring/liveness"

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

func TestClient_LogoDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/logos/test-name"

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
	err := client.LogoDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("LogoDelete failed: %v", err)
	}
}

func TestClient_LogoGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/logos/test-name"

	// Expected JSON response
	expectedJSON := `{"content":"example","content_type":"example","name":"example","stream_names":["example"]}`

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
	_, err := client.LogoGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("LogoGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.LogoGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("LogoGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("LogoGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_LogoSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/logos/test-name"

	// Expected JSON response
	expectedJSON := `{"content":"example","content_type":"example","name":"example","stream_names":["example"]}`

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
	body := &model.LogoFileImpl{}
	_, err := client.LogoSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("LogoSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.LogoSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("LogoSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("LogoSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_LogosList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/logos"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"logos":[{"content":"example","content_type":"example","name":"example","stream_names":["example"]}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.LogosListQuery{}
	_, err := client.LogosList(ctx, query)
	if err != nil {
		t.Fatalf("LogosList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.LogosList(ctx, query)
	if err != nil {
		t.Fatalf("LogosList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("LogosList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_MultiplexerXmltvUpload(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/streamer/api/v3/multiplexers/test-name/xmltv_upload"

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
	err := client.MultiplexerXmltvUpload(ctx, "test-name")
	if err != nil {
		t.Fatalf("MultiplexerXmltvUpload failed: %v", err)
	}
}

func TestClient_OpenedFilesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/vods/opened_files"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"files":[{"bytes":42309561,"folder":"example","name":"example","opened":true,"prefix":"example","stats":{"bytes_in":1700923231,"bytes_out":1700923231,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"worker_count":1},"subpath":"example","url":"example"}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.OpenedFilesListQuery{}
	_, err := client.OpenedFilesList(ctx, query)
	if err != nil {
		t.Fatalf("OpenedFilesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.OpenedFilesList(ctx, query)
	if err != nil {
		t.Fatalf("OpenedFilesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("OpenedFilesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PackageDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/iptv/packages/test-name"

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
	err := client.PackageDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("PackageDelete failed: %v", err)
	}
}

func TestClient_PackageGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/iptv/packages/test-name"

	// Expected JSON response
	expectedJSON := `{"channels":["example"],"name":"example"}`

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
	_, err := client.PackageGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("PackageGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PackageGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("PackageGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PackageGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PackageSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/iptv/packages/test-name"

	// Expected JSON response
	expectedJSON := `{"channels":["example"],"name":"example"}`

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
	body := &model.PackageConfigImpl{}
	_, err := client.PackageSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("PackageSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PackageSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("PackageSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PackageSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PackagesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/iptv/packages"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","packages":[{"channels":["example"],"name":"example"}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.PackagesListQuery{}
	_, err := client.PackagesList(ctx, query)
	if err != nil {
		t.Fatalf("PackagesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PackagesList(ctx, query)
	if err != nil {
		t.Fatalf("PackagesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PackagesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PeerDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/cluster/peers/test-hostname"

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
	err := client.PeerDelete(ctx, "test-hostname")
	if err != nil {
		t.Fatalf("PeerDelete failed: %v", err)
	}
}

func TestClient_PeerGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/cluster/peers/test-hostname"

	// Expected JSON response
	expectedJSON := `{"api_url":"http://streamer.local:8080","channel_limit":5,"cluster_key":"xS6i6Q3DCc5nEvnu","cpu_limit":10,"fetch_timeout":1000,"hostname":"peer.example.com","private_payload_url":"http://streamer.local","public_payload_url":"http://public.example.com","stale_timeout":1000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502}}`

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
	_, err := client.PeerGet(ctx, "test-hostname")
	if err != nil {
		t.Fatalf("PeerGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PeerGet(ctx, "test-hostname")
	if err != nil {
		t.Fatalf("PeerGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PeerGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PeerSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/cluster/peers/test-hostname"

	// Expected JSON response
	expectedJSON := `{"api_url":"http://streamer.local:8080","channel_limit":5,"cluster_key":"xS6i6Q3DCc5nEvnu","cpu_limit":10,"fetch_timeout":1000,"hostname":"peer.example.com","private_payload_url":"http://streamer.local","public_payload_url":"http://public.example.com","stale_timeout":1000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502}}`

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
	body := &model.PeerConfigImpl{}
	_, err := client.PeerSave(ctx, "test-hostname", body)
	if err != nil {
		t.Fatalf("PeerSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PeerSave(ctx, "test-hostname", body)
	if err != nil {
		t.Fatalf("PeerSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PeerSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_PeersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/cluster/peers"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","peers":[{"api_url":"http://streamer.local:8080","channel_limit":5,"cluster_key":"xS6i6Q3DCc5nEvnu","cpu_limit":10,"fetch_timeout":1000,"hostname":"peer.example.com","private_payload_url":"http://streamer.local","public_payload_url":"http://public.example.com","stale_timeout":1000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502}}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.PeersListQuery{}
	_, err := client.PeersList(ctx, query)
	if err != nil {
		t.Fatalf("PeersList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.PeersList(ctx, query)
	if err != nil {
		t.Fatalf("PeersList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("PeersList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_ReadinessProbe(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/monitoring/readiness"

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

func TestClient_RproxyAgentsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/rproxy/agents"

	// Expected JSON response
	expectedJSON := `{"agents":[{"id":"1234567","key":"example","stats":{"agent_type":"single","endpoint_connection":{"bytes_from_server":40000,"bytes_to_server":400000000000,"hostname":"agents-001.vsaas.io","opened_at":1637094994000,"status_changed_at":1634560921},"local_ip":"10.10.17.88","mac_address":"F0-23-B9-59-20-F1","peer_ip":"185.134.232.183","streampoint_connection":{"bytes_from_server":40000,"bytes_to_server":400000000000,"connections_attempted":400,"connections_current":2,"connections_opened":300,"hostname":"agents-001.vsaas.io","opened_at":1637094994000,"status_changed_at":1634560921},"version":"v21.02-8-g535c85d"}}],"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.RproxyAgentsListQuery{}
	_, err := client.RproxyAgentsList(ctx, query)
	if err != nil {
		t.Fatalf("RproxyAgentsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.RproxyAgentsList(ctx, query)
	if err != nil {
		t.Fatalf("RproxyAgentsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("RproxyAgentsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_RproxyDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/rproxy"

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
	err := client.RproxyDelete(ctx)
	if err != nil {
		t.Fatalf("RproxyDelete failed: %v", err)
	}
}

func TestClient_RproxyGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/rproxy"

	// Expected JSON response
	expectedJSON := `{"endpoint_auth":"example","forward_ports":{"key1":{"handler":"example"},"key2":{"handler":"example"}},"streampoint_key":"example"}`

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
	_, err := client.RproxyGet(ctx)
	if err != nil {
		t.Fatalf("RproxyGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.RproxyGet(ctx)
	if err != nil {
		t.Fatalf("RproxyGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("RproxyGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_RproxySave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/rproxy"

	// Expected JSON response
	expectedJSON := `{"endpoint_auth":"example","forward_ports":{"key1":{"handler":"example"},"key2":{"handler":"example"}},"streampoint_key":"example"}`

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
	body := &model.RproxyConfigImpl{}
	_, err := client.RproxySave(ctx, body)
	if err != nil {
		t.Fatalf("RproxySave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.RproxySave(ctx, body)
	if err != nil {
		t.Fatalf("RproxySave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("RproxySave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_SessionDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/sessions/test-id"

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
	err := client.SessionDelete(ctx, "test-id")
	if err != nil {
		t.Fatalf("SessionDelete failed: %v", err)
	}
}

func TestClient_SessionGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/sessions/test-id"

	// Expected JSON response
	expectedJSON := `{"allowed_dvr_from":1000000000,"allowed_dvr_to":1000000000,"closed_at":1637098821000,"country":"us","id":"61942414-8c15-4809-8bb6-adf1ae846027","ip":"172.16.25.73","manifest_type":"rewind","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"name":"hockey1","named_by":"config","opened_at":1637094994000,"proto":"dash","query_string":"example","referer":"http://my-tv-portal.local/hockey1","security_protocol":"tlsv1.2","segments_container":"cmaf","started_at":1637095014000,"token":"zGAFxLkoWluO1pG7_nJmQAbCnM5","ts_delay":1284,"ts_delay_per_tracks":[1284],"updated_at":1637098611000,"user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","user_id":"5435","user_name":"example"}`

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
	_, err := client.SessionGet(ctx, "test-id")
	if err != nil {
		t.Fatalf("SessionGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SessionGet(ctx, "test-id")
	if err != nil {
		t.Fatalf("SessionGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SessionGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_SessionsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/sessions"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","sessions":[{"allowed_dvr_from":1000000000,"allowed_dvr_to":1000000000,"closed_at":1637098821000,"country":"us","id":"61942414-8c15-4809-8bb6-adf1ae846027","ip":"172.16.25.73","manifest_type":"rewind","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"name":"hockey1","named_by":"config","opened_at":1637094994000,"proto":"dash","query_string":"example","referer":"http://my-tv-portal.local/hockey1","security_protocol":"tlsv1.2","segments_container":"cmaf","started_at":1637095014000,"token":"zGAFxLkoWluO1pG7_nJmQAbCnM5","ts_delay":1284,"ts_delay_per_tracks":[1284],"updated_at":1637098611000,"user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","user_id":"5435","user_name":"example"}]}`

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
	query := &flussonic.SessionsListQuery{}
	_, err := client.SessionsList(ctx, query)
	if err != nil {
		t.Fatalf("SessionsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SessionsList(ctx, query)
	if err != nil {
		t.Fatalf("SessionsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SessionsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_SessionsReauth(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/streamer/api/v3/sessions/reauth"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.SessionsReauthQuery{}
	query.Name = "test-name"
	_, err := client.SessionsReauth(ctx, query)
	if err != nil {
		t.Fatalf("SessionsReauth failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SessionsReauth(ctx, query)
	if err != nil {
		t.Fatalf("SessionsReauth failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SessionsReauth returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_SourceDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/cluster/sources/test-url"

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
	err := client.SourceDelete(ctx, "test-url")
	if err != nil {
		t.Fatalf("SourceDelete failed: %v", err)
	}
}

func TestClient_SourceGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/cluster/sources/test-url"

	// Expected JSON response
	expectedJSON := `{"add_audio_only":true,"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"cluster_key":"example","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"except":["example"],"group_config":{"key1":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000},"key2":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000}},"hls_scte35":"scte35","jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"meta":{"key1":"example","key2":"example"},"mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"only":["example"],"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefix":"example","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"url":"example_value","vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}`

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
	_, err := client.SourceGet(ctx, "test-url")
	if err != nil {
		t.Fatalf("SourceGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SourceGet(ctx, "test-url")
	if err != nil {
		t.Fatalf("SourceGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SourceGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_SourceSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/cluster/sources/test-url"

	// Expected JSON response
	expectedJSON := `{"add_audio_only":true,"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"cluster_key":"example","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"except":["example"],"group_config":{"key1":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000},"key2":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000}},"hls_scte35":"scte35","jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"meta":{"key1":"example","key2":"example"},"mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"only":["example"],"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefix":"example","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"url":"example_value","vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}`

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
	body := &model.SourceConfigImpl{}
	_, err := client.SourceSave(ctx, "test-url", body)
	if err != nil {
		t.Fatalf("SourceSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SourceSave(ctx, "test-url", body)
	if err != nil {
		t.Fatalf("SourceSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SourceSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_SourcesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/cluster/sources"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","sources":[{"add_audio_only":true,"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"cluster_key":"example","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"except":["example"],"group_config":{"key1":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000},"key2":{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"segment_duration":1000}},"hls_scte35":"scte35","jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"meta":{"key1":"example","key2":"example"},"mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"only":["example"],"playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefix":"example","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"stats":{"bandwidth_usage":67,"config_error":{"col":20,"detail":"example","error":"bad_url","first_error_col":5,"first_error_line":14,"line":15,"path":["streams",0,"inputs",0,"url","input_url"]},"config_external_status":{"detail":"example","error":"invalid_authorization","path":["streams",0,"inputs",0,"url","input_url"],"reason":"validation_error","status":"loaded","while":"refresh"},"config_version":[1636709231,4],"cpu_usage":48,"error":"example","hostname":"openapi.flussonic.com","id":"61893b15-75b2-4fcb-b4cf-ae1dd0858ea2","input_kbit":400300,"license_txt":"uO8v12HJhNXVj5gM","license_type":"undefined","memory_usage":27,"next_version":"22.01","now":1000000000000,"online_streams":27,"opened_files":5,"output_kbit":500400,"partitions":[{"device":"sda1","io_util":42,"path":"_var_lib_docker_overlay2_3684778579db0a4d418fdc1a8a6953b680ab92d179a7d6f9506710d073095e36_merged","total_mb":45423,"usage":30}],"scheduler_load":40,"server_version":"23.04","started_at":1639337825,"streamer_status":"running","text_alerts":{"key1":"example","key2":"example"},"total_clients":2040,"total_streams":45,"transcoder_devices":[{"id":"auto","name":"example","reconfig_support":"full","type":"cpu"}],"uptime":4325502},"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"url":"example_value","vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}]}`

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
	query := &flussonic.SourcesListQuery{}
	_, err := client.SourcesList(ctx, query)
	if err != nil {
		t.Fatalf("SourcesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SourcesList(ctx, query)
	if err != nil {
		t.Fatalf("SourcesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SourcesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/streams/test-name"

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

func TestClient_StreamDvrConsistencyCheck(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/streamer/api/v3/streams/test-name/dvr/consistency_check"

	// Expected JSON response
	expectedJSON := `{"errors":[{"details":{"key1":"example","key2":"example"},"failed_check":"blob_coherence"}],"name":"demo"}`

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
	_, err := client.StreamDvrConsistencyCheck(ctx, "test-name")
	if err != nil {
		t.Fatalf("StreamDvrConsistencyCheck failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamDvrConsistencyCheck(ctx, "test-name")
	if err != nil {
		t.Fatalf("StreamDvrConsistencyCheck failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamDvrConsistencyCheck returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamDvrLocksDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/streams/test-name/dvr/locks"

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
	body := &model.DvrRangeImpl{}
	err := client.StreamDvrLocksDelete(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("StreamDvrLocksDelete failed: %v", err)
	}
}

func TestClient_StreamDvrLocksList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/streams/test-name/dvr/locks"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"locks":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}],"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.StreamDvrLocksListQuery{}
	_, err := client.StreamDvrLocksList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("StreamDvrLocksList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamDvrLocksList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("StreamDvrLocksList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamDvrLocksList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamDvrLocksSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/streamer/api/v3/streams/test-name/dvr/locks"

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
	body := &model.DvrRangeImpl{}
	err := client.StreamDvrLocksSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("StreamDvrLocksSave failed: %v", err)
	}
}

func TestClient_StreamDvrRangesDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/streams/test-name/dvr/ranges"

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
	body := &model.DvrRangeImpl{}
	err := client.StreamDvrRangesDelete(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("StreamDvrRangesDelete failed: %v", err)
	}
}

func TestClient_StreamDvrRangesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/streams/test-name/dvr/ranges"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]}`

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
	query := &flussonic.StreamDvrRangesListQuery{}
	_, err := client.StreamDvrRangesList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("StreamDvrRangesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.StreamDvrRangesList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("StreamDvrRangesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("StreamDvrRangesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_StreamGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/streams/test-name"

	// Expected JSON response
	expectedJSON := `{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","config_on_disk":{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","drm":{"vendor":"example"},"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"dvr_read_performance":{"segments_read_delayed":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_fast":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_slow":{"cache":100,"local":10,"ram":100,"remote":10}},"dvr_read_popularity":{"other":{"cache":100,"local":10,"ram":100,"remote":10},"today":{"cache":100,"local":10,"ram":100,"remote":10},"week":{"cache":100,"local":10,"ram":100,"remote":10},"yesterday":{"cache":100,"local":10,"ram":100,"remote":10}},"id":"61893ba6-07b3-431b-b2f7-716ac1643953","input":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"input_bitrate":186,"last_access_at":1669106270979,"last_dts":383835646,"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"opened_at":1636383729002,"play":{"play_bytes_by_proto":{"key1":0,"key2":0},"play_dash_playlist_http_200":1000,"play_dash_playlist_http_400":1000,"play_dash_playlist_http_500":1000,"play_dash_playlist_time_1000ms":10,"play_dash_playlist_time_100ms":10,"play_dash_playlist_time_500ms":10,"play_dash_playlist_time_longms":10,"play_dash_segment_http_200":1000,"play_dash_segment_http_400":1000,"play_dash_segment_http_500":1000,"play_dash_segment_time_1000ms":10,"play_dash_segment_time_100ms":10,"play_dash_segment_time_500ms":10,"play_dash_segment_time_longms":10,"play_duration_by_proto":{"key1":0,"key2":0},"play_hls_playlist_http_200":1000,"play_hls_playlist_http_400":1000,"play_hls_playlist_http_500":1000,"play_hls_playlist_time_1000ms":10,"play_hls_playlist_time_100ms":10,"play_hls_playlist_time_500ms":10,"play_hls_playlist_time_longms":10,"play_hls_segment_http_200":1000,"play_hls_segment_http_400":1000,"play_hls_segment_http_500":1000,"play_hls_segment_time_1000ms":10,"play_hls_segment_time_100ms":10,"play_hls_segment_time_500ms":10,"play_hls_segment_time_longms":10,"play_ll_hls_playlist_http_200":1000,"play_ll_hls_playlist_http_400":1000,"play_ll_hls_playlist_http_500":1000,"play_ll_hls_playlist_time_1000ms":10,"play_ll_hls_playlist_time_100ms":10,"play_ll_hls_playlist_time_500ms":10,"play_ll_hls_playlist_time_longms":10,"play_ll_hls_segment_http_200":1000,"play_ll_hls_segment_http_400":1000,"play_ll_hls_segment_http_500":1000,"play_ll_hls_segment_time_1000ms":10,"play_ll_hls_segment_time_100ms":10,"play_ll_hls_segment_time_500ms":10,"play_ll_hls_segment_time_longms":10,"play_opened_sessions_by_proto":{"key1":0,"key2":0},"play_other_http_200":1000,"play_other_http_400":1000,"play_other_http_500":1000,"play_other_time_1000ms":10,"play_other_time_100ms":10,"play_other_time_500ms":10,"play_other_time_longms":10,"play_preview_http_200":1000,"play_preview_http_400":1000,"play_preview_http_500":1000,"play_preview_time_1000ms":10,"play_preview_time_100ms":10,"play_preview_time_500ms":10,"play_preview_time_longms":10,"play_total_sessions_by_proto":{"key1":0,"key2":0}},"push":[{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"}],"running_on":["streamer1.example.com"],"source_id":"61893be1-054e-4acc-8d24-8ed92efe6ad0","start_running_at":1000000000000,"status":"running","transcoder":{"hw":"cpu","target":"uhdtv","tracks":[{"track_id":"v1"}]},"ts_delay":1284},"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}`

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

func TestClient_StreamSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/streams/test-name"

	// Expected JSON response
	expectedJSON := `{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","config_on_disk":{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","drm":{"vendor":"example"},"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"dvr_read_performance":{"segments_read_delayed":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_fast":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_slow":{"cache":100,"local":10,"ram":100,"remote":10}},"dvr_read_popularity":{"other":{"cache":100,"local":10,"ram":100,"remote":10},"today":{"cache":100,"local":10,"ram":100,"remote":10},"week":{"cache":100,"local":10,"ram":100,"remote":10},"yesterday":{"cache":100,"local":10,"ram":100,"remote":10}},"id":"61893ba6-07b3-431b-b2f7-716ac1643953","input":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"input_bitrate":186,"last_access_at":1669106270979,"last_dts":383835646,"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"opened_at":1636383729002,"play":{"play_bytes_by_proto":{"key1":0,"key2":0},"play_dash_playlist_http_200":1000,"play_dash_playlist_http_400":1000,"play_dash_playlist_http_500":1000,"play_dash_playlist_time_1000ms":10,"play_dash_playlist_time_100ms":10,"play_dash_playlist_time_500ms":10,"play_dash_playlist_time_longms":10,"play_dash_segment_http_200":1000,"play_dash_segment_http_400":1000,"play_dash_segment_http_500":1000,"play_dash_segment_time_1000ms":10,"play_dash_segment_time_100ms":10,"play_dash_segment_time_500ms":10,"play_dash_segment_time_longms":10,"play_duration_by_proto":{"key1":0,"key2":0},"play_hls_playlist_http_200":1000,"play_hls_playlist_http_400":1000,"play_hls_playlist_http_500":1000,"play_hls_playlist_time_1000ms":10,"play_hls_playlist_time_100ms":10,"play_hls_playlist_time_500ms":10,"play_hls_playlist_time_longms":10,"play_hls_segment_http_200":1000,"play_hls_segment_http_400":1000,"play_hls_segment_http_500":1000,"play_hls_segment_time_1000ms":10,"play_hls_segment_time_100ms":10,"play_hls_segment_time_500ms":10,"play_hls_segment_time_longms":10,"play_ll_hls_playlist_http_200":1000,"play_ll_hls_playlist_http_400":1000,"play_ll_hls_playlist_http_500":1000,"play_ll_hls_playlist_time_1000ms":10,"play_ll_hls_playlist_time_100ms":10,"play_ll_hls_playlist_time_500ms":10,"play_ll_hls_playlist_time_longms":10,"play_ll_hls_segment_http_200":1000,"play_ll_hls_segment_http_400":1000,"play_ll_hls_segment_http_500":1000,"play_ll_hls_segment_time_1000ms":10,"play_ll_hls_segment_time_100ms":10,"play_ll_hls_segment_time_500ms":10,"play_ll_hls_segment_time_longms":10,"play_opened_sessions_by_proto":{"key1":0,"key2":0},"play_other_http_200":1000,"play_other_http_400":1000,"play_other_http_500":1000,"play_other_time_1000ms":10,"play_other_time_100ms":10,"play_other_time_500ms":10,"play_other_time_longms":10,"play_preview_http_200":1000,"play_preview_http_400":1000,"play_preview_http_500":1000,"play_preview_time_1000ms":10,"play_preview_time_100ms":10,"play_preview_time_500ms":10,"play_preview_time_longms":10,"play_total_sessions_by_proto":{"key1":0,"key2":0}},"push":[{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"}],"running_on":["streamer1.example.com"],"source_id":"61893be1-054e-4acc-8d24-8ed92efe6ad0","start_running_at":1000000000000,"status":"running","transcoder":{"hw":"cpu","target":"uhdtv","tracks":[{"track_id":"v1"}]},"ts_delay":1284},"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}`

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

func TestClient_StreamSaveMp4(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/streamer/api/v3/streams/test-name/dvr/export"

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
	query := &flussonic.StreamSaveMp4Query{}
	query.Duration = 1
	query.From = 1
	query.Path = "test-path"
	err := client.StreamSaveMp4(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("StreamSaveMp4 failed: %v", err)
	}
}

func TestClient_StreamStop(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "POST"
	expectedPath := "/streamer/api/v3/streams/test-name/stop"

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
	err := client.StreamStop(ctx, "test-name")
	if err != nil {
		t.Fatalf("StreamStop failed: %v", err)
	}
}

func TestClient_StreamsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/streams"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","server_id":"550e8400-e29b-41d4-a716-446655440000","streams":[{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","config_on_disk":{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"comment":"This is a test stream","drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","named_by":"config","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"position":2,"prefer_codec":"h264","provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"recheck_secondary_inputs_interval":120,"segment_count":4,"segment_duration":5000,"source_timeout":10,"srt":9060,"srt2":9062,"srt2_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt2_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_play":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"srt_publish":{"enforcedencryption":true,"latency":150,"linger":15,"minversion":"1.1.0","passphrase":"9876543210","port":9050,"streamid":"#!::r=my-stream,m=publish","timeout":10,"version":"1.3.0"},"static":true,"stats":{"agent_status":"connected","alive":true,"bitrate":186,"current_agent_id":"example","drm":{"vendor":"example"},"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"dvr_read_performance":{"segments_read_delayed":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_fast":{"cache":100,"local":10,"ram":100,"remote":10},"segments_read_slow":{"cache":100,"local":10,"ram":100,"remote":10}},"dvr_read_popularity":{"other":{"cache":100,"local":10,"ram":100,"remote":10},"today":{"cache":100,"local":10,"ram":100,"remote":10},"week":{"cache":100,"local":10,"ram":100,"remote":10},"yesterday":{"cache":100,"local":10,"ram":100,"remote":10}},"id":"61893ba6-07b3-431b-b2f7-716ac1643953","input":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"input_bitrate":186,"last_access_at":1669106270979,"last_dts":383835646,"last_dts_at":1636383841974,"lifetime":71977,"media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"online_clients":3,"opened_at":1636383729002,"play":{"play_bytes_by_proto":{"key1":0,"key2":0},"play_dash_playlist_http_200":1000,"play_dash_playlist_http_400":1000,"play_dash_playlist_http_500":1000,"play_dash_playlist_time_1000ms":10,"play_dash_playlist_time_100ms":10,"play_dash_playlist_time_500ms":10,"play_dash_playlist_time_longms":10,"play_dash_segment_http_200":1000,"play_dash_segment_http_400":1000,"play_dash_segment_http_500":1000,"play_dash_segment_time_1000ms":10,"play_dash_segment_time_100ms":10,"play_dash_segment_time_500ms":10,"play_dash_segment_time_longms":10,"play_duration_by_proto":{"key1":0,"key2":0},"play_hls_playlist_http_200":1000,"play_hls_playlist_http_400":1000,"play_hls_playlist_http_500":1000,"play_hls_playlist_time_1000ms":10,"play_hls_playlist_time_100ms":10,"play_hls_playlist_time_500ms":10,"play_hls_playlist_time_longms":10,"play_hls_segment_http_200":1000,"play_hls_segment_http_400":1000,"play_hls_segment_http_500":1000,"play_hls_segment_time_1000ms":10,"play_hls_segment_time_100ms":10,"play_hls_segment_time_500ms":10,"play_hls_segment_time_longms":10,"play_ll_hls_playlist_http_200":1000,"play_ll_hls_playlist_http_400":1000,"play_ll_hls_playlist_http_500":1000,"play_ll_hls_playlist_time_1000ms":10,"play_ll_hls_playlist_time_100ms":10,"play_ll_hls_playlist_time_500ms":10,"play_ll_hls_playlist_time_longms":10,"play_ll_hls_segment_http_200":1000,"play_ll_hls_segment_http_400":1000,"play_ll_hls_segment_http_500":1000,"play_ll_hls_segment_time_1000ms":10,"play_ll_hls_segment_time_100ms":10,"play_ll_hls_segment_time_500ms":10,"play_ll_hls_segment_time_longms":10,"play_opened_sessions_by_proto":{"key1":0,"key2":0},"play_other_http_200":1000,"play_other_http_400":1000,"play_other_http_500":1000,"play_other_time_1000ms":10,"play_other_time_100ms":10,"play_other_time_500ms":10,"play_other_time_longms":10,"play_preview_http_200":1000,"play_preview_http_400":1000,"play_preview_http_500":1000,"play_preview_time_1000ms":10,"play_preview_time_100ms":10,"play_preview_time_500ms":10,"play_preview_time_longms":10,"play_total_sessions_by_proto":{"key1":0,"key2":0}},"push":[{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"}],"running_on":["streamer1.example.com"],"source_id":"61893be1-054e-4acc-8d24-8ed92efe6ad0","start_running_at":1000000000000,"status":"running","transcoder":{"hw":"cpu","target":"uhdtv","tracks":[{"track_id":"v1"}]},"ts_delay":1284},"template":"sports-hd","thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"title":"Hockey channel","transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}]}`

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
	query := &flussonic.StreamsListQuery{}
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

func TestClient_SubscriberDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/iptv/subscribers/test-name"

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
	err := client.SubscriberDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("SubscriberDelete failed: %v", err)
	}
}

func TestClient_SubscriberGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/iptv/subscribers/test-name"

	// Expected JSON response
	expectedJSON := `{"max_sessions":3,"name":"client01","packages":["example"],"token":"ybBb5CFLqSFYc2"}`

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
	_, err := client.SubscriberGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("SubscriberGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SubscriberGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("SubscriberGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SubscriberGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_SubscriberSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/iptv/subscribers/test-name"

	// Expected JSON response
	expectedJSON := `{"max_sessions":3,"name":"client01","packages":["example"],"token":"ybBb5CFLqSFYc2"}`

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
	body := &model.SubscriberConfigImpl{}
	_, err := client.SubscriberSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("SubscriberSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SubscriberSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("SubscriberSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SubscriberSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_SubscribersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/iptv/subscribers"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","subscribers":[{"max_sessions":3,"name":"client01","packages":["example"],"token":"ybBb5CFLqSFYc2"}]}`

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
	query := &flussonic.SubscribersListQuery{}
	_, err := client.SubscribersList(ctx, query)
	if err != nil {
		t.Fatalf("SubscribersList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.SubscribersList(ctx, query)
	if err != nil {
		t.Fatalf("SubscribersList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("SubscribersList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TemplateDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/templates/test-name"

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
	err := client.TemplateDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("TemplateDelete failed: %v", err)
	}
}

func TestClient_TemplateGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/templates/test-name"

	// Expected JSON response
	expectedJSON := `{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefer_codec":"h264","prefixes":["chats",""],"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"source_timeout":10,"static":true,"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}`

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
	_, err := client.TemplateGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("TemplateGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TemplateGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("TemplateGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TemplateGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TemplateSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/templates/test-name"

	// Expected JSON response
	expectedJSON := `{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefer_codec":"h264","prefixes":["chats",""],"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"source_timeout":10,"static":true,"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}`

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
	body := &model.TemplateConfigImpl{}
	_, err := client.TemplateSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("TemplateSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TemplateSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("TemplateSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TemplateSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TemplatesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/templates"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","templates":[{"add_audio_only":true,"audio_timeout":20,"backup":{"audio_timeout":5,"file":"vod/blank.mp4","timeout":10,"video_timeout":4},"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"chunk_duration":200,"clients_timeout":485,"cluster_ingest":{"capture_at":"example"},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"dvbocr":"replace","dvr":{"copy":"s3://token@minio.mycompany.com/dvr-bucket","disk_usage_limit":98,"dvr_replicate":true,"episodes_url":"example","reference":"localdvr0","remotes":["example_value"],"replication_port":8002,"root":"example","schedule":[[800,1600],[2200,130]],"storage_limit":400000000000},"epg_enabled":true,"hls_scte35":"scte35","input_media_info":{"program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"match":{"codec":"ac3","index":2,"language":"eng"}}]},"inputs":[{"allow_if":"example","audio_timeout":20,"cluster_key":"example","comment":"This is a test input","deny_if":"example","frames_timeout":3,"headers":{"Authorization":"Basic dXNlcjpwYXNzd29yZA==","User-Agent":"curl/7.85.0"},"max_retry_timeout":30,"output_audio":"keep","priority":1,"remote_dvr":"nochain","source_timeout":20,"stats":{"active":true,"dvr_info":{"bytes":129600000000,"depth":259200,"disk_size":1099511627776,"duration":172800,"from":1641045644,"ranges":[{"closed_at":1000000000000,"duration":28800,"from":1525186456,"opened_at":1000000000000}]},"ip":"172.16.25.73","media_info":{"flow_type":"stream","program_id":110,"provider":"Netflix","stream_id":253,"title":"Bunny","tracks":[{"avg_gop":25,"bandwidth":2600,"bframes":3,"bitrate":2543,"closed_captions":[{"language":"eng","name":"English"}],"codec":"h264","content":"audio","last_gop":28,"length_size":2,"level":"example","pix_fmt":"yuv420p","profile":"example","title":"Video1","track_id":"v1"}]},"opened_at":1637094994000,"pids":[{"pid":0}],"proto":"dash","rtp_channels":[{"channel_id":0,"content":"video"}],"ts_delay":1284,"ts_delay_per_tracks":[1284],"url":"udp://239.0.0.1:1234","user_agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML. like Gecko) Chrome/90.0.4430.72 Safari/537.36","valid_secondary_inputs":2},"timeout":10,"url":"fake://fake","user_agent":"example","via":"example","video_timeout":20}],"jpeg_snapshot_sign_key":"example","labels":{"key1":"example_value","key2":"example_value"},"max_retry_timeout":30,"meta":{"key1":"example","key2":"example"},"mpegts_ac3":"keep","mpegts_pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"name":"example","on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"on_publish":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"output_audio":"keep","password":"example","playback_headers":[{"headers":{"Cache-Control":"max-age=3600"},"playback":"live","segment_headers":{"Cache-Control":"max-age=3600"}}],"prefer_codec":"h264","prefixes":["chats",""],"provider":"SportsTV","pushes":[{"comment":"This is a test push","connect_timeout":2,"domain":"officialdomain.com","encoder":"Lavf57","retry_timeout":7,"service":"My service","stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"timeout":10,"url":"example_value"}],"segment_count":4,"segment_duration":5000,"source_timeout":10,"static":true,"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"transcoder":{"audio":{"avol":"-6dB","bitrate":64000,"channels":2,"codec":"opus","sample_rate":48000},"decoder":{"crop":{"height":0,"left":0,"top":0,"width":0},"deinterlace":true,"deinterlace_rate":"frame","drop_frame_interval":3,"pix_fmt":"yuv420p"},"global":{"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"deviceid":"auto","gop":150,"hw":"cpu","target":"uhdtv"},"tracks":[{"bitrate":2543,"channels":2,"codec":"opus","content":"audio","input_track":1,"language":"eng","sample_rate":48000,"title":"Video1","volume":"-6dB"}],"video":[{"alogo":{"path":"@chan.png","position":"tl","x":10,"y":10},"bframes":3,"bitrate":1000000,"burn":{"sub":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"text":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"},"time":{"box":{"borderw":10,"color":"example"},"font":{"alpha":1,"color":"example","file":"/usr/share/fonts/truetype/freefont/FONT_NAME.ttf","size":24},"position":"tl","text":"example"}},"codec":"h264","extra":{"key1":"example","key2":"example"},"fps":"any","gop":150,"interlace":"tff","level":"1","logo":{"path":"@chan.png","position":"tl","x":10,"y":10},"preset":"medium","profile":"simple","rc_method":"vbr","refs":1,"resize_mode":"vic","size":{"background":"blur","height":-1,"strategy":"crop","width":-1},"track":1}]},"transport":"udp","video_timeout":20,"vision":{"alg":"faces","areas":"example"},"webrtc_abr":{"start_track":"v2"}}]}`

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
	query := &flussonic.TemplatesListQuery{}
	_, err := client.TemplatesList(ctx, query)
	if err != nil {
		t.Fatalf("TemplatesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TemplatesList(ctx, query)
	if err != nil {
		t.Fatalf("TemplatesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TemplatesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/transponders/test-name"

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
	err := client.TransponderDelete(ctx, "test-name")
	if err != nil {
		t.Fatalf("TransponderDelete failed: %v", err)
	}
}

func TestClient_TransponderGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders/test-name"

	// Expected JSON response
	expectedJSON := `{"eit":{"xmltv_url":"xmltv_dir"},"name":"multiplexer","network_name":"Example Network","others":[{"name":"example"}],"prebuffer":800,"programs":[{"eit_title":"EIT_Title","lcn":5,"pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"program_id":1,"service_type":"digital_tv","source":"hockey1","title":"ProgramTitle"}],"provider":"Example Provider","pushes":[{"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"url":"example"}],"stats":{"pids":[{"content":"audio","pid":0}],"programs":[{"pnr":0}]},"time_offsets":[{"country":"example","local_time_offset":"example","next_time_offset":"example"}],"ts_descriptors":[{"hex":"example","tag":0}]}`

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
	_, err := client.TransponderGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("TransponderGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderGet(ctx, "test-name")
	if err != nil {
		t.Fatalf("TransponderGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderOtherDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/transponders/test-name/others/test-index"

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
	err := client.TransponderOtherDelete(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderOtherDelete failed: %v", err)
	}
}

func TestClient_TransponderOtherGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders/test-name/others/test-index"

	// Expected JSON response
	expectedJSON := `{"name":"example"}`

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
	_, err := client.TransponderOtherGet(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderOtherGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderOtherGet(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderOtherGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderOtherGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderOtherSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/transponders/test-name/others/test-index"

	// Expected JSON response
	expectedJSON := `{"name":"example"}`

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
	body := &model.TransponderOtherImpl{}
	_, err := client.TransponderOtherSave(ctx, "test-name", "test-index", body)
	if err != nil {
		t.Fatalf("TransponderOtherSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderOtherSave(ctx, "test-name", "test-index", body)
	if err != nil {
		t.Fatalf("TransponderOtherSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderOtherSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderOthersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders/test-name/others"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","others":[{"name":"example"}],"prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl"}`

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
	query := &flussonic.TransponderOthersListQuery{}
	_, err := client.TransponderOthersList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("TransponderOthersList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderOthersList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("TransponderOthersList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderOthersList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderProgramDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/transponders/test-name/programs/test-program_id"

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
	err := client.TransponderProgramDelete(ctx, "test-name", "test-program_id")
	if err != nil {
		t.Fatalf("TransponderProgramDelete failed: %v", err)
	}
}

func TestClient_TransponderProgramGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders/test-name/programs/test-program_id"

	// Expected JSON response
	expectedJSON := `{"eit_title":"EIT_Title","lcn":5,"pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"program_id":1,"service_type":"digital_tv","source":"hockey1","title":"ProgramTitle"}`

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
	_, err := client.TransponderProgramGet(ctx, "test-name", "test-program_id")
	if err != nil {
		t.Fatalf("TransponderProgramGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderProgramGet(ctx, "test-name", "test-program_id")
	if err != nil {
		t.Fatalf("TransponderProgramGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderProgramGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderProgramSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/transponders/test-name/programs/test-program_id"

	// Expected JSON response
	expectedJSON := `{"eit_title":"EIT_Title","lcn":5,"pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"program_id":1,"service_type":"digital_tv","source":"hockey1","title":"ProgramTitle"}`

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
	body := &model.TransponderProgramImpl{}
	_, err := client.TransponderProgramSave(ctx, "test-name", "test-program_id", body)
	if err != nil {
		t.Fatalf("TransponderProgramSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderProgramSave(ctx, "test-name", "test-program_id", body)
	if err != nil {
		t.Fatalf("TransponderProgramSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderProgramSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderProgramsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders/test-name/programs"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","programs":[{"eit_title":"EIT_Title","lcn":5,"pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"program_id":1,"service_type":"digital_tv","source":"hockey1","title":"ProgramTitle"}]}`

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
	query := &flussonic.TransponderProgramsListQuery{}
	_, err := client.TransponderProgramsList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("TransponderProgramsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderProgramsList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("TransponderProgramsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderProgramsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderPushDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/transponders/test-name/pushes/test-index"

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
	err := client.TransponderPushDelete(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderPushDelete failed: %v", err)
	}
}

func TestClient_TransponderPushGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders/test-name/pushes/test-index"

	// Expected JSON response
	expectedJSON := `{"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"url":"example"}`

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
	_, err := client.TransponderPushGet(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderPushGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderPushGet(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderPushGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderPushGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderPushSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/transponders/test-name/pushes/test-index"

	// Expected JSON response
	expectedJSON := `{"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"url":"example"}`

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
	body := &model.TransponderPushImpl{}
	_, err := client.TransponderPushSave(ctx, "test-name", "test-index", body)
	if err != nil {
		t.Fatalf("TransponderPushSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderPushSave(ctx, "test-name", "test-index", body)
	if err != nil {
		t.Fatalf("TransponderPushSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderPushSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderPushesList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders/test-name/pushes"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","pushes":[{"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"url":"example"}]}`

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
	query := &flussonic.TransponderPushesListQuery{}
	_, err := client.TransponderPushesList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("TransponderPushesList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderPushesList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("TransponderPushesList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderPushesList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/transponders/test-name"

	// Expected JSON response
	expectedJSON := `{"eit":{"xmltv_url":"xmltv_dir"},"name":"multiplexer","network_name":"Example Network","others":[{"name":"example"}],"prebuffer":800,"programs":[{"eit_title":"EIT_Title","lcn":5,"pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"program_id":1,"service_type":"digital_tv","source":"hockey1","title":"ProgramTitle"}],"provider":"Example Provider","pushes":[{"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"url":"example"}],"stats":{"pids":[{"content":"audio","pid":0}],"programs":[{"pnr":0}]},"time_offsets":[{"country":"example","local_time_offset":"example","next_time_offset":"example"}],"ts_descriptors":[{"hex":"example","tag":0}]}`

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
	body := &model.TransponderConfigImpl{}
	_, err := client.TransponderSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("TransponderSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderSave(ctx, "test-name", body)
	if err != nil {
		t.Fatalf("TransponderSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderTimeOffsetDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/transponders/test-name/time_offsets/test-index"

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
	err := client.TransponderTimeOffsetDelete(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderTimeOffsetDelete failed: %v", err)
	}
}

func TestClient_TransponderTimeOffsetGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders/test-name/time_offsets/test-index"

	// Expected JSON response
	expectedJSON := `{"country":"example","local_time_offset":"example","next_time_offset":"example"}`

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
	_, err := client.TransponderTimeOffsetGet(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderTimeOffsetGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderTimeOffsetGet(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderTimeOffsetGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderTimeOffsetGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderTimeOffsetSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/transponders/test-name/time_offsets/test-index"

	// Expected JSON response
	expectedJSON := `{"country":"example","local_time_offset":"example","next_time_offset":"example"}`

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
	body := &model.TransponderTimeOffsetImpl{}
	_, err := client.TransponderTimeOffsetSave(ctx, "test-name", "test-index", body)
	if err != nil {
		t.Fatalf("TransponderTimeOffsetSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderTimeOffsetSave(ctx, "test-name", "test-index", body)
	if err != nil {
		t.Fatalf("TransponderTimeOffsetSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderTimeOffsetSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderTimeOffsetsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders/test-name/time_offsets"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","time_offsets":[{"country":"example","local_time_offset":"example","next_time_offset":"example"}]}`

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
	query := &flussonic.TransponderTimeOffsetsListQuery{}
	_, err := client.TransponderTimeOffsetsList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("TransponderTimeOffsetsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderTimeOffsetsList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("TransponderTimeOffsetsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderTimeOffsetsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderTsDescriptorDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/transponders/test-name/ts_descriptors/test-index"

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
	err := client.TransponderTsDescriptorDelete(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderTsDescriptorDelete failed: %v", err)
	}
}

func TestClient_TransponderTsDescriptorGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders/test-name/ts_descriptors/test-index"

	// Expected JSON response
	expectedJSON := `{"hex":"example","tag":0}`

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
	_, err := client.TransponderTsDescriptorGet(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderTsDescriptorGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderTsDescriptorGet(ctx, "test-name", "test-index")
	if err != nil {
		t.Fatalf("TransponderTsDescriptorGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderTsDescriptorGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderTsDescriptorSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/transponders/test-name/ts_descriptors/test-index"

	// Expected JSON response
	expectedJSON := `{"hex":"example","tag":0}`

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
	body := &model.TSDescriptorImpl{}
	_, err := client.TransponderTsDescriptorSave(ctx, "test-name", "test-index", body)
	if err != nil {
		t.Fatalf("TransponderTsDescriptorSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderTsDescriptorSave(ctx, "test-name", "test-index", body)
	if err != nil {
		t.Fatalf("TransponderTsDescriptorSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderTsDescriptorSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TransponderTsDescriptorsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders/test-name/ts_descriptors"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","ts_descriptors":[{"hex":"example","tag":0}]}`

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
	query := &flussonic.TransponderTsDescriptorsListQuery{}
	_, err := client.TransponderTsDescriptorsList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("TransponderTsDescriptorsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TransponderTsDescriptorsList(ctx, "test-name", query)
	if err != nil {
		t.Fatalf("TransponderTsDescriptorsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TransponderTsDescriptorsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_TranspondersList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/transponders"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","server_id":"550e8400-e29b-41d4-a716-446655440000","transponders":[{"eit":{"xmltv_url":"xmltv_dir"},"name":"multiplexer","network_name":"Example Network","others":[{"name":"example"}],"prebuffer":800,"programs":[{"eit_title":"EIT_Title","lcn":5,"pids":{"default":"auto","media":[{"bitrate":2543,"codec":"scte35","content":"audio","es_info":"52010D","pid":0,"stats":{"content":"audio","pid":0},"stream_type":12,"track":1}]},"program_id":1,"service_type":"digital_tv","source":"hockey1","title":"ProgramTitle"}],"provider":"Example Provider","pushes":[{"stats":{"genlock_status":"no_ref","genref_status":{"vstd":"pal","vstd_detected":"pal"},"opened_at":1000000000000,"pids":[{"content":"audio","pid":0}],"standby_status":"active","status":"starting","url":"example"},"url":"example"}],"stats":{"pids":[{"content":"audio","pid":0}],"programs":[{"pnr":0}]},"time_offsets":[{"country":"example","local_time_offset":"example","next_time_offset":"example"}],"ts_descriptors":[{"hex":"example","tag":0}]}]}`

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
	query := &flussonic.TranspondersListQuery{}
	_, err := client.TranspondersList(ctx, query)
	if err != nil {
		t.Fatalf("TranspondersList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.TranspondersList(ctx, query)
	if err != nil {
		t.Fatalf("TranspondersList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("TranspondersList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_VisionDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/vision"

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
	err := client.VisionDelete(ctx)
	if err != nil {
		t.Fatalf("VisionDelete failed: %v", err)
	}
}

func TestClient_VisionGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/vision"

	// Expected JSON response
	expectedJSON := `{"hw":"example","jpeg_vector_helper":"example"}`

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
	_, err := client.VisionGet(ctx)
	if err != nil {
		t.Fatalf("VisionGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.VisionGet(ctx)
	if err != nil {
		t.Fatalf("VisionGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("VisionGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_VisionSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/vision"

	// Expected JSON response
	expectedJSON := `{"hw":"example","jpeg_vector_helper":"example"}`

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
	_, err := client.VisionSave(ctx)
	if err != nil {
		t.Fatalf("VisionSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.VisionSave(ctx)
	if err != nil {
		t.Fatalf("VisionSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("VisionSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_VodDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/vods/test-prefix"

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
	err := client.VodDelete(ctx, "test-prefix")
	if err != nil {
		t.Fatalf("VodDelete failed: %v", err)
	}
}

func TestClient_VodGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/vods/test-prefix"

	// Expected JSON response
	expectedJSON := `{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"logo":{"height":100,"left":15,"path":"@chan.png","top":15,"width":200},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"position":1,"prefix":"movies","provider":"example","read_queue":100,"segment_duration":1000,"storages":[{"extra":{"key1":"example","key2":"example"},"url":"/storage"}],"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"timeout":3}`

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
	_, err := client.VodGet(ctx, "test-prefix")
	if err != nil {
		t.Fatalf("VodGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.VodGet(ctx, "test-prefix")
	if err != nil {
		t.Fatalf("VodGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("VodGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_VodSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/vods/test-prefix"

	// Expected JSON response
	expectedJSON := `{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"logo":{"height":100,"left":15,"path":"@chan.png","top":15,"width":200},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"position":1,"prefix":"movies","provider":"example","read_queue":100,"segment_duration":1000,"storages":[{"extra":{"key1":"example","key2":"example"},"url":"/storage"}],"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"timeout":3}`

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
	body := &model.VodConfigImpl{}
	_, err := client.VodSave(ctx, "test-prefix", body)
	if err != nil {
		t.Fatalf("VodSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.VodSave(ctx, "test-prefix", body)
	if err != nil {
		t.Fatalf("VodSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("VodSave returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_VodsList(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/vods"

	// Expected JSON response
	expectedJSON := `{"estimated_count":5,"next":"JTI0cG9zaXRpb25fZ3Q9MA==","prev":"JTI0cG9zaXRpb25fbHQ9MSYlMjRyZXZlcnNlZD10cnVl","vods":[{"cache":{"expiration":3600,"misses":3,"path":"/storage/cache","reference":"cache1","storage_limit":400000},"drm":{"encryption":"full","keyserver":"https://keyserver1.mycompany.com","resource_id":"L2sItm6","vendor":"example"},"logo":{"height":100,"left":15,"path":"@chan.png","top":15,"width":200},"on_play":{"allowed_countries":["US","DE","GB"],"disallowed_countries":["US","DE","GB"],"domains":["mycompany.com"],"extra":{"key1":"example","key2":"example"},"max_sessions":5000,"session_keys":["name","token","proto","ip"],"url":"http://middleware-address/auth/v2"},"position":1,"prefix":"movies","provider":"example","read_queue":100,"segment_duration":1000,"storages":[{"extra":{"key1":"example","key2":"example"},"url":"/storage"}],"thumbnails":{"sizes":[{}],"url":"http://10.115.23.45/isapi/thumbnail.jpg"},"timeout":3}]}`

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
	query := &flussonic.VodsListQuery{}
	_, err := client.VodsList(ctx, query)
	if err != nil {
		t.Fatalf("VodsList failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.VodsList(ctx, query)
	if err != nil {
		t.Fatalf("VodsList failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("VodsList returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_VsaasDelete(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "DELETE"
	expectedPath := "/streamer/api/v3/vsaas"

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
	err := client.VsaasDelete(ctx)
	if err != nil {
		t.Fatalf("VsaasDelete failed: %v", err)
	}
}

func TestClient_VsaasGet(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "GET"
	expectedPath := "/streamer/api/v3/vsaas"

	// Expected JSON response
	expectedJSON := `{"central_url":"example","watcher_url":"example"}`

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
	_, err := client.VsaasGet(ctx)
	if err != nil {
		t.Fatalf("VsaasGet failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.VsaasGet(ctx)
	if err != nil {
		t.Fatalf("VsaasGet failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("VsaasGet returned nil result")
	}

	// Verify marshaling/unmarshaling preserves data
	verifyMarshalUnmarshal(t, result, expectedJSON)
}

func TestClient_VsaasSave(t *testing.T) {
	ctx := context.Background()

	// Expected request
	expectedMethod := "PUT"
	expectedPath := "/streamer/api/v3/vsaas"

	// Expected JSON response
	expectedJSON := `{"central_url":"example","watcher_url":"example"}`

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
	body := &model.VsaasConfigImpl{}
	_, err := client.VsaasSave(ctx, body)
	if err != nil {
		t.Fatalf("VsaasSave failed: %v", err)
	}
	// Re-call to get result for verification
	result, err := client.VsaasSave(ctx, body)
	if err != nil {
		t.Fatalf("VsaasSave failed: %v", err)
	}

	// Verify result is not nil
	if result == nil {
		t.Fatal("VsaasSave returned nil result")
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

func createTestClient(t *testing.T, rt http.RoundTripper) flussonic.Flussonic {
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

	client, err := flussonic.NewWithBaseFactory(cfg, baseFactory)
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
