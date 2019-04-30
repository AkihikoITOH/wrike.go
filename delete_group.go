package wrike

import (
	"fmt"

	types "github.com/AkihikoITOH/wrike.go/types"
)

// DeleteGroup deletes an existing group by given ID.
// For details refer to https://developers.wrike.com/documentation/api/methods/delete-groups
func (api *API) DeleteGroup(id string) (*types.Groups, error) {
	path := fmt.Sprintf("/groups/%s", id)

	data, err := api.delete(path)
	if err != nil {
		return nil, err
	}

	return types.NewGroupsFromJSON(data)
}
