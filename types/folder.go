package types

import "encoding/json"

const (
	WsRoot   TreeScope = "WsRoot"
	RbRoot   TreeScope = "RbRoot"
	WsFolder TreeScope = "WsFolder"
	RbFolder TreeScope = "RbFolder"
	WsTask   TreeScope = "WsTask"
	RbTask   TreeScope = "RbTask"

	Green     ProjectStatus = "Green"
	Yellow    ProjectStatus = "Yellow"
	Red       ProjectStatus = "Red"
	Completed ProjectStatus = "Completed"
	OnHold    ProjectStatus = "OnHold"
	Cancelled ProjectStatus = "Cancelled"
)

type FolderTree struct {
	Kind string        `json:"kind"`
	Data []FolderEntry `json:"data"`
}

type Folders struct {
	Kind string   `json:"kind"`
	Data []Folder `json:"data"`
}

type TreeScope string
type ProjectStatus string

type Project struct {
	AuthorID      ContactID     `json:"authorId,omitempty"`
	OwnerIDs      []ContactID   `json:"ownerIds"`
	Status        ProjectStatus `json:"status"`
	StartDate     string        `json:"startDate,omitempty"`
	EndDate       string        `json:"endDate,omitempty"`
	CreatedDate   string        `json:"createdDate,omitempty"`
	CompletedDate string        `json:"completedDate,omitempty"`
}

type FolderEntry struct {
	ID       FolderID   `json:"id"`
	Title    string     `json:"title"`
	Color    string     `json:"color,omitempty"`
	ChildIDs []FolderID `json:"childIds,omitempty"`
	Scope    TreeScope  `json:"scope,omitempty"`
	Project  Project    `json:"project,omitempty"`
}

type Folder struct {
	ID               FolderID         `json:"id"`
	AccountID        AccountID        `json:"accountId"`
	Title            string           `json:"title"`
	CreatedDate      string           `json:"createdDate"`
	UpdatedDate      string           `json:"updatedDate"`
	BriefDescription *string          `json:"briefDescription,omitempty"`
	Description      string           `json:"description"`
	Color            *string          `json:"color,omitempty"`
	SharedIDs        []ContactID      `json:"sharedIds"`
	ParentIDs        []FolderID       `json:"parentIds"`
	ChildIDs         []FolderID       `json:"childIds"`
	SuperParentIDs   []FolderID       `json:"superParentIds"`
	Scope            TreeScope        `json:"scope"`
	HasAttachments   bool             `json:"hasAttachments"`
	AttachmentCount  *int             `json:"attachmentCount,omitempty"`
	Permalink        string           `json:"permalink"`
	WorkflowID       WorkflowID       `json:"workflowId"`
	Metadata         *[]Metadata      `json:"metadata,omitempty"`
	CustomFields     *[]CustomField   `json:"customFields,omitempty"`
	CustomColumnIDs  *[]CustomFieldID `json:"customColumnIds"`
	Project          *Project         `json:"project"`
}

// NewFolderTreeFromJSON parses the given JSON (as byte sequence) and returns a new Groups.
func NewFolderTreeFromJSON(data []byte) (*FolderTree, error) {
	var folderTree FolderTree
	err := json.Unmarshal(data, &folderTree)
	return &folderTree, err
}

// NewFoldersFromJSON parses the given JSON (as byte sequence) and returns a new Groups.
func NewFoldersFromJSON(data []byte) (*Folders, error) {
	var folders Folders
	err := json.Unmarshal(data, &folders)
	return &folders, err
}
