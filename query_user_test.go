package wrike_test

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
)

var userData = []byte(
	`{
     "kind": "users",
     "data":
     [
         {
             "id": "KUAAAAHP",
             "firstName": "Test",
             "lastName": "User",
             "type": "Person",
             "profiles":
             [
                 {
                     "accountId": "IEAAAAG3",
                     "email": "rvrphbdwox@airbpdhc5z.com",
                     "role": "User",
                     "external": false,
                     "admin": false,
                     "owner": true
                 }
             ],
             "avatarUrl": "https://dev.wrke.io/avatars//D7/A2/Box_ffffff00_84-84_v1.png",
             "timezone": "Europe/Moscow",
             "locale": "en",
             "deleted": false,
             "me": true
         }
     ]
  }`)

func Example_queryUser() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	api.QueryUser("KUAAAAHP")
}

func TestQueryUser(t *testing.T) {
	client := NewMockClient(userData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	user, err := api.QueryUser("KUAAAAHP")
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/user/KUAAAAHP")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(user.Data[0].ID), "KUAAAAHP")
}
