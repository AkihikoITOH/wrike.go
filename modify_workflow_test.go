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

var modifiedWorkflowData = []byte(
	`{
     "kind": "workflows",
     "data":
     [
         {
             "id": "IEAAAAAQK4AAAAFU",
             "name": "New workflow",
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
                     "id": "IEAAAAAQJMAAAAF6",
                     "name": "In design",
                     "standardName": false,
                     "color": "Green",
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

func NewModifyWorkflowParams() parameters.ModifyWorkflow {
	statusName := "In design"
	color := "Green"
	status := "Active"
	customStatus := parameters.CustomStatus{
		Name:  &statusName,
		Color: &color,
		Group: &status,
	}
	name := "New workflow"
	hidden := false
	params := parameters.ModifyWorkflow{
		Name:         &name,
		Hidden:       &hidden,
		CustomStatus: &customStatus,
	}
	return params
}

func Example_modifyWorkflow() {
	conf := wrike.NewConfig("wrike-access-token", "")
	api := wrike.NewAPI(conf)
	statusName := "In design"
	color := "Green"
	status := "Active"
	customStatus := parameters.CustomStatus{
		Name:  &statusName,
		Color: &color,
		Group: &status,
	}
	name := "New workflow"
	hidden := false
	params := parameters.ModifyWorkflow{
		Name:         &name,
		Hidden:       &hidden,
		CustomStatus: &customStatus,
	}
	api.ModifyWorkflow("IEAAAAAQK4AAAAFU", params)
}

func TestModifyWorkflow(t *testing.T) {
	id := "IEAAAAAQK4AAAAFU"
	client := NewMockClient(modifiedWorkflowData)
	api := &wrike.API{Config: mockAPIConfig, HTTPClient: client}
	params := NewModifyWorkflowParams()
	workflow, err := api.ModifyWorkflow(id, params)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodPut)
	assert.Equal(t, client.Request.URL.String(), "https://app-eu.wrike.com/api/v4/workflows/IEAAAAAQK4AAAAFU")
	body, _ := ioutil.ReadAll(client.Request.Body)
	data, _ := url.QueryUnescape(string(body))
	assert.Equal(t, data, "customStatus={\"name\":\"In design\",\"color\":\"Green\",\"group\":\"Active\"}&hidden=false&name=New workflow")
	SharedRequestTests(t, client.Request)

	assert.Equal(t, string(workflow.Data[0].ID), "IEAAAAAQK4AAAAFU")
}
