package types_test

import (
	"fmt"
	"testing"

	"gotest.tools/assert"

	"github.com/AkihikoITOH/wrike.go/types"
)

var taskData = []byte(
	`{
     "kind": "tasks",
     "data":
     [
         {
             "id": "IEAAAAAQKQAB5BG4",
             "accountId": "IEAAAAAQ",
             "title": "Test task",
             "description": "Test task description",
             "briefDescription": "Test task description",
             "parentIds":
             [
                 "IEAAAAAQI4AB5BGW",
                 "IEAAAAAQI4AB5BGU"
             ],
             "superParentIds":
             [
                 "IEAAAAAQI4AB5BGV"
             ],
             "sharedIds":
             [
                 "KUAAAAAQ"
             ],
             "responsibleIds":
             [
                 "KUAAAAAQ"
             ],
             "status": "Active",
             "importance": "High",
             "createdDate": "2019-02-18T11:17:43Z",
             "updatedDate": "2019-02-18T11:17:44Z",
             "dates":
             {
                 "type": "Planned",
                 "duration": 2880,
                 "start": "2019-02-14T09:00:00",
                 "due": "2019-02-21T17:00:00"
             },
             "scope": "WsTask",
             "authorIds":
             [
                 "KUAAAAAQ"
             ],
             "customStatusId": "IEAAAAAQJMAAAAAA",
             "hasAttachments": false,
             "attachmentCount": 0,
             "permalink": "https://dev.wrke.io/open.htm?id=2000092",
             "priority": "00000800800000000000b200",
             "superTaskIds":
             [
                 "IEAAAAAQKQAB5BG3"
             ],
             "subTaskIds":
             [
                 "IEAAAAAQKQAB5BG5"
             ],
             "dependencyIds":
             [
                 "IEAAAAAQIUAB5BG3KMAB5BG4"
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
                     "id": "IEAAAAAQJUAAAAAX",
                     "value": "testValue"
                 }
             ]
         },
         {
             "id": "IEAAAAAQKQAB5BG3",
             "accountId": "IEAAAAAQ",
             "title": "Test task",
             "description": "",
             "briefDescription": "",
             "parentIds":
             [
                 "IEAAAAAQI4AB5BGV"
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
             ],
             "status": "Active",
             "importance": "Normal",
             "createdDate": "2019-02-18T11:17:43Z",
             "updatedDate": "2019-02-18T11:17:44Z",
             "dates":
             {
                 "type": "Planned",
                 "duration": 2880,
                 "start": "2019-02-14T09:00:00",
                 "due": "2019-02-21T17:00:00"
             },
             "scope": "WsTask",
             "authorIds":
             [
                 "KUAAAAAQ"
             ],
             "customStatusId": "IEAAAAAQJMAAAAAA",
             "hasAttachments": false,
             "attachmentCount": 0,
             "permalink": "https://dev.wrke.io/open.htm?id=2000091",
             "priority": "00000800800000000000b000",
             "superTaskIds":
             [
             ],
             "subTaskIds":
             [
                 "IEAAAAAQKQAB5BG4"
             ],
             "dependencyIds":
             [
                 "IEAAAAAQIUAB5BG3KMAB5BG4",
                 "IEAAAAAQIUAB5BG5KMAB5BG3"
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
                     "id": "IEAAAAAQJUAAAAAX",
                     "value": "testValue"
                 },
                 {
                     "id": "IEAAAAAQJUAAAAAY",
                     "value": "testValue"
                 }
             ]
         }
     ]
  }`)

func TestNewTasksFromJSON(t *testing.T) {
	tasks, err := types.NewTasksFromJSON(taskData)

	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Assert(t, err == nil)
	data := tasks.Data
	assert.Equal(t, len(data), 2)
	assert.Equal(t, string(data[0].ID), "IEAAAAAQKQAB5BG4")
}
