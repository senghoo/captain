package models

import "time"

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
