package wrike

// DefaultAPIHost is the default API host name.
const DefaultAPIHost = "app-eu.wrike.com"

// Config contains configuration parameters of the BambooHR API.
type Config struct {
	apiAccessToken string
	apiHost        string
}

// NewConfig creates and returns a new Config object based on the given access token and api host.
func NewConfig(apiAccessToken, apiHost string) *Config {
	host := apiHost
	if len(host) == 0 {
		host = DefaultAPIHost
	}
	return &Config{apiAccessToken, host}
}

// APIHost returns the api host.
func (c *Config) APIHost() string {
	return c.apiHost
}

// APIAccessToken returns the access token.
func (c *Config) APIAccessToken() string {
	return c.apiAccessToken
}
