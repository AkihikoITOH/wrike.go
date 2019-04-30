package types

import "encoding/json"

const (
	Sun WorkDay = "Sun"
	Mon WorkDay = "Mon"
	Tue WorkDay = "Tue"
	Wed WorkDay = "Wed"
	Thu WorkDay = "Thu"
	Fri WorkDay = "Fri"
	Sat WorkDay = "Sat"
)

// Accounts represents a list of Wrike account.
type Accounts struct {
	Kind string    `json:"kind"`
	Data []Account `json:"data"`
}

type WorkDays []WorkDay
type WorkDay string

// CustomFieldSet represents a list of custom field objects.
type CustomFieldSet []CustomField

// Subscription represents a subscription object, a part of account.
type Subscription struct {
	Type      string `json:"type"`
	Paid      bool   `json:"paid"`
	UserLimit int    `json:"userLimit"`
}

type AccountID string

// Account represents a Wrike account.
type Account struct {
	ID             AccountID      `json:"id"`
	Name           string         `json:"name"`
	DateFormat     string         `json:"dateFormat"`
	FirstDayOfWeek string         `json:"firstDayOfWeek"`
	WorkDays       WorkDays       `json:"workDays"`
	RootFolderID   FolderID       `json:"rootFolderId"`
	RecycleBinID   FolderID       `json:"recycleBinID"`
	CreatedDate    string         `json:"createdDate"`
	Subscription   Subscription   `json:"subscription"`
	Metadata       []Metadata     `json:"metadata"`
	CustomFields   CustomFieldSet `json:"customFields"`
	JoinedDate     string         `json:"joinedDate"`
}

// NewAccountsFromJSON parses the given JSON (as byte sequence) and returns a new Accounts.
func NewAccountsFromJSON(data []byte) (*Accounts, error) {
	var accounts Accounts
	err := json.Unmarshal(data, &accounts)
	return &accounts, err
}
