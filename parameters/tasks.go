package parameters

import (
	"net/url"

	types "github.com/AkihikoITOH/wrike.go/types"
)

const (
	RecurrentField        TaskField = "recurrent"
	AttachmentCountField  TaskField = "attachmentCount"
	EffortAllocationField TaskField = "effortAllocation"
	AuthorIDsField        TaskField = "authorIds"
	HasAttachmentsField   TaskField = "hasAttachments"
	ParentIDsField        TaskField = "parentIds"
	SuperParentIDsField   TaskField = "superParentIds"
	SharedIDsField        TaskField = "sharedIds"
	ResponsibleIDsField   TaskField = "responsibleIds"
	DescriptionField      TaskField = "description"
	BriefDescriptionField TaskField = "briefDescription"
	SuperTaskIDsField     TaskField = "superTaskIds"
	SubTaskIDsField       TaskField = "subTaskIds"
	DependencyIDsField    TaskField = "dependencyIds"
	MetadataField         TaskField = "metadata"
	CustomFieldsField     TaskField = "customFields"

	CreatedDateSortField    TaskSortField = "CreatedDate"
	UpdatedDateSortField    TaskSortField = "UpdatedDate"
	CompletedDateSortField  TaskSortField = "CompletedDate"
	DueDateSortField        TaskSortField = "DueDate"
	StatusSortField         TaskSortField = "Status"
	ImportanceSortField     TaskSortField = "Importance"
	TitleSortField          TaskSortField = "Title"
	LastAccessDateSortField TaskSortField = "LastAccessDate"

	Asc  TaskSortOrder = "Asc"
	Desc TaskSortOrder = "Desc"
)

type TaskSortField string
type TaskSortOrder string

type TaskField string
type TaskFieldSet []TaskField

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f TaskFieldSet) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

type TaskDates struct {
	Type           types.TaskDatesType `json:"type,omitempty"`
	Duration       int                 `json:"duration,omitempty"`
	Start          string              `json:"start,omitempty"`
	Due            string              `json:"due,omitempty"`
	WorkOnWeekends bool                `json:"workOnWeekends,omitempty,omitempty"`
}

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f TaskDates) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

type TaskStatusSet []types.TaskStatus

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f TaskStatusSet) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

type TaskIDSet []types.TaskID

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f TaskIDSet) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

type DateOrRange struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
	Equal string `json:"equal,omitempty"`
}

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f *DateOrRange) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

type CustomStatusIDSet []types.CustomStatusID

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f CustomStatusIDSet) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

// QueryTasksByIDs contains parameters that will be passed to QueryTasksByIDs API.
type QueryTasksByIDs struct {
	Fields TaskFieldSet `url:"fields,omitempty"`
}

type TaskEffort struct {
	Mode            types.TaskEffortMode `json:"mode"`
	TotalEffort     int                  `json:"totalEffort,omitempty"`
	AllocatedEffort int                  `json:"allocatedEffort,omitempty"`
}

// EncodeValues defines how to encode the type. For more details, refer to https://godoc.org/github.com/google/go-querystring/query#Encoder
func (f TaskEffort) EncodeValues(key string, v *url.Values) error {
	return jsonAsURLParams(f, key, v)
}

// QueryTasks contains parameters that will be passed to QueryTasks API.
type QueryTasks struct {
	Descendants    *bool                `url:"descendants,omitempty"`
	Title          *string              `url:"title,omitempty"`
	Status         TaskStatusSet        `url:"status,omitempty"`
	Importance     types.TaskImportance `url:"importance,omitempty"`
	StartDate      *DateOrRange         `url:"startDate,omitempty"`
	DueDate        *DateOrRange         `url:"dueDate,omitempty"`
	ScheduledDate  *DateOrRange         `url:"scheduledDate,omitempty"`
	CreatedDate    *DateOrRange         `url:"createdDate,omitempty"`
	UpdatedDate    *DateOrRange         `url:"updatedDate,omitempty"`
	CompletedDate  *DateOrRange         `url:"completedDate,omitempty"`
	Authors        ContactIDSet         `url:"authors,omitempty"`
	Responsibles   ContactIDSet         `url:"responsibles,omitempty"`
	Permalink      *string              `url:"permalink,omitempty"`
	Type           types.TaskDatesType  `url:"type,omitempty"`
	Limit          *int                 `url:"limit,omitempty"`
	SortField      TaskSortField        `url:"sortField,omitempty"`
	SortOrder      TaskSortOrder        `url:"sortOrder,omitempty"`
	SubTasks       *bool                `url:"subTasks,omitempty"`
	PageSize       *int                 `url:"pageSize,omitempty"`
	NextPageToken  *string              `url:"nextPageToken,omitempty"`
	Metadata       *Metadata            `url:"metadata,omitempty"`
	CustomField    *CustomField         `url:"customField,omitempty"`
	CustomStatuses CustomStatusIDSet    `url:"customStatuses,omitempty"`
	Fields         TaskFieldSet         `url:"fields,omitempty"`
}

