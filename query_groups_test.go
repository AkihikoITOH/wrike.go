package wrike_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
	"github.com/AkihikoITOH/wrike.go/parameters"
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

var allGroupsData = []byte(
	`{
	   "kind": "groups",
	   "data":
	   [
	       {
	           "id": "KX77777G",
	           "accountId": "IEAAAAAQ",
	           "title": "Test child group",
	           "memberIds":
	           [
	           ],
	           "childIds":
	           [
	           ],
	           "parentIds":
	           [
	               "KX77777H"
	           ],
	           "avatarUrl": "/D5/48/Circle_ff80cbc4_84-67_v1.png"
	       },
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
	       },
	       {
	           "id": "KX77777I",
	           "accountId": "IEAAAAAQ",
	           "title": "Test parent group",
	           "memberIds":
	           [
	           ],
	           "childIds":
	           [
	               "KX77777H"
	           ],
	           "parentIds":
	           [
	           ],
	           "avatarUrl": "/4C/D1/Circle_fff48fb1_84-80_v1.png"
	       }
	   ]
	}`)

func NewQueryGroupParams() parameters.QueryGroup {
	return parameters.QueryGroup{Fields: &parameters.FieldSet{"metadata"}}
}

func NewQueryGroupsParams() parameters.QueryGroups {
	return parameters.QueryGroups{
		Fields:   &parameters.FieldSet{"metadata"},
		Metadata: &parameters.Metadata{Key: "Test Key", Value: "Test Value"},
	}
}

func Example_queryGroup() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.QueryGroup{Fields: &parameters.FieldSet{"metadata"}}
	api.QueryGroup("KX77777H", &params)
}

func Example_queryGroups() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.QueryGroups{
		Fields:   &parameters.FieldSet{"metadata"},
		Metadata: &parameters.Metadata{Key: "Test Key", Value: "Test Value"},
	}
	api.QueryGroups(&params)
}

func TestQueryGroup(t *testing.T) {
	client := NewMockClient(groupData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewQueryGroupParams()
	groups, err := api.QueryGroup("KX77777H", &params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	data, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, data, "https://app-eu.wrike.com/api/v4/groups/KX77777H?fields=[\"metadata\"]")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(groups.Data[0].ID), "KX77777H")
}

func TestQueryGroups(t *testing.T) {
	client := NewMockClient(allGroupsData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewQueryGroupsParams()
	groups, err := api.QueryGroups(&params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	data, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, data, "https://app-eu.wrike.com/api/v4/groups?fields=[\"metadata\"]&metadata={\"key\":\"Test Key\",\"value\":\"Test Value\"}")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(groups.Data), 3)
}
