package wrike_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
	parameters "github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
)

var createdTaskData = []byte(
	`{
   "kind": "tasks",
   "data":
   [
       {
           "id": "IEAAAAAQKQAB5BG5",
           "accountId": "IEAAAAAQ",
           "title": "Test task",
           "description": "Test task description",
           "briefDescription": "Test task description",
           "parentIds":
           [
               "IEAAAAAQI4AB5BGU"
           ],
           "superParentIds":
           [
               "IEAAAAAQI4AB5BGV",
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
           "status": "Active",
           "importance": "High",
           "createdDate": "2019-02-18T11:17:43Z",
           "updatedDate": "2019-02-18T11:17:43Z",
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
           "permalink": "https://dev.wrke.io/open.htm?id=2000093",
           "priority": "00000800800000000000b100",
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
       }
   ]
}`)

func NewCreateTaskParams() parameters.CreateTask {
	desc := "Test task description"
	t := true
	return parameters.CreateTask{
		SuperTasks:     parameters.TaskIDSet{"IEAAAAAQKQAB5BG4"},
		Metadata:       &parameters.MetadataSet{parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"}},
		PriorityBefore: types.TaskID("IEAAAAAQKQAB5BG3"),
		Importance:     types.HighImportance,
		CustomFields:   parameters.CustomFieldSet{parameters.CustomField{ID: "IEAAAAAQJUAAAAAX", Value: "testValue"}},
		Description:    &desc,
		Dates:          &parameters.TaskDates{Start: "2019-02-18", Due: "2019-02-25"},
		Title:          "Test task",
		Follow:         &t,
		Followers:      parameters.ContactIDSet{"KUAAAAAQ"},
		Responsibles:   parameters.ContactIDSet{"KUAAAAAQ"},
		Shareds:        parameters.ContactIDSet{"KUAAAAAQ"},
		Parents:        parameters.FolderIDSet{"IEAAAAAQI4AB5BGU"},
		Status:         parameters.TaskStatusSet{types.ActiveTaskStatus},
	}
}

func Example_createTask() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.CreateTask{}
	api.CreateTask(types.FolderID("IEAAAAAQI4AB5BGU"), params)
}

func TestCreateTask(t *testing.T) {
	client := NewMockClient(createdTaskData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewCreateTaskParams()
	tasks, err := api.CreateTask(types.FolderID("IEAAAAAQI4AB5BGU"), params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPost)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/folders/IEAAAAAQI4AB5BGU/tasks")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "customFields=[{\"id\":\"IEAAAAAQJUAAAAAX\",\"value\":\"testValue\"}]&dates={\"start\":\"2019-02-18\",\"due\":\"2019-02-25\"}&description=Test task description&follow=true&followers=[\"KUAAAAAQ\"]&importance=High&metadata=[{\"key\":\"testMetaKey\",\"value\":\"testMetaValue\"}]&parents=[\"IEAAAAAQI4AB5BGU\"]&priorityBefore=IEAAAAAQKQAB5BG3&responsibles=[\"KUAAAAAQ\"]&shareds=[\"KUAAAAAQ\"]&status=[\"Active\"]&superTasks=[\"IEAAAAAQKQAB5BG4\"]&title=Test task")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(tasks.Data[0].ID), "IEAAAAAQKQAB5BG5")
}
