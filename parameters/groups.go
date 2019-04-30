package parameters

import (
	types "github.com/AkihikoITOH/wrike.go/types"
)

// CreateGroup contains parameters that will be passed to CreateGroup API.
type CreateGroup struct {
	Title    string           `url:"title,omitempty"`
	Members  ContactIDSet     `url:"members,omitempty"`
	Parent   *types.ContactID `url:"parent,omitempty"`
	Avatar   *Avatar          `url:"avatar,omitempty"`
	Metadata *MetadataSet     `url:"metadata,omitempty"`
}

// ModifyGroup contains parameters that will be passed to ModifyGroup API.
type ModifyGroup struct {
	Title         *string      `url:"title,omitempty"`
	AddMembers    ContactIDSet `url:"addMembers,omitempty"`
	RemoveMembers ContactIDSet `url:"removeMembers,omitempty"`
	Parent        *string      `url:"parent,omitempty"`
	Avatar        *Avatar      `url:"avatar,omitempty"`
	Metadata      *MetadataSet `url:"metadata,omitempty"`
}

// QueryGroup contains parameters that will be passed to QueryGroup API.
type QueryGroup struct {
	Fields *FieldSet `url:"fields,omitempty"`
}

// QueryGroups contains parameters that will be passed to QueryGroups API.
type QueryGroups struct {
	Metadata *Metadata `url:"metadata,omitempty"`
	Fields   *FieldSet `url:"fields,omitempty"`
}
