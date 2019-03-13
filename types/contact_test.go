package types_test

import (
	"fmt"
	"testing"

	"gotest.tools/assert"

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

func TestNewContactsFromJSON(t *testing.T) {
	contacts, err := types.NewContactsFromJSON(allContactsData)

	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Assert(t, err == nil)
	data := contacts.Data
	assert.Assert(t, data[0].IsPerson())
	assert.Assert(t, !data[0].IsGroup())
	assert.Assert(t, data[1].IsGroup())
	assert.Assert(t, !data[1].IsPerson())
}
