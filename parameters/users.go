package parameters

import (
	"net/url"
)

// Profile contains profile data to be modified.
// Note that the fields are pointers so it can properly ignore blank values.
type Profile struct {
	AccountID *string `json:"accountId,omitempty"`
	Role      *string `json:"role,omitempty"`
	External  *bool   `json:"external,omitempty"`
}

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (p *Profile) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(p, key, v)
}

// ModifyUser contains parameters that will be passed to ModifyUser API.
type ModifyUser struct {
	Profile *Profile `url:"profile,omitempty"`
}
