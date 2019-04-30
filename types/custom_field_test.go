package types_test

import (
	"fmt"
	"testing"

	"gotest.tools/assert"

	"github.com/AkihikoITOH/wrike.go/types"
)

var customFieldsData = []byte(
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

func TestNewCustomFieldsFromJSON(t *testing.T) {
	customFields, err := types.NewCustomFieldsFromJSON(customFieldsData)

	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Assert(t, err == nil)
	data := customFields.Data
	assert.Equal(t, string(data[0].ID), "IEAAAAAQJUAAAAAP")
}
