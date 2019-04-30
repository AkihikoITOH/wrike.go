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

var allTasksData = []byte(
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

var tasksInFolderData = []byte(
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
       },
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
       }
   ]
}`)

var tasksByIDsData = []byte(
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

func NewQueryTasksParams() parameters.QueryTasks {
	t := true
	return parameters.QueryTasks{
		Metadata:      &parameters.Metadata{Key: "testMetaKey", Value: "testMetaValue"},
		CreatedDate:   &parameters.DateOrRange{Start: "2019-02-17T11:17:45Z"},
		DueDate:       &parameters.DateOrRange{Start: "2019-01-18"},
		SortOrder:     parameters.Desc,
		SortField:     parameters.CreatedDateSortField,
		ScheduledDate: &parameters.DateOrRange{Start: "2019-01-18", End: "2019-03-18"},
		UpdatedDate:   &parameters.DateOrRange{Start: "2019-02-17T11:17:45Z", End: "2019-02-19T11:17:45Z"},
		Fields: parameters.TaskFieldSet{
			parameters.AuthorIDsField,
			parameters.DependencyIDsField,
			parameters.SubTaskIDsField,
			parameters.SuperParentIDsField,
			parameters.SuperTaskIDsField,
			parameters.MetadataField,
			parameters.RecurrentField,
			parameters.AttachmentCountField,
			parameters.ResponsibleIDsField,
			parameters.DescriptionField,
			parameters.CustomFieldsField,
			parameters.BriefDescriptionField,
			parameters.HasAttachmentsField,
			parameters.ParentIDsField,
			parameters.SharedIDsField,
		},
		Descendants: &t,
		StartDate:   &parameters.DateOrRange{End: "2019-03-18T14:17:45"},
		Status:      parameters.TaskStatusSet{types.ActiveTaskStatus},
	}
}

func NewQueryTasksInFolderParams() parameters.QueryTasks {
	t := true
	return parameters.QueryTasks{
		CreatedDate:   &parameters.DateOrRange{Start: "2019-02-17T11:17:45Z", End: "2019-02-19T11:17:45Z"},
		DueDate:       &parameters.DateOrRange{End: "2019-03-18"},
		SortField:     parameters.CreatedDateSortField,
		ScheduledDate: &parameters.DateOrRange{Equal: "2019-02-18"},
		UpdatedDate:   &parameters.DateOrRange{End: "2019-02-19T11:17:45Z"},
		Fields: parameters.TaskFieldSet{
			parameters.AuthorIDsField,
			parameters.DependencyIDsField,
			parameters.SubTaskIDsField,
			parameters.SuperParentIDsField,
			parameters.SuperTaskIDsField,
			parameters.MetadataField,
			parameters.RecurrentField,
			parameters.AttachmentCountField,
			parameters.ResponsibleIDsField,
			parameters.DescriptionField,
			parameters.CustomFieldsField,
			parameters.BriefDescriptionField,
			parameters.HasAttachmentsField,
			parameters.ParentIDsField,
			parameters.SharedIDsField,
		},
		Descendants: &t,
		StartDate:   &parameters.DateOrRange{Start: "2019-01-18", End: "2019-03-18"},
		Status:      parameters.TaskStatusSet{types.ActiveTaskStatus},
	}
}

func NewQueryTasksByIDsParams() parameters.QueryTasksByIDs {
	return parameters.QueryTasksByIDs{}
}

func Example_queryTasks() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.QueryTasks{}
	api.QueryTasks(params)
}

func Example_queryTasksInFolder() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.QueryTasks{}
	api.QueryTasksInFolder(types.FolderID("IEAAAAAQI4AB5BGU"), params)
}

func Example_queryTasksByIDs() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.QueryTasksByIDs{}
	api.QueryTasksByIDs([]types.TaskID{"IEAAAAAQKQAB5BG3", "IEAAAAAQKQAB5BG5"}, params)
}

func TestQueryTasks(t *testing.T) {
	client := NewMockClient(allTasksData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewQueryTasksParams()
	tasks, err := api.QueryTasks(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	data, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, data, "https://app-eu.wrike.com/api/v4/tasks?createdDate={\"start\":\"2019-02-17T11:17:45Z\"}&descendants=true&dueDate={\"start\":\"2019-01-18\"}&fields=[\"authorIds\",\"dependencyIds\",\"subTaskIds\",\"superParentIds\",\"superTaskIds\",\"metadata\",\"recurrent\",\"attachmentCount\",\"responsibleIds\",\"description\",\"customFields\",\"briefDescription\",\"hasAttachments\",\"parentIds\",\"sharedIds\"]&metadata={\"key\":\"testMetaKey\",\"value\":\"testMetaValue\"}&scheduledDate={\"start\":\"2019-01-18\",\"end\":\"2019-03-18\"}&sortField=CreatedDate&sortOrder=Desc&startDate={\"end\":\"2019-03-18T14:17:45\"}&status=[\"Active\"]&updatedDate={\"start\":\"2019-02-17T11:17:45Z\",\"end\":\"2019-02-19T11:17:45Z\"}")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(tasks.Data), 2)
	assert.Equal(t, string(tasks.Data[0].ID), "IEAAAAAQKQAB5BG4")
}

func TestQueryTasksInFolder(t *testing.T) {
	client := NewMockClient(tasksInFolderData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewQueryTasksInFolderParams()
	tasks, err := api.QueryTasksInFolder(types.FolderID("IEAAAAAQI4AB5BGU"), params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	data, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, data, "https://app-eu.wrike.com/api/v4/folders/IEAAAAAQI4AB5BGU/tasks?createdDate={\"start\":\"2019-02-17T11:17:45Z\",\"end\":\"2019-02-19T11:17:45Z\"}&descendants=true&dueDate={\"end\":\"2019-03-18\"}&fields=[\"authorIds\",\"dependencyIds\",\"subTaskIds\",\"superParentIds\",\"superTaskIds\",\"metadata\",\"recurrent\",\"attachmentCount\",\"responsibleIds\",\"description\",\"customFields\",\"briefDescription\",\"hasAttachments\",\"parentIds\",\"sharedIds\"]&scheduledDate={\"equal\":\"2019-02-18\"}&sortField=CreatedDate&startDate={\"start\":\"2019-01-18\",\"end\":\"2019-03-18\"}&status=[\"Active\"]&updatedDate={\"end\":\"2019-02-19T11:17:45Z\"}")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(tasks.Data), 2)
	assert.Equal(t, string(tasks.Data[0].ID), "IEAAAAAQKQAB5BG3")
}

func TestQueryTasksByIDs(t *testing.T) {
	client := NewMockClient(tasksByIDsData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewQueryTasksByIDsParams()
	tasks, err := api.QueryTasksByIDs([]types.TaskID{"IEAAAAAQKQAB5BG3", "IEAAAAAQKQAB5BG5"}, params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	data, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, data, "https://app-eu.wrike.com/api/v4/tasks/IEAAAAAQKQAB5BG3,IEAAAAAQKQAB5BG5?")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(tasks.Data), 2)
	assert.Equal(t, string(tasks.Data[0].ID), "IEAAAAAQKQAB5BG3")
}
