package wrike_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
	"github.com/AkihikoITOH/wrike.go/parameters"
)

var modifiedUserData = []byte(
	`{
     "kind": "users",
     "data":
     [
       {
           "id": "KUAAAAHP",
           "firstName": "full",
           "lastName": "name",
           "type": "Person",
           "profiles":
           [
               {
                   "accountId": "IEAAAAAQ",
                   "email": "rxoxamhvky@jm5rpcdnwm.com",
                   "role": "Collaborator",
                   "external": true,
                   "admin": false,
                   "owner": false
               }
           ],
           "avatarUrl": "https://dev.wrke.io/avatars//33/42/Box_fff9a825_70-78_v1.png",
           "timezone": "US/Pacific",
           "locale": "en",
           "deleted": false,
           "metadata": []
       }
     ]
  }`)

func NewModifyUserParams() *parameters.ModifyUser {
	accountID := "IEAAAAAQ"
	role := "Collaborator"
	external := true
	params := parameters.ModifyUser{
		Profile: &parameters.Profile{
			AccountID: &accountID,
			Role:      &role,
			External:  &external,
		},
	}
	return &params
}

func Example_modifyUser() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	accountID := "IEAAAAAQ"
	role := "Collaborator"
	external := true
	params := parameters.ModifyUser{
		Profile: &parameters.Profile{
			AccountID: &accountID,
			Role:      &role,
			External:  &external,
		},
	}
	api.ModifyUser("KUAAAAHP", &params)
}

func TestModifyUser(t *testing.T) {
	id := "KUAAAAHP"
	client := NewMockClient(modifiedUserData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewModifyUserParams()
	user, err := api.ModifyUser(id, params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPut)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/user/KUAAAAHP")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "profile={\"accountId\":\"IEAAAAAQ\",\"role\":\"Collaborator\",\"external\":true}")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(user.Data[0].ID), "KUAAAAHP")
}
