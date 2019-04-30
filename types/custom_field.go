package types

import "encoding/json"

const (
	TextType       CustomFieldType = "Text"
	DropDownType   CustomFieldType = "DropDown"
	NumericType    CustomFieldType = "Numeric"
	CurrencyType   CustomFieldType = "Currency"
	PercentageType CustomFieldType = "Percentage"
	DateType       CustomFieldType = "Date"
	DurationType   CustomFieldType = "Duration"
	CheckBoxType   CustomFieldType = "CheckBox"
	ContactsType   CustomFieldType = "Contacts"
	MultipleType   CustomFieldType = "Multiple"
)

// CustomFields represents a list of Wrike custom fields.
type CustomFields struct {
	Kind string        `json:"kind"`
	Data []CustomField `json:"data"`
}

// Settings represents a settings object, a part of account.
type Settings struct {
	InheritanceType       string `json:"inheritanceType"`
	DecimalPlaces         int    `json:"decimalPlaces"`
	UseThousandsSeparator bool   `json:"useThousandsSeparator"`
	Aggregation           string `json:"aggregation"`
}

type CustomFieldID string
type CustomFieldType string

// CustomField represents a custom field object, a part of account.
type CustomField struct {
	ID        CustomFieldID   `json:"id"`
	AccountID AccountID       `json:"accountId"`
	Title     string          `json:"title"`
	Type      CustomFieldType `json:"type"`
	SharedIDs []ContactID     `json:"sharedIds"`
	Settings  Settings        `json:"settings"`
}

// NewCustomFieldsFromJSON parses the given JSON (as byte sequence) and returns a new CustomFields.
func NewCustomFieldsFromJSON(data []byte) (*CustomFields, error) {
	var customFields CustomFields
	err := json.Unmarshal(data, &customFields)
	return &customFields, err
}
