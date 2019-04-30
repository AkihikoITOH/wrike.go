package types

import "encoding/json"

const (
	ActiveTaskStatus    TaskStatus = "Active"
	CompletedTaskStatus TaskStatus = "Completed"
	DeferredTaskStatus  TaskStatus = "Deferred"
	CancelledTaskStatus TaskStatus = "Cancelled"

	HighImportance   TaskImportance = "High"
	NormalImportance TaskImportance = "Normal"
	LowImportance    TaskImportance = "Low"

	BacklogType   TaskDatesType = "Backlog"
	MilestoneType TaskDatesType = "Milestone"
	PlannedType   TaskDatesType = "Planned"

	NoneEffortMode     TaskEffortMode = "None"
	FullTimeEffortMode TaskEffortMode = "FullTime"
	BasicEffortMode    TaskEffortMode = "Basic"
	FlexibleEffortMode TaskEffortMode = "Flexible"
)

// Tasks represents a list of Wrike task.
type Tasks struct {
	Kind string `json:"kind"`
	Data []Task `json:"data"`
}

type DependencyID string
type TaskID string
type TaskStatus string
type TaskImportance string
type TaskDatesType string
type TaskEffortMode string

type TaskDates struct {
	Type           TaskDatesType `json:"type"`
	Duration       int           `json:"duration"`
	Start          string        `json:"start"`
	Due            string        `json:"due"`
	WorkOnWeekends bool          `json:"workOnWeekends,omitempty"`
}

type TaskEffort struct {
	Mode            TaskEffortMode `json:"mode"`
	TotalEffort     int            `json:"totalEffort,omitempty"`
	AllocatedEffort int            `json:"allocatedEffort,omitempty"`
}

type Task struct {
	ID               TaskID         `json:"id"`
	AccountID        AccountID      `json:"accountId"`
	Title            string         `json:"title"`
	Description      string         `json:"description,omitempty"`
	BriefDescription string         `json:"briefDescription,omitempty"`
	ParentIDs        []FolderID     `json:"parentIds,omitempty"`
	SuperParentIDs   []FolderID     `json:"superParentIds,omitempty"`
	SharedIDs        []ContactID    `json:"sharedIds,omitempty"`
	ResponsibleIDs   []ContactID    `json:"responsibleIds,omitempty"`
	Status           TaskStatus     `json:"status"`
	Importance       TaskImportance `json:"importance"`
	CreatedDate      string         `json:"createdDate"`
	UpdatedDate      string         `json:"updatedDate"`
	CompletedDate    string         `json:"completedDate"`
	Dates            TaskDates      `json:"dates"`
	Scope            TreeScope      `json:"scope"`
	AuthorIDs        []ContactID    `json:"authorIds,omitempty"`
	CustomStatusID   CustomStatusID `json:"customStatusId"`
	HasAttachments   bool           `json:"hasAttachments,omitempty"`
	AttachmentCount  int            `json:"attachmentCount,omitempty"`
	Permalink        string         `json:"permalink"`
	Priority         string         `json:"priority"`
	FollowedByMe     bool           `json:"followedByMe,omitempty"`
	FollwerIDs       []ContactID    `json:"followerIds,omitempty"`
	Recurrent        bool           `json:"recurrent,omitempty"`
	SuperTaskIDs     []TaskID       `json:"superTaskIds,omitempty"`
	SubTaskIDs       []TaskID       `json:"subTaskIds,omitempty"`
	DependencyIDs    []DependencyID `json:"dependencyIds,omitempty"`
	Metadata         []Metadata     `json:"metadata,omitempty"`
	CustomFields     []CustomField  `json:"customFields,omitempty"`
	EffortAllocation TaskEffort     `json:"effortAllocation,omitempty"`
}

// NewTasksFromJSON parses the given JSON (as byte sequence) and returns a new Tasks.
func NewTasksFromJSON(data []byte) (*Tasks, error) {
	var tasks Tasks
	err := json.Unmarshal(data, &tasks)
	return &tasks, err
}
