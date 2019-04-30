// Package types contains various types that are used to parse responses from the API.
package types

const (
	SubscriptionField Field = "subscription"
	MetadataField     Field = "metadata"
	CustomFieldsField Field = "customFields"
)

// Field represents a Wrike field.
type Field string

type FolderID string

// Metadata represents a metadata, part of other API objects.
type Metadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
