package github

import (
	"encoding/json"

	gc "github.com/google/go-github/github"
	"github.com/senghoo/captain/modules/settings"
	"github.com/senghoo/captain/modules/utils"
	"golang.org/x/oauth2"
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

func AuthCodeURL() (url, state string) {
	state = utils.RandomString(10)
	url = oauthConf.AuthCodeURL(state, oauth2.AccessTypeOnline)
	return
}

func Exchange(code string) (token string, err error) {
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
