package types

import "encoding/json"

// Groups represents a list of Wrike group.
type Groups struct {
	Kind string  `json:"kind"`
	Data []Group `json:"data"`
}

// Group represents a Wrike group.
type Group struct {
	ID        ContactID   `json:"id"`
	AccountID AccountID   `json:"accountId"`
	Title     string      `json:"title"`
	MemberIDs []ContactID `json:"memberIds"`
	ChildIDs  []ContactID `json:"childIds"`
	ParentIDs []ContactID `json:"parentIds"`
	AvatarURL string      `json:"avatarUrl"`
	MyTeam    bool        `json:"myTeam"`
	Metadata  []Metadata  `json:"metadata"`
}

// NewGroupsFromJSON parses the given JSON (as byte sequence) and returns a new Groups.
func NewGroupsFromJSON(data []byte) (*Groups, error) {
	var groups Groups
	err := json.Unmarshal(data, &groups)
	return &groups, err
}
