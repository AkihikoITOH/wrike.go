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
	"github.com/AkihikoITOH/wrike.go/types"
)

var createdCustomFieldData = []byte(
	`{
   "kind": "customfields",
   "data":
   [
       {
           "id": "IEAAAAAQJUAAAAAZ",
           "accountId": "IEAAAAAQ",
           "title": "Test custom field",
           "type": "Text",
           "sharedIds":
           [
               "KUAAAAAQ"
           ],
           "settings":
           {
               "inheritanceType": "All"
           }
       }
   ]
}`)

func NewCreateCustomFieldParams() parameters.CreateCustomField {
	return parameters.CreateCustomField{
		Title:   "Test custom field",
		Type:    types.TextType,
		Shareds: parameters.ContactIDSet{"KUAAAAAQ"},
	}
}

func Example_createCustomField() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.CreateCustomField{
		Title:   "Test custom field",
		Type:    types.TextType,
		Shareds: parameters.ContactIDSet{"KUAAAAAQ"},
	}
	api.CreateCustomField(params)
}

func TestCreateCustomField(t *testing.T) {
	client := NewMockClient(createdCustomFieldData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewCreateCustomFieldParams()
	customFields, err := api.CreateCustomField(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPost)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/customFields")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "shareds=[\"KUAAAAAQ\"]&title=Test custom field&type=Text")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(customFields.Data[0].ID), "IEAAAAAQJUAAAAAZ")
}
