package wrike_test

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
	"github.com/AkihikoITOH/wrike.go/types"
)

var deletedTaskData = []byte(
	`{
   "kind": "tasks",
   "data":
   [
       {
           "id": "IEAAAAAQKQAB5BG5",
           "accountId": "IEAAAAAQ",
           "title": "New title",
           "description": "New description",
           "briefDescription": "New description",
           "parentIds":
           [
               "IEAAAAAQI7777776"
           ],
           "superParentIds":
           [
           ],
           "sharedIds":
           [
               "KUAAAAAQ"
           ],
           "responsibleIds":
           [
               "KUAAAAAQ"
           ],
           "status": "Deferred",
           "importance": "Low",
           "createdDate": "2019-02-18T11:17:43Z",
           "updatedDate": "2019-02-18T11:17:45Z",
           "dates":
           {
               "type": "Planned",
               "duration": 1920,
               "start": "2019-02-18T09:00:00",
               "due": "2019-02-21T17:00:00"
           },
           "scope": "RbTask",
           "authorIds":
           [
               "KUAAAAAQ"
           ],
           "customStatusId": "IEAAAAAQJMAAAAAC",
           "hasAttachments": true,
           "attachmentCount": 1,
           "permalink": "https://dev.wrke.io/open.htm?id=2000093",
           "priority": "00000800800000000000ae00",
           "followedByMe": true,
           "followerIds":
           [
               "KUAAAAAQ"
           ],
           "superTaskIds":
           [
           ],
           "subTaskIds":
           [
           ],
           "dependencyIds":
           [
           ],
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
                   "id": "IEAAAAAQJUAAAAAY",
                   "value": "testValue"
               },
               {
                   "id": "IEAAAAAQJUAAAAAX",
                   "value": "testValue"
               }
           ]
       }
   ]
 }`)

func Example_deleteTask() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	api.DeleteTask(types.TaskID("IEAAAAAQKQAB5BG5"))
}

func TestDeleteTask(t *testing.T) {
	client := NewMockClient(deletedTaskData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	tasks, err := api.DeleteTask(types.TaskID("IEAAAAAQKQAB5BG5"))
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodDelete)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/tasks/IEAAAAAQKQAB5BG5")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(tasks.Data), 1)
}
