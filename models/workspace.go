package models

import (
	"path"
	"time"

	"github.com/senghoo/captain/modules/settings"
)

type Workspace struct {
	ID      int64
	Name    string    `xorm:"not null unique"`
	Created time.Time `xorm:"CREATED"`
	Updated time.Time `xorm:"UPDATED"`
	Deleted time.Time `xorm:"deleted"`
}

func NewWorkspace(name string) *Workspace {
	return &Workspace{
		Name: name,
	}
}

func (w *Workspace) Save() {
	if w.ID == 0 {
		x.Insert(w)
	} else {
		x.Id(w.ID).Update(w)
	}
}

func (w *Workspace) WorkDir() string {
	return path.Join(settings.GetWorkspacePath(), w.Name)
}
