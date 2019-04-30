package types

import "encoding/json"

const (
	ContactTypePerson = "Person"
	ContactTypeGroup  = "Group"
)

// Contacts represents a list of Wrike contacts.
type Contacts struct {
	Kind string    `json:"kind"`
	Data []Contact `json:"data"`
}

// ContactID is a string that represents a Wrike Contact ID
type ContactID string

// Contact represents a Wrike contact.
type Contact struct {
	ID        ContactID   `json:"id"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Type      string      `json:"type"`
	Profiles  []Profile   `json:"profiles"`
	AvatarURL string      `json:"avatarUrl"`
	Timezone  string      `json:"timezone"`
	Locale    string      `json:"locale"`
	Deleted   bool        `json:"deleted"`
	Me        bool        `json:"me"`
	MemberIDs []ContactID `json:"memberIds"`
	Metadata  []Metadata  `json:"metadata"`
	Phone     string      `json:"phone"`
}

// IsPerson returns true if and only if the contact is a person.
func (c *Contact) IsPerson() bool {
	return c.Type == ContactTypePerson
}

// IsGroup returns true if and only if the contact is a group.
func (c *Contact) IsGroup() bool {
	return c.Type == ContactTypeGroup
}

// Profile represents a profile, part of Contact object.
type Profile struct {
	AccountID string `json:"accountId"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	External  bool   `json:"external"`
	Admin     bool   `json:"admin"`
	Owner     bool   `json:"owner"`
}

// NewContactsFromJSON parses the given JSON (as byte sequence) and returns a new Contacts.
func NewContactsFromJSON(data []byte) (*Contacts, error) {
	var contacts Contacts
	err := json.Unmarshal(data, &contacts)
	return &contacts, err
}
