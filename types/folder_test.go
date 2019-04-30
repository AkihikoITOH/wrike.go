package types_test

import (
	"fmt"
	"testing"

	"gotest.tools/assert"

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

var folderTreeData = []byte(
	`{
     "kind": "folderTree",
     "data":
     [
         {
             "id": "IEAAAAAQI4AB5BGU",
             "title": "New title",
             "children":
             [
                 "IEAAAAAQI4AB5BGW",
                 "IEAAAAAQI4AB5BGV"
             ],
             "childIds":
             [
                 "IEAAAAAQI4AB5BGW",
                 "IEAAAAAQI4AB5BGV"
             ],
             "scope": "WsFolder",
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
             "id": "IEAAAAAQI4AB5BGW",
             "title": "Test folder",
             "children":
             [
             ],
             "childIds":
             [
             ],
             "scope": "WsFolder"
         },
         {
             "id": "IEAAAAAQI4AB5BGV",
             "title": "Test folder",
             "children":
             [
             ],
             "childIds":
             [
             ],
             "scope": "WsFolder",
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

func TestNewFolderTreeFromJSON(t *testing.T) {
	folders, err := types.NewFolderTreeFromJSON(folderTreeData)

	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Assert(t, err == nil)
	data := folders.Data
	assert.Equal(t, string(data[0].ID), "IEAAAAAQI4AB5BGU")
}

func TestNewFoldersFromJSON(t *testing.T) {
	folders, err := types.NewFoldersFromJSON(foldersData)

	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Assert(t, err == nil)
	data := folders.Data
	assert.Equal(t, len(data), 2)
	assert.Equal(t, string(data[0].ID), "IEAAAAAQI4AB5BGU")
}
