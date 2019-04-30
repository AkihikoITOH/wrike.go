package wrike

import (
	"fmt"

	"github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// ModifyGroup modifiies a group with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/modify-groups
func (api *API) ModifyGroup(id string, params *parameters.ModifyGroup) (*types.Groups, error) {
	path := fmt.Sprintf("/groups/%s", id)

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.put(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewGroupsFromJSON(data)
}
