package parameters

// QueryAccounts contains parameters that will be passed to QueryAccounts API.
type QueryAccounts struct {
	Metadata *Metadata `url:"metadata,omitempty"`
	Fields   *FieldSet `url:"fields,omitempty"`
}

// ModifyAccount contains parameters that will be passed to ModifyAccount API.
type ModifyAccount struct {
	Metadata *MetadataSet `url:"metadata,omitempty"`
}
