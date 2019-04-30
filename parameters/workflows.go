package parameters

import (
	"net/url"

	types "github.com/AkihikoITOH/wrike.go/types"
)

// CreateWorkflow contains parameters that will be passed to CreateWorkflow API.
type CreateWorkflow struct {
	Name string `url:"name"`
}

// CustomStatus contains custom status data to be added/modified.
// Note that the fields are pointers so it can properly ignore blank values.
type CustomStatus struct {
	ID           *types.CustomStatusID `json:"id,omitempty"`
	Name         *string               `json:"name,omitempty"`
	StandardName *bool                 `json:"standardName,omitempty"`
	Color        *string               `json:"color,omitempty"`
	Standard     *bool                 `json:"standard,omitempty"`
	Group        *string               `json:"group,omitempty"`
	Hidden       *bool                 `json:"hidden,omitempty"`
}

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (cs *CustomStatus) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(cs, key, v)
}

// ModifyWorkflow contains parameters that will be passed to ModifyWorkflow API.
type ModifyWorkflow struct {
	Name         *string       `url:"name,omitempty"`
	Hidden       *bool         `url:"hidden,omitempty"`
	CustomStatus *CustomStatus `url:"customStatus,omitempty"`
}
