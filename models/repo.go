package models

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/senghoo/captain/modules/git"
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
func (r *Repository) StatusString() string {
	switch r.Status {
	case REPOSITORY_STATUS_IDLE:
		return "idle"
	case REPOSITORY_STATUS_CLONEING:
		return "clone"
	default:
		return "unknown"
	}
}

func (r *Repository) Update() {
	if r.Exists() {
		r.Pull()
	} else {
		r.clone()
	}
}

func (r *Repository) Pull() (string, error) {
	p, err := r.Path()
	if err != nil {
		return "", err
	}
	return git.Pull(p)
}

func (r *Repository) Clone() {
	r.clone()
}

func (r *Repository) Exists() bool {
	p, err := r.Path()
	if err != nil {
		return false
	}

	_, err = os.Stat(p)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (r *Repository) clone() error {
	p, err := r.Path()
	if err != nil {
		return err
	}
	git.Clone(r.CloneURL, p)
	return nil
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
	return path.Join(ws.WorkDir(), "repos", r.FullName), nil
}