// CreateTask contains parameters that will be passed to CreateTask API.
type CreateTask struct {
	Title            string               `url:"title,omitempty"`
	Description      *string              `url:"description,omitempty"`
	Status           TaskStatusSet        `url:"status,omitempty"`
	Importance       types.TaskImportance `url:"importance,omitempty"`
	Dates            *TaskDates           `url:"dates,omitempty"`
	Shareds          ContactIDSet         `url:"shareds,omitempty"`
	Parents          FolderIDSet          `url:"parents,omitempty"`
	Responsibles     ContactIDSet         `url:"responsibles,omitempty"`
	Followers        ContactIDSet         `url:"followers,omitempty"`
	Follow           *bool                `url:"follow,omitempty"`
	PriorityBefore   types.TaskID         `url:"priorityBefore,omitempty"`
	PriorityAfter    types.TaskID         `url:"priorityAfter,omitempty"`
	SuperTasks       TaskIDSet            `url:"superTasks,omitempty"`
	Metadata         *MetadataSet         `url:"metadata,omitempty"`
	CustomFields     CustomFieldSet       `url:"customFields,omitempty"`
	CustomStatus     types.CustomStatusID `url:"customStatus,omitempty"`
	EffortAllocation *TaskEffort          `url:"effortAllocation,omitempty"`
	Fields           TaskFieldSet         `url:"fields,omitempty"`
}

// ModifyTask contains parameters that will be passed to ModifyTask API.
type ModifyTask struct {
	Title              *string              `url:"title,omitempty"`
	Description        *string              `url:"description,omitempty"`
	Status             TaskStatusSet        `url:"status,omitempty"`
	Importance         types.TaskImportance `url:"importance,omitempty"`
	Dates              *TaskDates           `url:"dates,omitempty"`
	AddParents         FolderIDSet          `url:"addParents,omitempty"`
	RemoveParents      FolderIDSet          `url:"removeParents,omitempty"`
	AddShareds         ContactIDSet         `url:"addShareds,omitempty"`
	RemoveShareds      ContactIDSet         `url:"removeShareds,omitempty"`
	AddResponsibles    ContactIDSet         `url:"addResponsibles,omitempty"`
	RemoveResponsibles ContactIDSet         `url:"removeResponsibles,omitempty"`
	AddFollowers       ContactIDSet         `url:"addFollowers,omitempty"`
	Follow             *bool                `url:"follow,omitempty"`
	PriorityBefore     types.TaskID         `url:"priorityBefore,omitempty"`
	PriorityAfter      types.TaskID         `url:"priorityAfter,omitempty"`
	AddSuperTasks      TaskIDSet            `url:"addSuperTasks,omitempty"`
	RemoveSuperTasks   TaskIDSet            `url:"removeSuperTasks,omitempty"`
	Metadata           *MetadataSet         `url:"metadata,omitempty"`
	CustomFields       CustomFieldSet       `url:"customFields,omitempty"`
	CustomStatus       types.CustomStatusID `url:"customStatus,omitempty"`
	Restore            *bool                `url:"restore,omitempty"`
	EffortAllocation   *TaskEffort          `url:"effortAllocation,omitempty"`
	Fields             TaskFieldSet         `url:"fields,omitempty"`
}

// ModifyTasks contains parameters that will be passed to ModifyTasks API.
type ModifyTasks struct {
	CustomFields     CustomFieldSet `url:"customFields,omitempty"`
	EffortAllocation *TaskEffort    `url:"effortAllocation,omitempty"`
}
