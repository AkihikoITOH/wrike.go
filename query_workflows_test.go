package wrike_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
)

var allWorkflowsData = []byte(
	`{
     "kind": "workflows",
     "data":
     [
         {
             "id": "IEAAAAAQK777777Q",
             "name": "Default Workflow",
             "standard": true,
             "hidden": false,
             "customStatuses":
             [
                 {
                     "id": "IEAAAAAQJMAAAAAA",
                     "name": "New",
                     "standardName": true,
                     "color": "Blue",
                     "standard": true,
                     "group": "Active",
                     "hidden": false
                 },
                 {
                     "id": "IEAAAAAQJMAAAAAK",
                     "name": "In Progress",
                     "standardName": true,
                     "color": "Turquoise",
                     "standard": false,
                     "group": "Active",
                     "hidden": false
                 },
                 {
                     "id": "IEAAAAAQJMAAAAAB",
                     "name": "Completed",
                     "standardName": true,
                     "color": "Green",
                     "standard": true,
                     "group": "Completed",
                     "hidden": false
                 }
             ]
         },
         {
             "id": "IEAAAAAQK4AAAAEM",
             "name": "New workflow",
             "standard": false,
             "hidden": false,
             "customStatuses":
             [
                 {
                     "id": "IEAAAAAQJMAAAAEM",
                     "name": "Active",
                     "standardName": false,
                     "color": "Blue",
                     "standard": false,
                     "group": "Active",
                     "hidden": false
                 },
                 {
                     "id": "IEAAAAAQJMAAAAEW",
                     "name": "In design",
                     "standardName": false,
                     "color": "Green",
                     "standard": false,
                     "group": "Active",
                     "hidden": false
                 },
                 {
                     "id": "IEAAAAAQJMAAAAEN",
                     "name": "Completed",
                     "standardName": false,
                     "color": "Green",
                     "standard": false,
                     "group": "Completed",
                     "hidden": false
                 }
             ]
         },
         {
             "id": "IEAAAAAQK4AAAAFA",
             "name": "New workflow",
             "standard": false,
             "hidden": false,
             "customStatuses":
             [
                 {
                     "id": "IEAAAAAQJMAAAAFA",
                     "name": "Active",
                     "standardName": false,
                     "color": "Blue",
                     "standard": false,
                     "group": "Active",
                     "hidden": false
                 },
                 {
                     "id": "IEAAAAAQJMAAAAFK",
                     "name": "In design",
                     "standardName": false,
                     "color": "Green",
                     "standard": false,
                     "group": "Active",
                     "hidden": false
                 },
                 {
                     "id": "IEAAAAAQJMAAAAFB",
                     "name": "Completed",
                     "standardName": false,
                     "color": "Green",
                     "standard": false,
                     "group": "Completed",
                     "hidden": false
                 }
             ]
         }
     ]
  }`)

func Example_queryWorkflows() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	api.QueryWorkflows()
}

func TestQueryWorkflows(t *testing.T) {
	client := NewMockClient(allWorkflowsData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	workflows, err := api.QueryWorkflows()
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	data, _ := url.QueryUnescape(client.Request.URL.String())
	assert.Equal(t, data, "https://app-eu.wrike.com/api/v4/workflows")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(workflows.Data), 3)
}
