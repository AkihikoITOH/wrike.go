package types_test

import (
	"fmt"
	"testing"

	"gotest.tools/assert"

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

func TestNewAccountsFromJSON(t *testing.T) {
	accounts, err := types.NewAccountsFromJSON(allAccountsData)

	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Assert(t, err == nil)
	data := accounts.Data
	assert.Equal(t, len(data), 1)
	account := data[0]
	assert.Equal(t, string(account.ID), "IEAAAAAQ")
	assert.Equal(t, len(account.WorkDays), 3)
	assert.Equal(t, len(account.Metadata), 1)
	assert.Equal(t, len(account.CustomFields), 3)
}
