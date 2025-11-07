package authorization_test

import (
	// nolint:gosec
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flussonic/go-flussonic/authorization"
)

func TestClusterKey_ToHeader(t *testing.T) {
	t.Parallel()

	t.Run("generates correct format", func(t *testing.T) {
		t.Parallel()
		key := "test-cluster-key"
		auth := authorization.ClusterKey(key)
		header := auth.ToHeader()

		// Format should be "Cluster <timestamp>:<hash>"
		pattern := regexp.MustCompile(`^Cluster \d+:[a-f0-9]{40}$`)
		assert.Regexp(t, pattern, header, "header should match Cluster format")
	})

	t.Run("includes current timestamp", func(t *testing.T) {
		t.Parallel()
		key := "test-key"
		auth := authorization.ClusterKey(key)
		before := time.Now().Unix()
		header := auth.ToHeader()
		after := time.Now().Unix()

		// Extract timestamp from header
		parts := strings.Split(header, ":")
		require.Len(t, parts, 2, "header should have two parts separated by colon")
		require.True(t, strings.HasPrefix(parts[0], "Cluster "), "first part should start with 'Cluster '")

		timestampStr := strings.TrimPrefix(parts[0], "Cluster ")
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		require.NoError(t, err, "timestamp should be valid integer")

		assert.GreaterOrEqual(t, timestamp, before, "timestamp should be >= before")
		assert.LessOrEqual(t, timestamp, after, "timestamp should be <= after")
	})

	t.Run("generates correct hash", func(t *testing.T) {
		t.Parallel()
		key := "test-cluster-key-123"
		auth := authorization.ClusterKey(key)

		// Get header to extract timestamp
		header := auth.ToHeader()
		parts := strings.Split(header, ":")
		require.Len(t, parts, 2, "header should have two parts")

		timestampStr := strings.TrimPrefix(parts[0], "Cluster ")
		hashStr := parts[1]

		// Manually compute expected hash
		s := fmt.Sprintf("%s:%s", timestampStr, key)
		// nolint:gosec
		hasher := sha1.New()
		hasher.Write([]byte(s))
		expectedHash := hex.EncodeToString(hasher.Sum(nil))

		assert.Equal(t, expectedHash, hashStr, "hash should match computed value")
	})

	t.Run("generates different hashes for different keys", func(t *testing.T) {
		t.Parallel()
		key1 := "key1"
		key2 := "key2"
		auth1 := authorization.ClusterKey(key1)
		auth2 := authorization.ClusterKey(key2)

		header1 := auth1.ToHeader()
		header2 := auth2.ToHeader()

		// Extract hashes
		hash1 := strings.Split(header1, ":")[1]
		hash2 := strings.Split(header2, ":")[1]

		assert.NotEqual(t, hash1, hash2, "different keys should produce different hashes")
	})

	t.Run("handles empty key", func(t *testing.T) {
		t.Parallel()
		auth := authorization.ClusterKey("")
		header := auth.ToHeader()

		pattern := regexp.MustCompile(`^Cluster \d+:[a-f0-9]{40}$`)
		assert.Regexp(t, pattern, header, "should handle empty key")
	})

	t.Run("generates different timestamps on multiple calls", func(t *testing.T) {
		t.Parallel()
		key := "test-key"
		auth := authorization.ClusterKey(key)

		header1 := auth.ToHeader()
		time.Sleep(1 * time.Second) // Ensure different timestamp
		header2 := auth.ToHeader()

		// Extract timestamps
		parts1 := strings.Split(header1, ":")
		parts2 := strings.Split(header2, ":")
		timestamp1, _ := strconv.ParseInt(strings.TrimPrefix(parts1[0], "Cluster "), 10, 64)
		timestamp2, _ := strconv.ParseInt(strings.TrimPrefix(parts2[0], "Cluster "), 10, 64)

		assert.Greater(t, timestamp2, timestamp1, "second call should have later timestamp")
	})
}

