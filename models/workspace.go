package models

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/senghoo/captain/modules/settings"
)

type Workspace struct {
	ID      int64
	Name    string `xorm:"not null unique"`
	BuildNo int64
	Created time.Time `xorm:"CREATED"`
	Updated time.Time `xorm:"UPDATED"`
	Deleted time.Time `xorm:"deleted"`
}

func NewWorkspace(name string) *Workspace {
	return &Workspace{
		Name: name,
	}
}

func Workspaces() ([]*Workspace, error) {
	var workspaces []*Workspace
	return workspaces, x.Asc("id").Find(&workspaces)
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

func (w *Workspace) Repositories() ([]*Repository, error) {
	var repos []*Repository
	err := x.Find(&repos, &Repository{WorkspaceID: w.ID})
	return repos, err
}

func (w *Workspace) AddRepository(repo *Repository) {
	repo.WorkspaceID = w.ID
	x.Insert(repo)
}

func (w *Workspace) Builds() ([]*Build, error) {
	var builds []*Build
	err := x.Find(&builds, &Build{WorkspaceID: w.ID})
	return builds, err
}

type Build struct {
	ID          int64
	WorkspaceID int64
	workspace   *Workspace `xorm:"-"`
	BuildNo     int64
	Type        string
	Name        string      `xorm:"not null unique"`
	Created     time.Time   `xorm:"CREATED"`
	Updated     time.Time   `xorm:"UPDATED"`
	Deleted     time.Time   `xorm:"deleted"`
	logger      *log.Logger `xorm:"-"`
	Context     map[string]interface{}
}

func (w *Workspace) NewBuild(t string) (*Build, error) {
	buildNo := w.BuildNo
	w.BuildNo++
	w.Save()
	build := &Build{
		WorkspaceID: w.ID,
		BuildNo:     buildNo,
		Type:        t,
		Name:        fmt.Sprintf("%s:%s:%d", w.Name, t, buildNo),
	}
	_, err := x.Insert(build)
	if err != nil {
		return nil, err
	}

	return build, err
}

func (b *Build) Logger() *log.Logger {
	if b.logger != nil {
		return b.logger
	}

	f, _ := os.Create(b.LogFile())

	b.logger = log.New(f, fmt.Sprintf("build[%d]", b.BuildNo), log.Lshortfile)
	return b.logger
}

func (b *Build) LogFile() string {
	return path.Join(b.Path(), "log.log")
}

func (b *Build) Workspace() *Workspace {
	if b.workspace != nil {
		return b.workspace
	}

	if b.WorkspaceID != 0 {
		ws := new(Workspace)
		has, _ := GetByID(b.WorkspaceID, ws)
		if has {
			return ws
		}
	}
	return nil
}

func (b *Build) Path() string {
	ws := b.Workspace()
	p := path.Join(ws.WorkDir(), "builds", fmt.Sprintf("%d", b.BuildNo))
	os.MkdirAll(p, 0700)
	return p
}
