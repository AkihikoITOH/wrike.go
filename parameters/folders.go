package parameters

import (
	"net/url"

	"github.com/AkihikoITOH/wrike.go/types"
)

const (
	EqualTo          Comparator = "EqualTo"
	IsEmpty          Comparator = "IsEmpty"
	IsNotEmpty       Comparator = "IsNotEmpty"
	LessThan         Comparator = "LessThan"
	LessOrEqualTo    Comparator = "LessOrEqualTo"
	GreaterThan      Comparator = "GreaterThan"
	GreaterOrEqualTo Comparator = "GreaterOrEqualTo"
	InRange          Comparator = "InRange"
	NotInRange       Comparator = "NotInRange"
	Contains         Comparator = "Contains"
	StartsWith       Comparator = "StartsWith"
	EndsWith         Comparator = "EndsWith"
	ContainsAll      Comparator = "ContainsAll"
	ContainsAny      Comparator = "ContainsAny"
)

type DateRange struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type Comparator string

type CustomFieldFilter struct {
	ID         string     `json:"id"`
	Comparator Comparator `json:"comparator,omitempty"`
	Value      *string    `json:"value,omitempty"`
	MinValue   *string    `json:"minValue,omitempty"`
	MaxValue   *string    `json:"maxValue,omitempty"`
	Values     []string   `json:"values,omitempty"`
}

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f *CustomFieldFilter) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

// GetFolderTree contains parameters that will be passed to GetFolderTree API.
type GetFolderTree struct {
	Permalink   *string            `url:"permalink,omitempty"`
	Descendants *bool              `url:"descendants,omitempty"`
	Metadata    *Metadata          `url:"metadata,omitempty"`
	CustomField *CustomFieldFilter `url:"customField,omitempty"`
	UpdatedDate *DateRange         `url:"updatedDate,omitempty"`
	Project     *bool              `url:"project,omitempty"`
	Deleted     *bool              `url:"deleted,omitempty"`
	Fields      *FieldSet          `url:"fields,omitempty"`
}

// GetFolderSubtree contains parameters that will be passed to GetFolderSubTree API.
type GetFolderSubtree struct {
	Permalink   *string            `url:"permalink,omitempty"`
	Descendants *bool              `url:"descendants,omitempty"`
	Metadata    *Metadata          `url:"metadata,omitempty"`
	CustomField *CustomFieldFilter `url:"customField,omitempty"`
	UpdatedDate *DateRange         `url:"updatedDate,omitempty"`
	Project     *bool              `url:"project,omitempty"`
	Fields      *FieldSet          `url:"fields,omitempty"`
}

type CustomFieldIDSet []types.CustomFieldID

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f CustomFieldIDSet) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

type CustomField struct {
	ID    types.CustomFieldID `json:"id"`
	Value string              `json:"value,omitempty"`
}

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f *CustomField) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

type CustomFieldSet []CustomField

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f CustomFieldSet) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

type Project struct {
	OwnerIDs  ContactIDSet        `json:"ownerIds,omitempty"`
	Status    types.ProjectStatus `json:"status,omitempty"`
	StartDate string              `json:"startDate,omitempty"`
	EndDate   string              `json:"endDate,omitempty"`
}

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (p Project) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(p, key, v)
}

type ProjectModify struct {
	OwnersAdd    ContactIDSet        `json:"ownerAdd,omitempty"`
	OwnersRemove ContactIDSet        `json:"ownerRemove,omitempty"`
	Status       types.ProjectStatus `json:"status,omitempty"`
	StartDate    string              `json:"startDate,omitempty"`
	EndDate      string              `json:"endDate,omitempty"`
}

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (p ProjectModify) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(p, key, v)
}

type CreateFolder struct {
	Title         string           `url:"title"`
	Description   string           `url:"description,omitempty"`
	Shareds       ContactIDSet     `url:"shareds,omitempty"`
	Metadata      *MetadataSet     `url:"metadata,omitempty"`
	CustomFields  CustomFieldSet   `url:"customFields,omitempty"`
	CustomColumns CustomFieldIDSet `url:"customColumns,omitempty"`
	Project       Project          `url:"project,omitempty"`
}

// GetFolders contains parameters that will be passed to GetFolder API.
type GetFolders struct {
	Fields *FieldSet `url:"fields,omitempty"`
}

// CopyFolder contains parameters that will be passed to CopyFolder api.
type CopyFolder struct {
	Parent             types.FolderID `url:"parent"`
	Title              string         `url:"title"`
	TitlePrefix        string         `url:"titlePrefix,omitempty"`
	CopyDescriptions   *bool          `url:"copyDescriptions,omitempty"`
	CopyResponsibles   *bool          `url:"copyResponsibles,omitempty"`
	AddResponsibles    ContactIDSet   `url:"addResponsibles,omitempty"`
	RemoveResponsibles ContactIDSet   `url:"removeResponsibles,omitempty"`
	CopyCustomFields   *bool          `url:"copyCustomFields,omitempty"`
	CopyCustomStatuses *bool          `url:"copyCustomStatuses,omitempty"`
	CopyStatuses       *bool          `url:"copyStatuses,omitempty"`
	CopyParents        *bool          `url:"copyParents,omitempty"`
	RescheduleDate     string         `url:"rescheduleDate,omitempty"`
	RescheduleMode     string         `url:"rescheduleMode,omitempty"`
	EntryLimit         *int           `url:"entryLimit,omitempty"`
}

type ModifyFolder struct {
	Title         string           `url:"title"`
	Description   *string          `url:"description"`
	AddParents    FolderIDSet      `url:"addParents,omitempty"`
	RemoveParents FolderIDSet      `url:"removeParents,omitempty"`
	AddShareds    ContactIDSet     `url:"addShareds,omitempty"`
	RemoveShareds ContactIDSet     `url:"removeShareds,omitempty"`
	Metadata      *MetadataSet     `url:"metadata,omitempty"`
	Restore       *bool            `url:"restore,omitempty"`
	CustomFields  CustomFieldSet   `url:"customFields,omitempty"`
	CustomColumns CustomFieldIDSet `url:"customColumns,omitempty"`
	Project       ProjectModify    `url:"project"`
}

type ModifyFolders struct {
	CustomFields CustomFieldSet `url:"customFields,omitempty"`
}
