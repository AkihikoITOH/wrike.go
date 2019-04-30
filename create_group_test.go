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

var createdGroupData = []byte(
	`{
     "kind": "groups",
     "data":
     [
         {
             "id": "KX77777H",
             "accountId": "IEAAAAAQ",
             "title": "Test group",
             "memberIds":
             [
                 "KUAAAAHK",
                 "KUAAAAAQ"
             ],
             "childIds":
             [
             ],
             "parentIds":
             [
                 "KX77777I"
             ],
             "avatarUrl": "/2E/52/Circle_ffe7fac6_84-71_v1.png",
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

func NewCreateGroupParams() parameters.CreateGroup {
	parent := types.ContactID("KX77777I")
	return parameters.CreateGroup{
		Title:    "Test",
		Members:  parameters.ContactIDSet{"KUAAAAHK", "KUAAAAAQ"},
		Parent:   &parent,
		Avatar:   &parameters.Avatar{Letters: "TG", Color: "#e7fac6"},
		Metadata: &parameters.MetadataSet{parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"}},
	}
}

func Example_createGroup() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	parent := types.ContactID("KX77777I")
	params := parameters.CreateGroup{
		Title:    "Test",
		Members:  parameters.ContactIDSet{"KUAAAAHK", "KUAAAAAQ"},
		Parent:   &parent,
		Avatar:   &parameters.Avatar{Letters: "TG", Color: "#e7fac6"},
		Metadata: &parameters.MetadataSet{parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"}},
	}
	api.CreateGroup(params)
}

func TestCreateGroup(t *testing.T) {
	client := NewMockClient(createdGroupData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewCreateGroupParams()
	groups, err := api.CreateGroup(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPost)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/groups")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "avatar={\"letters\":\"TG\",\"color\":\"#e7fac6\"}&members=[\"KUAAAAHK\",\"KUAAAAAQ\"]&metadata=[{\"key\":\"testMetaKey\",\"value\":\"testMetaValue\"}]&parent=KX77777I&title=Test")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(groups.Data[0].ID), "KX77777H")
}
