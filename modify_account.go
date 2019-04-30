package wrike

import (
	"github.com/AkihikoITOH/wrike.go/parameters"
	types "github.com/AkihikoITOH/wrike.go/types"
	"github.com/google/go-querystring/query"
)

// ModifyAccount modifies the current account with given parameters.
// For details refer to https://developers.wrike.com/documentation/api/methods/modify-account
func (api *API) ModifyAccount(params *parameters.ModifyAccount) (*types.Accounts, error) {
	path := "/account"

	body, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	data, err := api.put(path, &body)
	if err != nil {
		return nil, err
	}

	return types.NewAccountsFromJSON(data)
}
