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

var modifiedCustomFieldData = []byte(
	`{
	   "kind": "customfields",
	   "data":
	   [
	       {
	           "id": "IEAAAAX3JUAAAAHL",
	           "accountId": "IEAAAAX3",
	           "title": "New test custom field",
	           "type": "Numeric",
	           "sharedIds":
	           [
	               "KUAAAA3E"
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

func NewModifyCustomFieldParams() parameters.ModifyCustomField {
	title := "New test custom field"
	typ := types.NumericType
	params := parameters.ModifyCustomField{
		AddShareds:    parameters.ContactIDSet{"KUAAAA3E"},
		RemoveShareds: parameters.ContactIDSet{"KUAAAA3F"},
		Title:         &title,
		Type:          &typ,
	}
	return params
}

func Example_modifyCustomField() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	title := "New test custom field"
	typ := types.NumericType
	params := parameters.ModifyCustomField{
		AddShareds:    parameters.ContactIDSet{"KUAAAA3E"},
		RemoveShareds: parameters.ContactIDSet{"KUAAAA3F"},
		Title:         &title,
		Type:          &typ,
	}
	api.ModifyCustomField("IEAAAAX3JUAAAAHL", &params)
}

func TestModifyCustomField(t *testing.T) {
	id := "IEAAAAX3JUAAAAHL"
	client := NewMockClient(modifiedCustomFieldData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewModifyCustomFieldParams()
	custom_field, err := api.ModifyCustomField(id, &params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPut)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/customfields/IEAAAAX3JUAAAAHL")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "addShareds=[\"KUAAAA3E\"]&removeShareds=[\"KUAAAA3F\"]&title=New test custom field&type=Numeric")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(custom_field.Data[0].ID), "IEAAAAX3JUAAAAHL")
}
