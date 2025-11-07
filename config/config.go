package config

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/flussonic/go-flussonic/authorization"
)

const (
	DefaultOriginator = "flussonic-sdk"
)

// ParseURL parses a Flussonic URL and returns a Config
// URL format: proto://<authorization>@host:port?param1=val&param2=val
// Examples:
//   - http://bearer-key@streamer.example.com:80
//   - https://basic_user:basic_password@streamer.example.com:443
//   - http://streamer.example.com?cluster_key=1234567890 (port will default to 80)
func ParseURL(rawURL string) (*Config, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	cfg := &Config{
		Protocol: parsedURL.Scheme,
		Hostname: parsedURL.Hostname(),
	}

	if query := parsedURL.Query(); query.Get("cluster_key") != "" {
		cfg.Auth = authorization.ClusterKey(query.Get("cluster_key"))
	}

	if parsedURL.User != nil {
		if pass, ok := parsedURL.User.Password(); ok {
			cfg.Auth = authorization.BasicAuth(parsedURL.User.Username(), pass)
		} else {
			cfg.Auth = authorization.BearerAuth(parsedURL.User.Username())
		}
	}
	// Extract port
	if parsedURL.Port() != "" {
		port, err := strconv.Atoi(parsedURL.Port())
		if err != nil {
			return nil, fmt.Errorf("invalid port: %w", err)
		}
		cfg.Port = port
	}

	if cfg.Port == 0 {
		if cfg.Protocol == "https" {
			cfg.Port = 443
		} else {
			cfg.Port = 80
		}
	}

	// Extract originator from query parameters
	if parsedURL.Query().Get("originator") != "" {
		cfg.Originator = parsedURL.Query().Get("originator")
	}

	// Parse query parameters if needed (for future extensions)
	// queryParams := parsedURL.Query()

	return cfg, nil
}

// Config defines configuration for the Flussonic client
type Config struct {
	Auth       authorization.AuthKey
	HTTPClient *http.Client
	Hostname   string
	ClusterKey string
	Protocol   string
	Originator string
	Port       int
	Retry      uint
}
