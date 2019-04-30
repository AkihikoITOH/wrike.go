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

var modifiedContactData = []byte(
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
       }
	   ]
	}`)

func NewModifyContactParams() parameters.ModifyContact {
	params := parameters.ModifyContact{
		Metadata: &parameters.MetadataSet{
			parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"},
		},
	}
	return params
}

func Example_modifyContact() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.ModifyContact{
		Metadata: &parameters.MetadataSet{
			parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"},
		},
	}
	api.ModifyContact("KUAAAAAQ", &params)
}

func TestModifyContact(t *testing.T) {
	id := "KUAAAAAQ"
	client := NewMockClient(modifiedContactData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewModifyContactParams()
	contact, err := api.ModifyContact(id, &params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPut)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/contacts/KUAAAAAQ")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "metadata=[{\"key\":\"testMetaKey\",\"value\":\"testMetaValue\"}]")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(contact.Data[0].ID), "KUAAAAAQ")
}
