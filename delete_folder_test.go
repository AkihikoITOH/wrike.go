package wrike_test

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
)

var deletedFolderData = []byte(
	`{
     "kind": "folders",
     "data":
     [
         {
             "id": "IEAAAAAQI4AB5BGU",
             "accountId": "IEAAAAAQ",
             "title": "New title",
             "createdDate": "2019-02-18T11:17:42Z",
             "updatedDate": "2019-02-18T11:17:46Z",
             "description": "New description",
             "sharedIds":
             [
                 "KUAAAAAQ"
             ],
             "parentIds":
             [
                 "IEAAAAAQI7777776"
             ],
             "childIds":
             [
                 "IEAAAAAQI4AB5BGV",
                 "IEAAAAAQI4AB5BGW"
             ],
             "superParentIds":
             [
             ],
             "scope": "RbFolder",
             "hasAttachments": false,
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
         }
     ]
  }`)

func Example_deleteFolder() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	api.DeleteFolder("IEAAAAAQI4AB5BGU")
}

func TestDeleteFolder(t *testing.T) {
	client := NewMockClient(deletedFolderData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	folders, err := api.DeleteFolder("IEAAAAAQI4AB5BGU")
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodDelete)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/folders/IEAAAAAQI4AB5BGU")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(folders.Data[0].ID), "IEAAAAAQI4AB5BGU")
}
