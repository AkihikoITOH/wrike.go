// Package parameters contains types that represent sets of parameters that are passed to the API.
package parameters

import (
	"encoding/json"
	"net/url"

	"github.com/AkihikoITOH/wrike.go/types"
)

type FolderIDSet []types.FolderID

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (ids FolderIDSet) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(ids, key, v)
}

type ContactIDSet []types.ContactID

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (ids ContactIDSet) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(ids, key, v)
}

func (ids ContactIDSet) ToSlice() []string {
	s := make([]string, 0)
	for _, id := range ids {
		s = append(s, string(id))
	}
	return s
}

// Metadata represents a metadata, part of other API objects.
type Metadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (m *Metadata) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(m, key, v)
}

// MetadataSet represents a list of pointers to Metadata
type MetadataSet []Metadata

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (s *MetadataSet) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(s, key, v)
}

type FieldSet []types.Field

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f *FieldSet) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

type Avatar struct {
	Letters string `json:"letters"`
	Color   string `json:"color"`
}

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (a *Avatar) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(a, key, v)
}

// jsonAsURLParams marshals an object into JSON and set it as an url value with the given key.
func jsonAsURLParams(obj interface{}, key string, v *url.Values) error {
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	v.Set(key, string(b))
	return nil
}
