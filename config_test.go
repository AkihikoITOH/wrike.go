package wrike_test

import (
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
)

func TestNewConfig(t *testing.T) {
	// With default host
	accessToken := "wrike-access-token"

	config := wrike.NewConfig(accessToken, "")

	assert.Equal(t, config.APIAccessToken(), accessToken)
	assert.Equal(t, config.APIHost(), "app-eu.wrike.com")

	// With custom host
	accessToken = "wrike-access-token"
	apiHost := "custom.api-host.example.com"

	config = wrike.NewConfig(accessToken, apiHost)

	assert.Equal(t, config.APIAccessToken(), accessToken)
	assert.Equal(t, config.APIHost(), apiHost)
}
