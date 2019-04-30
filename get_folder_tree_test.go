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

var folderTreeData = []byte(
	`{
     "kind": "folderTree",
     "data":
     [
         {
             "id": "IEAAAAAQI7777777",
             "title": "Root",
             "children":
             [
                 "IEAAAAAQI4AB5BEM",
                 "IEAAAAAQI4AB5BEK",
                 "IEAAAAAQI4AB5BEL"
             ],
             "childIds":
             [
                 "IEAAAAAQI4AB5BEM",
                 "IEAAAAAQI4AB5BEK",
                 "IEAAAAAQI4AB5BEL"
             ],
             "scope": "WsRoot"
         },
         {
             "id": "IEAAAAAQI7777776",
             "title": "Recycle Bin",
             "children":
             [
                 "IEAAAAAQI4AB5BGN",
                 "IEAAAAAQI4AB5BGK",
                 "IEAAAAAQI4AB5BGD"
             ],
             "childIds":
             [
                 "IEAAAAAQI4AB5BGN",
                 "IEAAAAAQI4AB5BGK",
                 "IEAAAAAQI4AB5BGD"
             ],
             "scope": "RbRoot"
         },
         {
             "id": "IEAAAAAQI4AB5BEI",
             "title": "Welcome to Wrike!",
             "children":
             [
             ],
             "childIds":
             [
             ],
             "scope": "WsFolder"
         }
     ]
  }`)

var folderSubtreeData = []byte(
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

func NewGetFolderTreeParams() parameters.GetFolderTree {
	return parameters.GetFolderTree{}
}

func NewGetFolderSubtreeParams() parameters.GetFolderSubtree {
	return parameters.GetFolderSubtree{}
}

func Example_getFolderTree() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.GetFolderTree{}
	api.GetFolderTree(&params)
}

func Example_getFolderSubtree() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.GetFolderSubtree{}
	api.GetFolderSubtree("IEAAAAAQI4AB5BGU", &params)
}

func TestGetFolderTree(t *testing.T) {
	client := NewMockClient(folderTreeData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewGetFolderTreeParams()
	folders, err := api.GetFolderTree(&params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	data, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, data, "https://app-eu.wrike.com/api/v4/folders?")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(folders.Data), 3)
	assert.Equal(t, string(folders.Data[0].ID), "IEAAAAAQI7777777")
}

func TestGetFolderSubtree(t *testing.T) {
	client := NewMockClient(folderSubtreeData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewGetFolderSubtreeParams()
	folders, err := api.GetFolderSubtree("IEAAAAAQI4AB5BGU", &params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	data, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, data, "https://app-eu.wrike.com/api/v4/folders/IEAAAAAQI4AB5BGU/folders?")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(folders.Data), 3)
	assert.Equal(t, string(folders.Data[0].ID), "IEAAAAAQI4AB5BGU")
}
