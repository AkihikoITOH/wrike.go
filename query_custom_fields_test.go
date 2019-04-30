package wrike_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
)

var allCustomFieldsData = []byte(
	`{
     "kind": "customfields",
     "data":
     [
         {
             "id": "IEAAAAAQJUAAAAAP",
             "accountId": "IEAAAAAQ",
             "title": "test 1",
             "type": "Text",
             "sharedIds":
             [
             ],
             "settings":
             {
                 "inheritanceType": "All"
             }
         },
         {
             "id": "IEAAAAAQJUAAAAAQ",
             "accountId": "IEAAAAAQ",
             "title": "test 2",
             "type": "Text",
             "sharedIds":
             [
             ],
             "settings":
             {
                 "inheritanceType": "All"
             }
         },
         {
             "id": "IEAAAAAQJUAAAAAT",
             "accountId": "IEAAAAAQ",
             "title": "test 1",
             "type": "Text",
             "sharedIds":
             [
             ],
             "settings":
             {
                 "inheritanceType": "All"
             }
         }
     ]
  }`)

var customFieldsByID = []byte(
	`{
     "kind": "customfields",
     "data":
     [
         {
             "id": "IEAAAAAQJUAAAAA2",
             "accountId": "IEAAAAAQ",
             "title": "Test custom field",
             "type": "Text",
             "sharedIds":
             [
             ],
             "settings":
             {
                 "inheritanceType": "All"
             }
         },
         {
             "id": "IEAAAAAQJUAAAAAZ",
             "accountId": "IEAAAAAQ",
             "title": "New test custom field",
             "type": "Numeric",
             "sharedIds":
             [
                 "KUAAAAAQ"
             ],
             "settings":
             {
                 "inheritanceType": "All",
                 "decimalPlaces": 2,
                 "useThousandsSeparator": false,
                 "aggregation": "Sum"
             }
         }
     ]
  }`)

func Example_queryCustomFields() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	api.QueryCustomFields()
}

func Example_queryCustomFieldsByIDs() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	ids := []string{"IEAAAAAQJUAAAAAZ", "IEAAAAAQJUAAAAA2"}
	api.QueryCustomFieldsByIDs(ids)
}

func TestQueryCustomFields(t *testing.T) {
	client := NewMockClient(allCustomFieldsData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	customFields, err := api.QueryCustomFields()
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	data, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, data, "https://app-eu.wrike.com/api/v4/customfields")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(customFields.Data), 3)
}

func TestQueryCustomFieldsByIDs(t *testing.T) {
	client := NewMockClient(customFieldsByID)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	ids := []string{"IEAAAAAQJUAAAAAZ", "IEAAAAAQJUAAAAA2"}
	customFields, err := api.QueryCustomFieldsByIDs(ids)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	data, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, data, "https://app-eu.wrike.com/api/v4/customfields/IEAAAAAQJUAAAAAZ,IEAAAAAQJUAAAAA2")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(customFields.Data), 2)
}
