package models

import (
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
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

func GithubClient(token string) *github.Client {
	t := new(oauth2.Token)
	json.Unmarshal([]byte(token), t)
	oauthClient := oauthConf.Client(oauth2.NoContext, t)
	return github.NewClient(oauthClient)
}

type GithubAccount struct {
	ID          int64
	Name        string `xorm:"not null unique"`
	AccessToken string
	Created     time.Time      `xorm:"CREATED"`
	Updated     time.Time      `xorm:"UPDATED"`
	Deleted     time.Time      `xorm:"deleted"`
	client      *github.Client `xorm:"-"`
}

func GetGithubAccountByID(id int64) (*GithubAccount, error) {
	s := new(GithubAccount)
	has, err := x.Id(id).Get(s)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("Docker server id: %d not exist", id)
	}
	return s, nil
}

func NewGithubAccount(token string) (a *GithubAccount) {
	a = &GithubAccount{
		AccessToken: token,
	}
	a.UpdateName()
	return
}

func (a *GithubAccount) Client() *github.Client {
	if a.client == nil {
		a.client = GithubClient(a.AccessToken)
	}
	return a.client
}

func (a *GithubAccount) GithubUser() (user *github.User, err error) {
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

func (a *GithubAccount) Repos() (r []*Repository, err error) {
	repos, _, err := a.Client().Repositories.List("", nil)
	if err != nil {
		return
	}

	emptyIfNil := func(s *string) string {
		if s != nil {
			return *s
		}
		return ""
	}

	for _, repo := range repos {
		n := &Repository{
			Name:          emptyIfNil(repo.Name),
			Site:          "github",
			FullName:      emptyIfNil(repo.FullName),
			Description:   emptyIfNil(repo.Description),
			DefaultBranch: emptyIfNil(repo.DefaultBranch),
			MasterBranch:  emptyIfNil(repo.MasterBranch),
			Homepage:      emptyIfNil(repo.Homepage),
			Language:      emptyIfNil(repo.Language),
			CreatedAt:     repo.CreatedAt.Time,
			PushedAt:      repo.PushedAt.Time,
			UpdatedAt:     repo.UpdatedAt.Time,
			HTMLURL:       emptyIfNil(repo.HTMLURL),
			CloneURL:      emptyIfNil(repo.CloneURL),
			GitURL:        emptyIfNil(repo.GitURL),
			SSHURL:        emptyIfNil(repo.SSHURL),
		}
		if repo.Homepage != nil {
			n.Homepage = *repo.Homepage
		}
		if repo.Language != nil {
			n.Language = *repo.Language
		}

		if repo.MasterBranch != nil {
			n.MasterBranch = *repo.MasterBranch
		}

		r = append(r, n)
	}
	return
}

func GithubAccounts() ([]*GithubAccount, error) {
	var accounts []*GithubAccount
	return accounts, x.Asc("id").Find(&accounts)
}

func CountGithubAccounts() int64 {
	count, _ := x.Count(new(GithubAccount))
	return count
}
