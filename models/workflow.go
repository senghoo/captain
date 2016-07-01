package models

import (
	"fmt"

	"github.com/senghoo/captain/commands"
	"github.com/senghoo/captain/models"
)

type Workflow struct {
	ID          int64
	WorkspaceID int64
	Name        string `xorm:"not null unique"`
	Config      string
	node        *command.Node `xorm:"-"`
}

func NewWorkflow(workspaceID int64, name, config string) *Workflow {
	return &Workflow{
		WorkspaceID: workspaceID,
		Name:        name,
		Config:      config,
	}
}

func (w *Workflow) Node() (*Node, error) {
	if w.node != nil {
		return w.node, nil
	}
	node, err := command.ParseNode([]byte(w.Config))
	if err == nil {
		w.node = node
	}
	return node, err
}

func (w *Workflow) Workspace() (*Workspace, error) {
	ws := new(Workspace)
	has, err := models.GetByID(w.WorkspaceID, ws)
	if !has {
		return nil, fmt.Errorf("workspace %d not found", w.WorkspaceID)
	}
	return ws, err
}

func (w *Workflow) Execute() error {
	node, err := w.Node()
	if err != nil {
		return errr
	}

	ws, err := w.Workspace()
	if err != nil {
		return errr
	}

	build, err := ws.NewBuild(w.Name)
	if err != nil {
		return nil
	}

	command.RunNode(node, build)
}
