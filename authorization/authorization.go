package authorization

import (
	// nolint:gosec
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// AuthKey represents interface for various authorization types
type AuthKey interface {
	// ToHeader returns Authorization header value
	ToHeader() string
}

// ClusterKey implements authorization in Cluster timestamp:hash format
type clusterKey struct {
	key string
}

// ClusterKey creates a new AuthKey with ClusterKey authorization
func ClusterKey(key string) AuthKey {
	return &clusterKey{key: key}
}

// ToHeader generates authorization header value in Cluster format
// nolint:gosec
func (c *clusterKey) ToHeader() string {
	unixTime := time.Now().Unix()
	s := fmt.Appendf(nil, "%s:%s", strconv.FormatInt(unixTime, 10), c.key)

	hasher := sha1.New()
	hasher.Write(s)

	return fmt.Sprintf("Cluster %d:%s", unixTime, hex.EncodeToString(hasher.Sum(nil)))
}

// basicAuth implements Basic authorization
type basicAuth struct {
	username string
	password string
}

// BasicAuth creates a new AuthKey with Basic authorization
func BasicAuth(username, password string) AuthKey {
	return &basicAuth{
		username: username,
		password: password,
	}
}

// ToHeader generates Basic authorization header value
func (b *basicAuth) ToHeader() string {
	credentials := fmt.Sprintf("%s:%s", b.username, b.password)
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
	return fmt.Sprintf("Basic %s", encoded)
}

// bearerAuth implements Bearer authorization
type bearerAuth struct {
	token string
}

// BearerAuth creates a new AuthKey with Bearer authorization
func BearerAuth(token string) AuthKey {
	return &bearerAuth{token: token}
}

// ToHeader generates Bearer authorization header value
func (b *bearerAuth) ToHeader() string {
	return fmt.Sprintf("Bearer %s", b.token)
}

// ClusterKeyOld represents additional headers for old authorization format
type ClusterKeyOld struct {
	Hash string
	Time int64
}

// GenerateClusterKeyOld creates additional headers for old API
// Returns main Authorization header value and additional data
// nolint:gosec
func GenerateClusterKeyOld(clusterKey string) (authValue string, old ClusterKeyOld) {
	unixTime := time.Now().Unix()
	s := []byte(fmt.Sprintf("%s:%s", strconv.FormatInt(unixTime, 10), clusterKey))

	hasher := sha1.New()
	hasher.Write(s)

	hash := hex.EncodeToString(hasher.Sum(nil))
	authValue = fmt.Sprintf("Cluster %d:%s", unixTime, hash)

	return authValue, ClusterKeyOld{
		Time: unixTime,
		Hash: hash,
	}
}
