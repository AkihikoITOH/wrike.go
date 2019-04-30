package wrike

import (
	"fmt"

	types "github.com/AkihikoITOH/wrike.go/types"
)

// QueryUser fetches a User by id.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-user
func (api *API) QueryUser(id string) (*types.Contacts, error) {
	path := fmt.Sprintf("/users/%s", id)

	data, err := api.get(path, nil)
	if err != nil {
		return nil, err
	}

	return types.NewContactsFromJSON(data)
}
