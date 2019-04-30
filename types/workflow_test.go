package types_test

import (
	"fmt"
	"testing"

	"gotest.tools/assert"

	"github.com/AkihikoITOH/wrike.go/types"
)

var workflowData = []byte(
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

func TestNewWorkflowsFromJSON(t *testing.T) {
	workflows, err := types.NewWorkflowsFromJSON(workflowData)

	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Assert(t, err == nil)
	data := workflows.Data
	assert.Equal(t, len(data), 3)
	assert.Equal(t, string(data[0].ID), "IEAAAAAQK777777Q")
	assert.Equal(t, len([]types.CustomStatus(data[0].CustomStatuses)), 3)
}
