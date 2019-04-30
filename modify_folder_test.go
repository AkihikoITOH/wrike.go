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

var modifiedFolderData = []byte(
	`{
     "kind": "folders",
     "data":
     [
         {
             "id": "IEAAAAAQI4AB5BGU",
             "accountId": "IEAAAAAQ",
             "title": "New title",
             "createdDate": "2019-02-18T11:17:42Z",
             "updatedDate": "2019-02-18T11:17:44Z",
             "description": "New description",
             "sharedIds":
             [
                 "KUAAAAAQ"
             ],
             "parentIds":
             [
                 "IEAAAAAQI7777777"
             ],
             "childIds":
             [
                 "IEAAAAAQI4AB5BGV",
                 "IEAAAAAQI4AB5BGW"
             ],
             "superParentIds":
             [
             ],
             "scope": "WsFolder",
             "hasAttachments": true,
             "permalink": "https://dev.wrke.io/open.htm?id=2000084",
             "workflowId": "IEAAAAAQK777777Q",
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
                     "id": "IEAAAAAQJUAAAAAX",
                     "value": "testValue"
                 }
             ],
             "project":
             {
                 "authorId": "KUAAAAAQ",
                 "ownerIds":
                 [
                     "KUAAAAAQ"
                 ],
                 "status": "Green",
                 "startDate": "2019-02-18",
                 "endDate": "2019-02-25",
                 "createdDate": "2019-02-18T11:17:44Z"
             }
         }
     ]
  }`)

var modifiedFoldersData = []byte(
	`{
     "kind": "folders",
     "data":
     [
         {
             "id": "IEAAAAAQI4AB5BGU",
             "accountId": "IEAAAAAQ",
             "title": "New title",
             "createdDate": "2019-02-18T11:17:42Z",
             "updatedDate": "2019-02-18T11:17:44Z",
             "description": "New description",
             "sharedIds":
             [
                 "KUAAAAAQ"
             ],
             "parentIds":
             [
                 "IEAAAAAQI7777777"
             ],
             "childIds":
             [
                 "IEAAAAAQI4AB5BGV",
                 "IEAAAAAQI4AB5BGW"
             ],
             "superParentIds":
             [
             ],
             "scope": "WsFolder",
             "hasAttachments": true,
             "permalink": "https://dev.wrke.io/open.htm?id=2000084",
             "workflowId": "IEAAAAAQK777777Q",
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
                     "id": "IEAAAAAQJUAAAAAX",
                     "value": "testValue"
                 },
                 {
                     "id": "IEAAAAAQJUAAAAAY",
                     "value": "testValue"
                 }
             ],
             "project":
             {
                 "authorId": "KUAAAAAQ",
                 "ownerIds":
                 [
                     "KUAAAAAQ"
                 ],
                 "status": "Green",
                 "startDate": "2019-02-18",
                 "endDate": "2019-02-25",
                 "createdDate": "2019-02-18T11:17:44Z"
             }
         },
         {
             "id": "IEAAAAAQI4AB5BGV",
             "accountId": "IEAAAAAQ",
             "title": "Test folder",
             "createdDate": "2019-02-18T11:17:42Z",
             "updatedDate": "2019-02-18T11:17:44Z",
             "description": "Test description",
             "sharedIds":
             [
                 "KUAAAAAQ"
             ],
             "parentIds":
             [
                 "IEAAAAAQI4AB5BGU"
             ],
             "childIds":
             [
             ],
             "superParentIds":
             [
             ],
             "scope": "WsFolder",
             "hasAttachments": false,
             "permalink": "https://dev.wrke.io/open.htm?id=2000085",
             "workflowId": "IEAAAAAQK777777Q",
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
                     "id": "IEAAAAAQJUAAAAAX",
                     "value": "testValue"
                 },
                 {
                     "id": "IEAAAAAQJUAAAAAY",
                     "value": "testValue"
                 }
             ],
             "project":
             {
                 "authorId": "KUAAAAAQ",
                 "ownerIds":
                 [
                     "KUAAAAAQ"
                 ],
                 "status": "Green",
                 "startDate": "2019-02-18",
                 "endDate": "2019-02-25",
                 "createdDate": "2019-02-18T11:17:42Z"
             }
         }
     ]
  }`)

func NewModifyFolderParams() parameters.ModifyFolder {
	desc := "New description"
	params := parameters.ModifyFolder{
		Metadata:     &parameters.MetadataSet{parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"}},
		AddShareds:   parameters.ContactIDSet{"KUAAAAAQ"},
		CustomFields: parameters.CustomFieldSet{parameters.CustomField{ID: "IEAAAAAQJUAAAAAX", Value: "testValue"}},
		Description:  &desc,
		Project: parameters.ProjectModify{
			OwnersAdd:    parameters.ContactIDSet{"KUAAAAAQ"},
			OwnersRemove: parameters.ContactIDSet{"KUAAAAHK"},
			Status:       types.Green,
			StartDate:    "2019-02-18",
			EndDate:      "2019-02-25",
		},
		AddParents: parameters.FolderIDSet{"IEAAAAAQI7777777"},
		Title:      "New title",
	}
	return params
}

func NewModifyFoldersParams() parameters.ModifyFolders {
	params := parameters.ModifyFolders{
		CustomFields: parameters.CustomFieldSet{
			parameters.CustomField{ID: "IEAAAAAQJUAAAAAX", Value: "testValue"},
			parameters.CustomField{ID: "IEAAAAAQJUAAAAAY", Value: "testValue"},
		},
	}
	return params
}

func Example_modifyFolder() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.ModifyFolder{}
	api.ModifyFolder("KUAAAAHP", params)
}

func Example_modifyFolders() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.ModifyFolders{}
	api.ModifyFolders([]types.FolderID{"IEAAAAAQI4AB5BGU", "IEAAAAAQI4AB5BGV"}, params)
}

func TestModifyFolder(t *testing.T) {
	client := NewMockClient(modifiedFolderData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewModifyFolderParams()
	folder, err := api.ModifyFolder("IEAAAAAQI4AB5BGU", params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPut)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/folders/IEAAAAAQI4AB5BGU")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "addParents=[\"IEAAAAAQI7777777\"]&addShareds=[\"KUAAAAAQ\"]&customFields=[{\"id\":\"IEAAAAAQJUAAAAAX\",\"value\":\"testValue\"}]&description=New description&metadata=[{\"key\":\"testMetaKey\",\"value\":\"testMetaValue\"}]&project={\"ownerAdd\":[\"KUAAAAAQ\"],\"ownerRemove\":[\"KUAAAAHK\"],\"status\":\"Green\",\"startDate\":\"2019-02-18\",\"endDate\":\"2019-02-25\"}&title=New title")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(folder.Data[0].ID), "IEAAAAAQI4AB5BGU")
}

func TestModifyFolders(t *testing.T) {
	client := NewMockClient(modifiedFolderData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewModifyFoldersParams()
	folder, err := api.ModifyFolders([]types.FolderID{"IEAAAAAQI4AB5BGU", "IEAAAAAQI4AB5BGV"}, params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPut)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/folders/IEAAAAAQI4AB5BGU,IEAAAAAQI4AB5BGV")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "customFields=[{\"id\":\"IEAAAAAQJUAAAAAX\",\"value\":\"testValue\"},{\"id\":\"IEAAAAAQJUAAAAAY\",\"value\":\"testValue\"}]")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(folder.Data[0].ID), "IEAAAAAQI4AB5BGU")
}
