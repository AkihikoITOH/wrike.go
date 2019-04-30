package types

import (
	"encoding/json"
)

// Workflows represents a list of Wrike workflow.
type Workflows struct {
	Kind string     `json:"kind"`
	Data []Workflow `json:"data"`
}

type WorkflowID string

// Workflow represents a Wrike workflow.
type Workflow struct {
	ID             WorkflowID     `json:"id"`
	Name           string         `json:"name"`
	Standard       bool           `json:"standard"`
	Hidden         bool           `json:"hidden"`
	CustomStatuses CustomStatuses `json:"customStatuses"`
}

type CustomStatusID string
type CustomStatuses []CustomStatus
type CustomStatus struct {
	ID           CustomStatusID `json:"id"`
	Name         string         `json:"name"`
	StandardName bool           `json:"standardName"`
	Color        string         `json:"color"`
	Standard     bool           `json:"standard"`
	Group        string         `json:"group"`
	Hidden       bool           `json:"hidden"`
}

// NewWorkflowsFromJSON parses the given JSON (as byte sequence) and returns a new Workflows.
func NewWorkflowsFromJSON(data []byte) (*Workflows, error) {
	var workflows Workflows
	err := json.Unmarshal(data, &workflows)
	return &workflows, err
}
