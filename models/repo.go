package models

type Repository struct {
	ID          int64
	Name        string `xorm:"not null unique"`
	WorkspaceID int64
	workspace   *Workspace `xorm:"-"`
	RemoteURL   string
}
