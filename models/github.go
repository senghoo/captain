package models

import (
	"time"

	gc "github.com/google/go-github/github"
	"github.com/senghoo/captain/modules/github"
)

type GithubAccount struct {
	ID          int64
	OwnerID     int64
	Name        string
	AccessToken string
	Created     time.Time  `xorm:"CREATED"`
	Updated     time.Time  `xorm:"UPDATED"`
	client      *gc.Client `xorm:"-"`
}

func NewGithubAccount(oid int64, token string) (a *GithubAccount) {
	a = &GithubAccount{
		OwnerID:     oid,
		AccessToken: token,
	}
	a.UpdateName()
	return
}

func (a *GithubAccount) Client() *gc.Client {
	if a.client == nil {
		a.client = github.GithubClient(a.AccessToken)
	}
	return a.client
}

func (a *GithubAccount) GithubUser() (user *gc.User, err error) {
	user, _, err = a.Client().Users.Get("")
	return
}

func (a *GithubAccount) UpdateName() {
	user, err := a.GithubUser()
	if err != nil {
		return
	}
	a.Name = *user.Name
}

func (a *GithubAccount) Save() {
	if a.ID == 0 {
		x.Insert(a)
	} else {
		x.Id(a.ID).Update(a)
	}
}
