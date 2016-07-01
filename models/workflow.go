package models

import "fmt"

type Workflow struct {
	ID          int64
	WorkspaceID int64
	Name        string `xorm:"not null unique"`
	Config      string
}

func NewWorkflow(workspaceID int64, name, config string) *Workflow {
	return &Workflow{
		WorkspaceID: workspaceID,
		Name:        name,
		Config:      config,
	}
}

func (w *Workflow) Workspace() (*Workspace, error) {
	ws := new(Workspace)
	has, err := GetByID(w.WorkspaceID, ws)
	if !has {
		return nil, fmt.Errorf("workspace %d not found", w.WorkspaceID)
	}
	return ws, err
}
