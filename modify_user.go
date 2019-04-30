package wrike

import (
	"fmt"

	"github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// ModifyUser modifiies a user with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/modify-user
func (api *API) ModifyUser(id string, params *parameters.ModifyUser) (*types.Contacts, error) {
	path := fmt.Sprintf("/user/%s", id)

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.put(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewContactsFromJSON(data)
}
