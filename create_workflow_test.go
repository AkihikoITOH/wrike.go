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
)

var createdWorkflowData = []byte(
	`{
     "kind": "workflows",
     "data":
     [
         {
             "id": "IEAAAAAQK4AAAAFU",
             "name": "Test workflow",
             "standard": false,
             "hidden": false,
             "customStatuses":
             [
                 {
                     "id": "IEAAAAAQJMAAAAFU",
                     "name": "Active",
                     "standardName": false,
                     "color": "Blue",
                     "standard": false,
                     "group": "Active",
                     "hidden": false
                 },
                 {
                     "id": "IEAAAAAQJMAAAAFV",
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

func NewCreateWorkflowParams() parameters.CreateWorkflow {
	return parameters.CreateWorkflow{Name: "Test workflow"}
}

func Example_createWorkflow() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	params := parameters.CreateWorkflow{Name: "Test workflow"}
	api.CreateWorkflow(params)
}

func TestCreateWorkflow(t *testing.T) {
	client := NewMockClient(createdWorkflowData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewCreateWorkflowParams()
	workflows, err := api.CreateWorkflow(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPost)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/workflows")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "name=Test workflow")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(workflows.Data[0].ID), "IEAAAAAQK4AAAAFU")
}
