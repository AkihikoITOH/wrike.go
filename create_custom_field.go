package wrike

import (
	"github.com/AkihikoITOH/wrike.go/parameters"
	"github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// CreateCustomField creates a new customField with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/create-customFields
func (api *API) CreateCustomField(params parameters.CreateCustomField) (*types.CustomFields, error) {
	path := "/customFields"

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.post(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewCustomFieldsFromJSON(data)
}
