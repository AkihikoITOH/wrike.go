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

var modifiedTaskData = []byte(
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
               "IEAAAAAQI4AB5BGU",
               "IEAAAAAQI4AB5BGV"
           ],
           "superParentIds":
           [
               "IEAAAAAQI4AB5BGW"
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
           "updatedDate": "2019-02-18T11:17:44Z",
           "dates":
           {
               "type": "Planned",
               "duration": 1920,
               "start": "2019-02-18T09:00:00",
               "due": "2019-02-21T17:00:00"
           },
           "scope": "WsTask",
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
               "IEAAAAAQKQAB5BG4"
           ],
           "subTaskIds":
           [
           ],
           "dependencyIds":
           [
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
                   "id": "IEAAAAAQJUAAAAAY",
                   "value": "testValue"
               }
           ]
       }
   ]
}`)

var modifiedTasksData = []byte(
	`{
   "kind": "tasks",
   "data":
   [
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
               "start": "2019-02-18T09:00:00",
               "due": "2019-02-25T17:00:00"
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
       },
       {
           "id": "IEAAAAAQKQAB5BG5",
           "accountId": "IEAAAAAQ",
           "title": "New title",
           "description": "New description",
           "briefDescription": "New description",
           "parentIds":
           [
               "IEAAAAAQI4AB5BGU",
               "IEAAAAAQI4AB5BGV"
           ],
           "superParentIds":
           [
               "IEAAAAAQI4AB5BGW"
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
           "updatedDate": "2019-02-18T11:17:44Z",
           "dates":
           {
               "type": "Planned",
               "duration": 1920,
               "start": "2019-02-18T09:00:00",
               "due": "2019-02-21T17:00:00"
           },
           "scope": "WsTask",
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
               "IEAAAAAQKQAB5BG4"
           ],
           "subTaskIds":
           [
           ],
           "dependencyIds":
           [
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

func NewModifyTaskParams() parameters.ModifyTask {
	title := "New title"
	desc := "New description"
	params := parameters.ModifyTask{
		RemoveSuperTasks: parameters.TaskIDSet{"IEAAAAAQKQAB5BG3"},
		PriorityAfter:    types.TaskID("IEAAAAAQKQAB5BG3"),
		Importance:       types.LowImportance,
		CustomFields:     parameters.CustomFieldSet{parameters.CustomField{ID: "IEAAAAAQJUAAAAAX"}, parameters.CustomField{ID: "IEAAAAAQJUAAAAAY", Value: "testValue"}},
		Description:      &desc,
		AddResponsibles:  parameters.ContactIDSet{"KUAAAAAQ"},
		Dates:            &parameters.TaskDates{Start: "2019-02-18", Due: "2019-02-21"},
		AddFollowers:     parameters.ContactIDSet{"KUAAAAAQ"},
		AddParents:       parameters.FolderIDSet{"IEAAAAAQI4AB5BGV"},
		Title:            &title,
		Status:           parameters.TaskStatusSet{types.DeferredTaskStatus},
	}
	return params
}

func NewModifyTasksParams() parameters.ModifyTasks {
	params := parameters.ModifyTasks{
		CustomFields: parameters.CustomFieldSet{
			parameters.CustomField{ID: "IEAAAAAQJUAAAAAX", Value: "testValue"},
			parameters.CustomField{ID: "IEAAAAAQJUAAAAAY", Value: "testValue"},
		},
	}
	return params
}

func Example_modifyTask() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	title := "New title"
	desc := "New description"
	params := parameters.ModifyTask{
		RemoveSuperTasks: parameters.TaskIDSet{"IEAAAAAQKQAB5BG3"},
		PriorityAfter:    types.TaskID("IEAAAAAQKQAB5BG3"),
		Importance:       types.LowImportance,
		CustomFields:     parameters.CustomFieldSet{parameters.CustomField{ID: "IEAAAAAQJUAAAAAX"}, parameters.CustomField{ID: "IEAAAAAQJUAAAAAY", Value: "testValue"}},
		Description:      &desc,
		AddResponsibles:  parameters.ContactIDSet{"KUAAAAAQ"},
		Dates:            &parameters.TaskDates{Start: "2019-02-18", Due: "2019-02-21"},
		AddFollowers:     parameters.ContactIDSet{"KUAAAAAQ"},
		AddParents:       parameters.FolderIDSet{"IEAAAAAQI4AB5BGV"},
		Title:            &title,
		Status:           parameters.TaskStatusSet{types.DeferredTaskStatus},
	}
	api.ModifyTask(types.TaskID("IEAAAAAQKQAB5BG5"), params)
}

func Example_modifyTasks() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.ModifyTasks{
		CustomFields: parameters.CustomFieldSet{
			parameters.CustomField{ID: "IEAAAAAQJUAAAAAX", Value: "testValue"},
			parameters.CustomField{ID: "IEAAAAAQJUAAAAAY", Value: "testValue"},
		},
	}
	api.ModifyTasks([]types.TaskID{"IEAAAAAQKQAB5BG5", "IEAAAAAQKQAB5BG3"}, params)
}

func TestModifyTask(t *testing.T) {
	client := NewMockClient(modifiedTaskData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewModifyTaskParams()
	task, err := api.ModifyTask(types.TaskID("IEAAAAAQKQAB5BG5"), params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPut)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/tasks/IEAAAAAQKQAB5BG5")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "addFollowers=[\"KUAAAAAQ\"]&addParents=[\"IEAAAAAQI4AB5BGV\"]&addResponsibles=[\"KUAAAAAQ\"]&customFields=[{\"id\":\"IEAAAAAQJUAAAAAX\"},{\"id\":\"IEAAAAAQJUAAAAAY\",\"value\":\"testValue\"}]&dates={\"start\":\"2019-02-18\",\"due\":\"2019-02-21\"}&description=New description&importance=Low&priorityAfter=IEAAAAAQKQAB5BG3&removeSuperTasks=[\"IEAAAAAQKQAB5BG3\"]&status=[\"Deferred\"]&title=New title")
	assert.Equal(t, string(task.Data[0].ID), "IEAAAAAQKQAB5BG5")
}

func TestModifyTasks(t *testing.T) {
	client := NewMockClient(modifiedTasksData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewModifyTasksParams()
	task, err := api.ModifyTasks([]types.TaskID{"IEAAAAAQKQAB5BG5", "IEAAAAAQKQAB5BG3"}, params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPut)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/tasks/IEAAAAAQKQAB5BG5,IEAAAAAQKQAB5BG3")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "customFields=[{\"id\":\"IEAAAAAQJUAAAAAX\",\"value\":\"testValue\"},{\"id\":\"IEAAAAAQJUAAAAAY\",\"value\":\"testValue\"}]")
	assert.Equal(t, string(task.Data[0].ID), "IEAAAAAQKQAB5BG3")
}
