package wrike

import (
	"encoding/json"
	"fmt"
	"strings"

	types "github.com/AkihikoITOH/wrike.go/types"
)

// QueryContactsParams contains parameters that will be passed to QueryContacts API.
type QueryContactsParams struct {
	Me       bool
	Metadata *types.Metadata
	Deleted  bool
	Fields   []string
}

// ToQueryParams converts the QueryContactsParams object to a query params string.
func (qp *QueryContactsParams) ToQueryParams() string {
	params := make([]string, 0)
	if qp.Me {
		params = append(params, "me=true")
	}
	if qp.Metadata != nil {
		metadata, err := json.Marshal(qp.Metadata)
		if err == nil {
			params = append(params, fmt.Sprintf("metadata=%s", string(metadata)))
		}
	}
	if qp.Deleted {
		params = append(params, "deleted=true")
	}
	if len(qp.Fields) > 0 {
		params = append(params, fmt.Sprintf("fields=[\"%s\"]", strings.Join(qp.Fields, "\",\"")))
	}

	return strings.Join(params, "&")
}

// QueryContacts fetches a list of contacts.
// For details refer to https://developers.wrike.com/documentation/api/methods/query-contacts
func (api *API) QueryContacts(params *QueryContactsParams) (*types.Contacts, error) {
	path := "/contacts"

	if params != nil {
		path += fmt.Sprintf("?%s", params.ToQueryParams())
	}

	data, err := api.get(path)
	if err != nil {
		return nil, err
	}

	return types.NewContactsFromJSON(data)
}
