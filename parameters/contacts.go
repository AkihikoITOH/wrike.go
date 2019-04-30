package parameters

// QueryContacts contains parameters that will be passed to QueryContacts API.
type QueryContacts struct {
	Me       *bool     `url:"me,omitempty"`
	Metadata *Metadata `url:"metadata,omitempty"`
	Deleted  *bool     `url:"deleted,omitempty"`
	Fields   *FieldSet `url:"fields,omitempty"`
}

// ModifyContact contains parameters that will be passed to ModifyContact API.
type ModifyContact struct {
	Metadata *MetadataSet `url:"metadata,omitempty"`
}
