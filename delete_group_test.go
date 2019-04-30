package wrike_test

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
)

var deletedGroupData = []byte(
	`{
    "kind": "groups",
    "data": []
  }`)

func Example_deleteGroup() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	api.DeleteGroup("KX77777H")
}

func TestDeleteGroup(t *testing.T) {
	client := NewMockClient(deletedGroupData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	groups, err := api.DeleteGroup("KX77777H")
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodDelete)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/groups/KX77777H")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(groups.Data), 0)
}
