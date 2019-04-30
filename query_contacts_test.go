package wrike_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
	"github.com/AkihikoITOH/wrike.go/parameters"
)

var contactsByIDsData = []byte(
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
             "id": "KUAAAAHK",
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
             "metadata":
             [
             ]
         }
     ]
  }`)

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

func NewQueryContactsParams() *parameters.QueryContacts {
	return &parameters.QueryContacts{
		Metadata: &parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"},
		Fields:   &parameters.FieldSet{"metadata"},
	}
}

func Example_queryContacts() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := &parameters.QueryContacts{
		Metadata: &parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"},
		Fields:   &parameters.FieldSet{"metadata"},
	}
	api.QueryContacts(params)
}

func Example_queryContactsByIDs() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := &parameters.QueryContacts{
		Metadata: &parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"},
		Fields:   &parameters.FieldSet{"metadata"},
	}
	api.QueryContactsByIDs(parameters.ContactIDSet{"KUAAAAHK", "KUAAAAAQ"}, params)
}

func TestQueryContacts(t *testing.T) {
	client := NewMockClient(allContactsData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewQueryContactsParams()
	contacts, err := api.QueryContacts(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	urlStr, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, urlStr, "https://app-eu.wrike.com/api/v4/contacts?fields=[\"metadata\"]&metadata={\"key\":\"testMetaKey\",\"value\":\"testMetaValue\"}")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(contacts.Data), 2)
}

func TestQueryContactsByIDs(t *testing.T) {
	client := NewMockClient(contactsByIDsData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewQueryContactsParams()
	contacts, err := api.QueryContactsByIDs(parameters.ContactIDSet{"KUAAAAHK", "KUAAAAAQ"}, params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	urlStr, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, urlStr, "https://app-eu.wrike.com/api/v4/contacts/KUAAAAHK,KUAAAAAQ?fields=[\"metadata\"]&metadata={\"key\":\"testMetaKey\",\"value\":\"testMetaValue\"}")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(contacts.Data), 2)
}
