package wrike_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
	"github.com/AkihikoITOH/wrike.go/parameters"
	"github.com/AkihikoITOH/wrike.go/types"
)

var allAccountsData = []byte(
	`{
     "kind": "accounts",
     "data":
     [
         {
             "id": "IEAAAAAQ",
             "name": "Account",
             "dateFormat": "MM/dd/yyyy",
             "firstDayOfWeek": "Mon",
             "workDays":
             [
                 "Mon",
                 "Tue",
                 "Wed"
             ],
             "rootFolderId": "IEAAAAAQI7777777",
             "recycleBinId": "IEAAAAAQI7777776",
             "createdDate": "2019-02-11T09:45:40Z",
             "subscription":
             {
                 "type": "Enterprise",
                 "paid": false,
                 "userLimit": 50
             },
             "metadata":
             [
                 {
                     "key": "testMetaKey",
                     "value": "testMetaValue"
                 }
             ],
             "customFields":
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
             ],
             "joinedDate": "2019-02-11T09:45:40Z"
         }
     ]
  }`)

func NewQueryAccountsParams() parameters.QueryAccounts {
	return parameters.QueryAccounts{
		Fields: &parameters.FieldSet{types.CustomFieldsField, types.SubscriptionField, types.MetadataField},
	}
}

func Example_queryAccounts() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.QueryAccounts{
		Fields: &parameters.FieldSet{types.CustomFieldsField, types.SubscriptionField, types.MetadataField},
	}
	api.QueryAccounts(&params)
}

func TestQueryAccounts(t *testing.T) {
	client := NewMockClient(allAccountsData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewQueryAccountsParams()
	accounts, err := api.QueryAccounts(&params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	urlStr, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, urlStr, "https://app-eu.wrike.com/api/v4/accounts?fields=[\"customFields\",\"subscription\",\"metadata\"]")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(accounts.Data), 1)
	assert.Equal(t, string(accounts.Data[0].ID), "IEAAAAAQ")
}
