package models

import (
	"fmt"
	"strings"
	"time"
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
