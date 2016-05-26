package models

import (
	"errors"
	"fmt"
	"path"
	"strings"
	"time"
)

const (
	REPOSITORY_STATUS_IDLE = iota
	REPOSITORY_STATUS_CLONEING
)

type Repository struct {
	ID            int64
	Name          string
	Site          string
	WorkspaceID   int64
	workspace     *Workspace `xorm:"-"`
	FullName      string     `xorm:"not null unique"`
	Description   string
	Homepage      string
	DefaultBranch string
	MasterBranch  string
	CreatedAt     time.Time
	PushedAt      time.Time
	UpdatedAt     time.Time
	HTMLURL       string
	CloneURL      string
	GitURL        string
	SSHURL        string
	Language      string
	Status        int
	Created       time.Time `xorm:"CREATED"`
	Updated       time.Time `xorm:"UPDATED"`
	Deleted       time.Time `xorm:"deleted"`
}

func GetRepositoryByIdentify(identify string) (*Repository, error) {
	s := strings.Split(identify, ":")
	if len(s) != 2 {
		return nil, fmt.Errorf("unknown identify %s", identify)
	}

	switch s[0] {
	case "github":
		acc, err := GetGithubAccount()
		if err != nil {
			return nil, err
		}
		return acc.GetRepoByFullName(s[1])
	}
	return nil, fmt.Errorf("unknown identify %s", identify)
}

func (r *Repository) Identify() string {
	return fmt.Sprintf("%s:%s", r.Site, r.FullName)
}

func (r *Repository) Clone() {
	if r.Status == REPOSITORY_STATUS_IDLE {
		go r.clone()
	}
}

func (r *Repository) clone() {
	r.Status = REPOSITORY_STATUS_CLONEING
	x.Id(r.ID).Update(r)
	defer func() {
		r.Status = REPOSITORY_STATUS_IDLE
		x.Id(r.ID).Update(r)
	}()

}

func (r *Repository) Workspace() *Workspace {
	if r.workspace != nil {
		return r.workspace
	}

	if r.WorkspaceID != 0 {
		ws := new(Workspace)
		has, _ := GetByID(r.WorkspaceID, ws)
		if has {
			return ws
		}
	}
	return nil
}

func (r *Repository) Path() (string, error) {
	ws := r.Workspace()
	if ws == nil {
		return "", errors.New("Workspace not exists")
	}
	return path.Join(ws.WorkDir(), "repos", r.FullName)
}
