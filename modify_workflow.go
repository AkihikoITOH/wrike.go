package wrike

import (
	"fmt"

	parameters "github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// ModifyWorkflow modifiies a workflow with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/modify-workflow
func (api *API) ModifyWorkflow(id string, params parameters.ModifyWorkflow) (*types.Workflows, error) {
	path := fmt.Sprintf("/workflows/%s", id)

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.put(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewWorkflowsFromJSON(data)
}
