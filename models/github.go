package models

import (
	"time"

	gc "github.com/google/go-github/github"
	"github.com/senghoo/captain/modules/github"
)

type GithubAccount struct {
	ID          int64
	OwnerID     int64  `xorm:"not null unique(owner_name)"`
	Name        string `xorm:"not null unique(owner_name)"`
	AccessToken string
	Created     time.Time  `xorm:"CREATED"`
	Updated     time.Time  `xorm:"UPDATED"`
	Deleted     time.Time  `xorm:"deleted"`
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
	a.Name = *user.Login
}

func (a *GithubAccount) Save() {
	if a.ID == 0 {
		// find same user
		cond := &GithubAccount{
			OwnerID: a.OwnerID,
			Name:    a.Name,
		}
		has, _ := x.Get(cond)
		if has {
			// got it
			x.Id(cond.ID).Update(a)
			return
		}
		x.Insert(a)
	} else {
		x.Id(a.ID).Update(a)
	}
}
