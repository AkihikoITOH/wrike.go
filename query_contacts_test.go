package wrike_test

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
	"github.com/AkihikoITOH/wrike.go/types"
)

var allContactsData = []byte(
	`{
     "kind": "contacts",
     "data":
     [
         {
             "id": "KUAAAAAQ",
             "firstName": "Alexander",
             "lastName": "Spitsyn",
             "type": "Person",
             "profiles":
             [
                 {
                     "accountId": "IEAAAAAQ",
                     "email": "rxoxamhvky@jm5rpcdnwm.com",
                     "role": "User",
                     "external": false,
                     "admin": false,
                     "owner": true
                 }
             ],
             "avatarUrl": "https://dev.wrke.io/avatars//4E/7D/Box_ffffff00_65-83_v1.png",
             "timezone": "Europe/Moscow",
             "locale": "en",
             "deleted": false,
             "me": true,
             "metadata":
             [
                 {
                     "key": "testMetaKey",
                     "value": "testMetaValue"
                 }
             ],
             "phone": "+1 111111111"
         },
         {
             "id": "KX77777H",
             "firstName": "New test group",
             "lastName": "",
             "type": "Group",
             "profiles":
             [
                 {
                     "accountId": "IEAAAAAQ",
                     "role": "User",
                     "external": false,
                     "admin": false,
                     "owner": false
                 }
             ],
             "avatarUrl": "https://dev.wrke.io/avatars//D5/E6/Circle_ffc5cbd9_78-71_v1.png",
             "timezone": "Z",
             "locale": "en",
             "deleted": false,
             "memberIds":
             [
                 "KUAAAAAQ"
             ],
             "metadata":
             [
                 {
                     "key": "testMetaKey",
                     "value": "testMetaValue"
                 }
             ]
         }
     ]
  }`)

func NewParams() *wrike.QueryContactsParams {
	return &wrike.QueryContactsParams{
		Metadata: &types.Metadata{Key: "testMetaKey", Value: "testMetaValue"},
		Fields:   []string{"metadata"},
	}
}

func Example_queryContacts() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := &wrike.QueryContactsParams{
		Metadata: &types.Metadata{Key: "testMetaKey", Value: "testMetaValue"},
		Fields:   []string{"metadata"},
	}
	api.QueryContacts(params)
}

func TestQueryContacts(t *testing.T) {
	client := NewMockClient(allContactsData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewParams()
	contacts, err := api.QueryContacts(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/contacts?metadata={\"key\":\"testMetaKey\",\"value\":\"testMetaValue\"}&fields=[\"metadata\"]")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(contacts.Data), 2)
}
