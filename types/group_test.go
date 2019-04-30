package types_test

import (
	"fmt"
	"testing"

	"gotest.tools/assert"

	"github.com/AkihikoITOH/wrike.go/types"
)

var groupData = []byte(
	`{
	   "kind": "groups",
	   "data":
	   [
       {
           "id": "KX77777H",
           "accountId": "IEAAAAAQ",
           "title": "New test group",
           "memberIds":
           [
               "KUAAAAAQ"
           ],
           "childIds":
           [
               "KX77777G"
           ],
           "parentIds":
           [
               "KX77777I"
           ],
           "avatarUrl": "/D5/E6/Circle_ffc5cbd9_78-71_v1.png"
       }
	   ]
	}`)

func TestNewGroupsFromJSON(t *testing.T) {
	groups, err := types.NewGroupsFromJSON(groupData)

	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Assert(t, err == nil)
	data := groups.Data
	assert.Equal(t, string(data[0].ID), "KX77777H")
}
