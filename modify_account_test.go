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

var modifiedAccountData = []byte(
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
             "joinedDate": "2019-02-11T09:45:40Z"
         }
     ]
  }`)

func NewModifyAccountParams() parameters.ModifyAccount {
	params := parameters.ModifyAccount{
		Metadata: &parameters.MetadataSet{
			parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"},
		},
	}
	return params
}

func Example_modifyAccount() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.ModifyAccount{
		Metadata: &parameters.MetadataSet{
			parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"},
		},
	}
	api.ModifyAccount(&params)
}

func TestModifyAccount(t *testing.T) {
	client := NewMockClient(modifiedAccountData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewModifyAccountParams()
	account, err := api.ModifyAccount(&params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPut)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/account")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "metadata=[{\"key\":\"testMetaKey\",\"value\":\"testMetaValue\"}]")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(account.Data[0].ID), "IEAAAAAQ")
}
