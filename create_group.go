package wrike

import (
	parameters "github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// CreateGroup creates a new group with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/create-groups
func (api *API) CreateGroup(params parameters.CreateGroup) (*types.Groups, error) {
	path := "/groups"

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.post(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewGroupsFromJSON(data)
}
