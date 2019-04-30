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

var copiedFolderData = []byte(
	`{
     "kind": "folders",
     "data":
     [
         {
             "id": "IEAAAAAQI4AB5BGY",
             "accountId": "IEAAAAAQ",
             "title": "Test folder copy",
             "createdDate": "2019-02-18T11:17:43Z",
             "updatedDate": "2019-02-18T11:17:43Z",
             "description": "",
             "sharedIds":
             [
                 "KUAAAAHK",
                 "KUAAAAAQ"
             ],
             "parentIds":
             [
                 "IEAAAAAQI4AB5BGX"
             ],
             "childIds":
             [
                 "IEAAAAAQI4AB5BGZ",
                 "IEAAAAAQI4AB5BG2"
             ],
             "superParentIds":
             [
             ],
             "scope": "WsFolder",
             "hasAttachments": false,
             "permalink": "https://dev.wrke.io/open.htm?id=2000088",
             "workflowId": "IEAAAAAQK777777Q",
             "metadata":
             [
             ],
             "customFields":
             [
             ]
         }
     ]
  }`)

func NewCopyFolderParams() parameters.CopyFolder {
	f := false
	t := true
	entryLimit := 50
	return parameters.CopyFolder{
		Parent:             "IEAAAAAQI4AB5BGX",
		CopyParents:        &f,
		EntryLimit:         &entryLimit,
		CopyDescriptions:   &t,
		CopyCustomFields:   &t,
		CopyCustomStatuses: &t,
		CopyResponsibles:   &t,
		RemoveResponsibles: parameters.ContactIDSet{"KUAAAAAQ"},
		AddResponsibles:    parameters.ContactIDSet{"KUAAAAHK"},
		RescheduleMode:     "Start",
		RescheduleDate:     "2019-02-18",
		Title:              "Test folder copy",
	}
}

func Example_copyFolder() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	f := false
	t := true
	entryLimit := 50
	params := parameters.CopyFolder{
		Parent:             "IEAAAAAQI4AB5BGX",
		CopyParents:        &f,
		EntryLimit:         &entryLimit,
		CopyDescriptions:   &t,
		CopyCustomFields:   &t,
		CopyCustomStatuses: &t,
		CopyResponsibles:   &t,
		RemoveResponsibles: parameters.ContactIDSet{"KUAAAAAQ"},
		AddResponsibles:    parameters.ContactIDSet{"KUAAAAHK"},
		RescheduleMode:     "Start",
		RescheduleDate:     "2019-02-18",
		Title:              "Test folder copy",
	}
	api.CopyFolder("IEAAAAAQI4AB5BGU", params)
}

func TestCopyFolder(t *testing.T) {
	client := NewMockClient(copiedFolderData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewCopyFolderParams()
	folders, err := api.CopyFolder("IEAAAAAQI4AB5BGU", params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPost)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/copy_folder/IEAAAAAQI4AB5BGU")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "addResponsibles=[\"KUAAAAHK\"]&copyCustomFields=true&copyCustomStatuses=true&copyDescriptions=true&copyParents=false&copyResponsibles=true&entryLimit=50&parent=IEAAAAAQI4AB5BGX&removeResponsibles=[\"KUAAAAAQ\"]&rescheduleDate=2019-02-18&rescheduleMode=Start&title=Test folder copy")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(folders.Data[0].ID), "IEAAAAAQI4AB5BGY")
}
