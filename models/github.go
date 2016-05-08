package models

import (
	"encoding/json"
	"time"

	"golang.org/x/oauth2"

	gc "github.com/google/go-github/github"
	"github.com/senghoo/captain/modules/settings"
	"github.com/senghoo/captain/modules/utils"
	githuboauth "golang.org/x/oauth2/github"
)

var oauthConf *oauth2.Config

func init() {
	oauthConf = &oauth2.Config{
		ClientID:     settings.GetOrDefault("github.client_id", ""),
		ClientSecret: settings.GetOrDefault("github.client_secret", ""),
		Scopes:       []string{"user:email"},
		Endpoint:     githuboauth.Endpoint,
	}
}

func GithubAuthCodeURL() (url, state string) {
	state = utils.RandomString(10)
	url = oauthConf.AuthCodeURL(state, oauth2.AccessTypeOnline)
	return
}

func GithubTokenExchange(code string) (token string, err error) {
	t, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return
	}
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return
	}
	token = string(jsonBytes)
	return
}

func GithubClient(token string) *gc.Client {
	t := new(oauth2.Token)
	json.Unmarshal([]byte(token), t)
	oauthClient := oauthConf.Client(oauth2.NoContext, t)
	return gc.NewClient(oauthClient)
}

type GithubAccount struct {
	ID          int64
	Name        string `xorm:"not null unique"`
	AccessToken string
	Created     time.Time  `xorm:"CREATED"`
	Updated     time.Time  `xorm:"UPDATED"`
	Deleted     time.Time  `xorm:"deleted"`
	client      *gc.Client `xorm:"-"`
}

func NewGithubAccount(token string) (a *GithubAccount) {
	a = &GithubAccount{
		AccessToken: token,
	}
	a.UpdateName()
	return
}

func (a *GithubAccount) Client() *gc.Client {
	if a.client == nil {
		a.client = GithubClient(a.AccessToken)
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
			Name: a.Name,
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

func GithubAccounts() ([]*GithubAccount, error) {
	var accounts []*GithubAccount
	return accounts, x.Asc("id").Find(&accounts)
}

func CountGithubAccounts() int64 {
	count, _ := x.Count(new(GithubAccount))
	return count
}
