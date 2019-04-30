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

var foldersData = []byte(
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
             "briefDescription": "New description",
             "description": "New description",
             "color": "None",
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
             "briefDescription": "Test description",
             "description": "Test description",
             "color": "None",
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

func NewGetFoldersParams() parameters.GetFolders {
	return parameters.GetFolders{Fields: &parameters.FieldSet{"childIds", "briefDescription", "color"}}
}

func Example_getFolders() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.GetFolders{Fields: &parameters.FieldSet{"childIds", "briefDescription", "color"}}
	api.GetFolders([]types.FolderID{"IEAAAAAQI4AB5BGU", "IEAAAAAQI4AB5BGV"}, &params)
}

func TestGetFolders(t *testing.T) {
	client := NewMockClient(foldersData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewGetFoldersParams()
	folders, err := api.GetFolders([]types.FolderID{"IEAAAAAQI4AB5BGU", "IEAAAAAQI4AB5BGV"}, &params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	data, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, data, "https://app-eu.wrike.com/api/v4/folders/IEAAAAAQI4AB5BGU,IEAAAAAQI4AB5BGV?fields=[\"childIds\",\"briefDescription\",\"color\"]")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(folders.Data), 2)
	assert.Equal(t, string(folders.Data[0].ID), "IEAAAAAQI4AB5BGU")
}
