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

var modifiedGroupData = []byte(
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
             "avatarUrl": "/D5/E6/Circle_ffc5cbd9_78-71_v1.png",
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

func NewModifyGroupParams() parameters.ModifyGroup {
	title := "New test group"
	parent := "KX77777I"
	params := parameters.ModifyGroup{
		Title:         &title,
		AddMembers:    parameters.ContactIDSet{"KUAAAAAQ"},
		RemoveMembers: parameters.ContactIDSet{"KUAAAAHK"},
		Parent:        &parent,
		Avatar:        &parameters.Avatar{Letters: "NG", Color: "#c5cbd9"},
		Metadata:      &parameters.MetadataSet{parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"}},
	}
	return params
}

func Example_modifyGroup() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	title := "New test group"
	parent := "KX77777I"
	params := parameters.ModifyGroup{
		Title:         &title,
		AddMembers:    parameters.ContactIDSet{"KUAAAAAQ"},
		RemoveMembers: parameters.ContactIDSet{"KUAAAAHK"},
		Parent:        &parent,
		Avatar:        &parameters.Avatar{Letters: "NG", Color: "#c5cbd9"},
		Metadata:      &parameters.MetadataSet{parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"}},
	}
	api.ModifyGroup("KX77777H", &params)
}

func TestModifyGroup(t *testing.T) {
	id := "KX77777H"
	client := NewMockClient(modifiedGroupData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewModifyGroupParams()
	group, err := api.ModifyGroup(id, &params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPut)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/groups/KX77777H")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "addMembers=[\"KUAAAAAQ\"]&avatar={\"letters\":\"NG\",\"color\":\"#c5cbd9\"}&metadata=[{\"key\":\"testMetaKey\",\"value\":\"testMetaValue\"}]&parent=KX77777I&removeMembers=[\"KUAAAAHK\"]&title=New test group")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(group.Data[0].ID), "KX77777H")
}
