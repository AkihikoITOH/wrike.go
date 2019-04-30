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

var createdFolderData = []byte(
	`{
     "kind": "folders",
     "data":
     [
         {
             "id": "IEAAAAAQI4AB5BGV",
             "accountId": "IEAAAAAQ",
             "title": "Test folder",
             "createdDate": "2019-02-18T11:17:42Z",
             "updatedDate": "2019-02-18T11:17:42Z",
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

func NewCreateFolderParams() parameters.CreateFolder {
	return parameters.CreateFolder{
		Metadata:     &parameters.MetadataSet{parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"}},
		CustomFields: parameters.CustomFieldSet{parameters.CustomField{ID: "IEAAAAAQJUAAAAAX", Value: "testValue"}},
		Description:  "Test description",
		Project:      parameters.Project{OwnerIDs: parameters.ContactIDSet{"KUAAAAAQ"}, Status: types.Green, StartDate: "2019-02-18", EndDate: "2019-02-25"},
		Title:        "Test folder",
		Shareds:      parameters.ContactIDSet{"KUAAAAAQ"},
	}
}

func Example_createFolder() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.CreateFolder{
		Metadata:     &parameters.MetadataSet{parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"}},
		CustomFields: parameters.CustomFieldSet{parameters.CustomField{ID: "IEAAAAAQJUAAAAAX", Value: "testValue"}},
		Description:  "Test description",
		Project:      parameters.Project{OwnerIDs: parameters.ContactIDSet{"KUAAAAAQ"}, Status: types.Green, StartDate: "2019-02-18", EndDate: "2019-02-25"},
		Title:        "Test folder",
		Shareds:      parameters.ContactIDSet{"KUAAAAAQ"},
	}
	api.CreateFolder("IEAAAAAQI4AB5BGU", params)
}

func TestCreateFolder(t *testing.T) {
	client := NewMockClient(createdFolderData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewCreateFolderParams()
	folders, err := api.CreateFolder("IEAAAAAQI4AB5BGU", params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPost)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/folders/IEAAAAAQI4AB5BGU/folders")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "customFields=[{\"id\":\"IEAAAAAQJUAAAAAX\",\"value\":\"testValue\"}]&description=Test description&metadata=[{\"key\":\"testMetaKey\",\"value\":\"testMetaValue\"}]&project={\"ownerIds\":[\"KUAAAAAQ\"],\"status\":\"Green\",\"startDate\":\"2019-02-18\",\"endDate\":\"2019-02-25\"}&shareds=[\"KUAAAAAQ\"]&title=Test folder")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(folders.Data[0].ID), "IEAAAAAQI4AB5BGV")
}
