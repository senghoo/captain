package models

import (
	"fmt"
	"time"

	"github.com/senghoo/captain/modules/utils"
)

type Workflow struct {
	ID          int64
	WorkspaceID int64
	Name        string `xorm:"not null unique"`
	Config      string
	ConfigType  string
	Created     time.Time `xorm:"CREATED"`
	Updated     time.Time `xorm:"UPDATED"`
	Deleted     time.Time `xorm:"deleted"`
}

func NewWorkflow(workspaceID int64, name, config, configType string) *Workflow {
	return &Workflow{
		WorkspaceID: workspaceID,
		Name:        name,
		Config:      config,
		ConfigType:  configType,
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

func (w *Workflow) GithubWebhook() (hook *GithubWebhook, err error) {
	hook = &GithubWebhook{
		WorkflowID: w.ID,
	}
	has, err := x.Get(hook)
	if err != nil {
		return nil, err
	}

	if has {
		return hook, nil
	}

	hook.Secret = utils.RandomString(64)
	_, err = Insert(hook)
	return hook, err
}
