package wrike

import (
	"fmt"

	parameters "github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// QueryGroup fetches a group by id.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-groups
func (api *API) QueryGroup(id string, params *parameters.QueryGroup) (*types.Groups, error) {
	path := fmt.Sprintf("/groups/%s", id)

	qp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.get(path, &qp)
	if err != nil {
		return nil, err
	}

	return types.NewGroupsFromJSON(data)
}

// QueryGroups fetches a list of groups.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-groups
func (api *API) QueryGroups(params *parameters.QueryGroups) (*types.Groups, error) {
	path := "/groups"

	qp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.get(path, &qp)
	if err != nil {
		return nil, err
	}

	return types.NewGroupsFromJSON(data)
}