func TestBasicAuth_ToHeader(t *testing.T) {
	t.Parallel()

	t.Run("generates correct format", func(t *testing.T) {
		t.Parallel()
		auth := authorization.BasicAuth("user", "pass")
		header := auth.ToHeader()

		assert.True(t, strings.HasPrefix(header, "Basic "), "header should start with 'Basic '")
	})

	t.Run("encodes credentials correctly", func(t *testing.T) {
		t.Parallel()
		username := "testuser"
		password := "testpass"
		auth := authorization.BasicAuth(username, password)
		header := auth.ToHeader()

		// Extract base64 part
		encoded := strings.TrimPrefix(header, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		require.NoError(t, err, "encoded part should be valid base64")

		credentials := string(decoded)
		expected := fmt.Sprintf("%s:%s", username, password)
		assert.Equal(t, expected, credentials, "decoded credentials should match input")
	})

	t.Run("handles special characters", func(t *testing.T) {
		t.Parallel()
		username := "user@example.com"
		password := "p@ss:w0rd!"
		auth := authorization.BasicAuth(username, password)
		header := auth.ToHeader()

		encoded := strings.TrimPrefix(header, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		require.NoError(t, err)

		credentials := string(decoded)
		expected := fmt.Sprintf("%s:%s", username, password)
		assert.Equal(t, expected, credentials, "should handle special characters")
	})

	t.Run("handles empty username", func(t *testing.T) {
		t.Parallel()
		auth := authorization.BasicAuth("", "password")
		header := auth.ToHeader()

		encoded := strings.TrimPrefix(header, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		require.NoError(t, err)

		credentials := string(decoded)
		assert.Equal(t, ":password", credentials, "should handle empty username")
	})

	t.Run("handles empty password", func(t *testing.T) {
		t.Parallel()
		auth := authorization.BasicAuth("username", "")
		header := auth.ToHeader()

		encoded := strings.TrimPrefix(header, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		require.NoError(t, err)

		credentials := string(decoded)
		assert.Equal(t, "username:", credentials, "should handle empty password")
	})

	t.Run("handles empty both", func(t *testing.T) {
		t.Parallel()
		auth := authorization.BasicAuth("", "")
		header := auth.ToHeader()

		encoded := strings.TrimPrefix(header, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		require.NoError(t, err)

		credentials := string(decoded)
		assert.Equal(t, ":", credentials, "should handle empty both")
	})

	t.Run("produces consistent output for same inputs", func(t *testing.T) {
		t.Parallel()
		auth := authorization.BasicAuth("user", "pass")
		header1 := auth.ToHeader()
		header2 := auth.ToHeader()

		assert.Equal(t, header1, header2, "same inputs should produce same output")
	})
}

func TestBearerAuth_ToHeader(t *testing.T) {
	t.Parallel()

	t.Run("generates correct format", func(t *testing.T) {
		t.Parallel()
		token := "test-token-123"
		auth := authorization.BearerAuth(token)
		header := auth.ToHeader()

		expected := fmt.Sprintf("Bearer %s", token)
		assert.Equal(t, expected, header, "header should match Bearer format")
	})

	t.Run("handles empty token", func(t *testing.T) {
		t.Parallel()
		auth := authorization.BearerAuth("")
		header := auth.ToHeader()

		assert.Equal(t, "Bearer ", header, "should handle empty token")
	})

	t.Run("handles special characters in token", func(t *testing.T) {
		t.Parallel()
		token := "token-with-special-chars-!@#$%^&*()"
		auth := authorization.BearerAuth(token)
		header := auth.ToHeader()

		expected := fmt.Sprintf("Bearer %s", token)
		assert.Equal(t, expected, header, "should handle special characters")
	})

	t.Run("handles long token", func(t *testing.T) {
		t.Parallel()
		token := strings.Repeat("a", 1000)
		auth := authorization.BearerAuth(token)
		header := auth.ToHeader()

		expected := fmt.Sprintf("Bearer %s", token)
		assert.Equal(t, expected, header, "should handle long token")
	})

	t.Run("produces consistent output for same token", func(t *testing.T) {
		t.Parallel()
		token := "test-token"
		auth := authorization.BearerAuth(token)
		header1 := auth.ToHeader()
		header2 := auth.ToHeader()

		assert.Equal(t, header1, header2, "same token should produce same output")
	})
}

func TestGenerateClusterKeyOld(t *testing.T) {
	t.Parallel()

	t.Run("generates correct auth value format", func(t *testing.T) {
		t.Parallel()
		key := "test-cluster-key"
		authValue, old := authorization.GenerateClusterKeyOld(key)

		// Format should be "Cluster <timestamp>:<hash>"
		pattern := regexp.MustCompile(`^Cluster \d+:[a-f0-9]{40}$`)
		assert.Regexp(t, pattern, authValue, "authValue should match Cluster format")
		assert.NotEmpty(t, old.Hash, "Hash should not be empty")
		assert.Greater(t, old.Time, int64(0), "Time should be positive")
	})

	t.Run("includes current timestamp", func(t *testing.T) {
		t.Parallel()
		key := "test-key"
		before := time.Now().Unix()
		authValue, old := authorization.GenerateClusterKeyOld(key)
		after := time.Now().Unix()

		assert.GreaterOrEqual(t, old.Time, before, "Time should be >= before")
		assert.LessOrEqual(t, old.Time, after, "Time should be <= after")

		// Verify timestamp in authValue matches old.Time
		parts := strings.Split(authValue, ":")
		require.Len(t, parts, 2, "authValue should have two parts")
		timestampStr := strings.TrimPrefix(parts[0], "Cluster ")
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		require.NoError(t, err)

		assert.Equal(t, old.Time, timestamp, "timestamp in authValue should match old.Time")
	})

	t.Run("generates correct hash", func(t *testing.T) {
		t.Parallel()
		key := "test-cluster-key-456"
		authValue, old := authorization.GenerateClusterKeyOld(key)

		// Manually compute expected hash
		s := fmt.Sprintf("%d:%s", old.Time, key)
		// nolint:gosec
		hasher := sha1.New()
		hasher.Write([]byte(s))
		expectedHash := hex.EncodeToString(hasher.Sum(nil))

		assert.Equal(t, expectedHash, old.Hash, "Hash should match computed value")

		// Verify hash in authValue matches
		parts := strings.Split(authValue, ":")
		require.Len(t, parts, 2, "authValue should have two parts")
		hashInAuthValue := parts[1]
		assert.Equal(t, expectedHash, hashInAuthValue, "hash in authValue should match computed value")
	})

	t.Run("authValue matches old struct values", func(t *testing.T) {
		t.Parallel()
		key := "test-key"
		authValue, old := authorization.GenerateClusterKeyOld(key)

		// Reconstruct expected authValue from old struct
		expectedAuthValue := fmt.Sprintf("Cluster %d:%s", old.Time, old.Hash)
		assert.Equal(t, expectedAuthValue, authValue, "authValue should match reconstructed value from old struct")
	})

	t.Run("handles empty key", func(t *testing.T) {
		t.Parallel()
		authValue, old := authorization.GenerateClusterKeyOld("")

		pattern := regexp.MustCompile(`^Cluster \d+:[a-f0-9]{40}$`)
		assert.Regexp(t, pattern, authValue, "should handle empty key")
		assert.NotEmpty(t, old.Hash, "Hash should not be empty even with empty key")
		assert.Greater(t, old.Time, int64(0), "Time should be positive")
	})

	t.Run("generates different values on multiple calls", func(t *testing.T) {
		t.Parallel()
		key := "test-key"
		authValue1, old1 := authorization.GenerateClusterKeyOld(key)
		time.Sleep(1 * time.Second) // Ensure different timestamp
		authValue2, old2 := authorization.GenerateClusterKeyOld(key)

		assert.NotEqual(t, authValue1, authValue2, "different calls should produce different authValue")
		assert.Greater(t, old2.Time, old1.Time, "second call should have later timestamp")
		// Hashes will be different because timestamps are different
		assert.NotEqual(t, old1.Hash, old2.Hash, "different timestamps should produce different hashes")
	})

	t.Run("generates different hashes for different keys", func(t *testing.T) {
		t.Parallel()
		key1 := "key1"
		key2 := "key2"
		_, old1 := authorization.GenerateClusterKeyOld(key1)
		_, old2 := authorization.GenerateClusterKeyOld(key2)

		// Note: Hashes might be the same if generated at the exact same timestamp
		// but in practice they should be different
		// We can verify the hash computation is key-dependent
		s1 := fmt.Sprintf("%d:%s", old1.Time, key1)
		s2 := fmt.Sprintf("%d:%s", old2.Time, key2)
		if old1.Time == old2.Time {
			assert.NotEqual(t, old1.Hash, old2.Hash, "same timestamp but different keys should produce different hashes")
		}
		assert.NotEqual(t, s1, s2, "input strings should be different")
	})
}

func TestAuthKey_Interface(t *testing.T) {
	t.Parallel()

	// Helper function to verify interface compliance
	checkAuthKey := func(t *testing.T, auth authorization.AuthKey) {
		t.Helper()
		header := auth.ToHeader()
		assert.NotEmpty(t, header, "ToHeader should return non-empty string")
	}

	t.Run("ClusterKey implements AuthKey", func(t *testing.T) {
		t.Parallel()
		checkAuthKey(t, authorization.ClusterKey("test"))
	})

	t.Run("BasicAuth implements AuthKey", func(t *testing.T) {
		t.Parallel()
		checkAuthKey(t, authorization.BasicAuth("user", "pass"))
	})

	t.Run("BearerAuth implements AuthKey", func(t *testing.T) {
		t.Parallel()
		checkAuthKey(t, authorization.BearerAuth("token"))
	})
}
