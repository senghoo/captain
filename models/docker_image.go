package models

import "time"

type Registry struct {
	ID       int64
	URL      string
	Username string
	Password string
	Created  time.Time `xorm:"CREATED"`
	Updated  time.Time `xorm:"UPDATED"`
	Deleted  time.Time `xorm:"deleted"`
}

type DockerImage struct {
	ID           int64
	Name         string
	RegistryID   int64
	RepositoryID int64
	Created      time.Time `xorm:"CREATED"`
	Updated      time.Time `xorm:"UPDATED"`
	Deleted      time.Time `xorm:"deleted"`
}

type DockerImageVersion struct {
	ID      int64
	ImageID int64
	BuildID int64
	Version string
}
