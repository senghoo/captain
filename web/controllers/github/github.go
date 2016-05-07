package github

import (
	"github.com/senghoo/captain/models"
	gh "github.com/senghoo/captain/modules/github"
	"github.com/senghoo/captain/web/middleware"
)

const accountPerPage = 30

func Auth(ctx *middleware.Context) {
	url, state := gh.AuthCodeURL()
	ctx.Session.Set("github_state", state)
	ctx.Redirect(url)
}

func Callback(ctx *middleware.Context) {
	state, ok := ctx.Session.Get("github_state").(string)
	if !ok {
		ctx.Redirect("/")
		return
	}

	if state != ctx.Query("state") {
		ctx.Redirect("/")
		return
	}

	code := ctx.Query("code")
	token, err := gh.Exchange(code)
	if err != nil {
		ctx.Redirect("/")
		return
	}

	a := models.NewGithubAccount(token)
	a.Save()
	ctx.Session.Delete("github_state")
	ctx.Redirect("/")
}

func List(ctx *middleware.Context) {
	accounts, err := models.GithubAccounts()
	if err != nil {
		return
	}
	ctx.Data["GithubAccounts"] = accounts

	ctx.HTML(200, "github/list")
}
