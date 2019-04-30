package parameters

import "github.com/AkihikoITOH/wrike.go/types"

type Settings struct {
	InheritanceType       *string      `url:"inheritanceType,omitempty"`
	DecimalPlaces         *int         `url:"decimalPlaces,omitempty"`
	UseThousandsSeparator *bool        `url:"useThousandsSeparator,omitempty"`
	Currency              *string      `url:"currency,omitempty"`
	Aggregation           *string      `url:"aggregation,omitempty"`
	Values                *[]string    `url:"values,omitempty"`
	AllowOtherValues      *bool        `url:"allowOtherValues,omitempty"`
	Contacts              ContactIDSet `url:"contacts,omitempty"`
}

// CreateCustomField contains parameters that will be passed to CreateCustomField API.
type CreateCustomField struct {
	Title    string                `url:"title"`
	Type     types.CustomFieldType `url:"type"`
	Shareds  ContactIDSet          `url:"shareds,omitempty"`
	Settings *Settings             `url:"settings,omitempty"`
}

// ModifyCustomField contains parameters that will be passed to ModifyCustomField API.
type ModifyCustomField struct {
	Title         *string                `url:"title"`
	Type          *types.CustomFieldType `url:"type"`
	AddShareds    ContactIDSet           `url:"addShareds,omitempty"`
	RemoveShareds ContactIDSet           `url:"removeShareds,omitempty"`
	Settings      *Settings              `url:"settings,omitempty"`
}
