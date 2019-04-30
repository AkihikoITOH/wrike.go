package wrike

import (
	"fmt"
	"strings"

	types "github.com/AkihikoITOH/wrike.go/types"
)

// QueryCustomFields fetches all custom fields in the account.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-custom-fields
func (api *API) QueryCustomFields() (*types.CustomFields, error) {
	path := "/customfields"

	data, err := api.get(path, nil)
	if err != nil {
		return nil, err
	}

	return types.NewCustomFieldsFromJSON(data)
}

// QueryCustomFieldsByIDs fetches all custom fields by the given IDs.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-custom-fields
func (api *API) QueryCustomFieldsByIDs(ids []string) (*types.CustomFields, error) {
	path := fmt.Sprintf("/customfields/%s", strings.Join(ids, ","))

	data, err := api.get(path, nil)
	if err != nil {
		return nil, err
	}

	return types.NewCustomFieldsFromJSON(data)
}
