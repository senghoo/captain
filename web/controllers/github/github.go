package github

import (
	"github.com/senghoo/captain/models"
	gh "github.com/senghoo/captain/modules/github"
	"github.com/senghoo/captain/web/middleware"
)

func Auth(ctx *middleware.Context) {
	url, state := gh.AuthCodeURL()
	ctx.Session.Set("github_state", state)
	ctx.Redirect(url)
}

func Callback(ctx *middleware.Context) {
	state, ok := ctx.Session.Get("github_state").(string)
	ctx.Session.Delete("github_state")
	if !ok {
		ctx.Redirect("/")
		return
	}

	if state != ctx.Params("state") {
		ctx.Redirect("/")
	}

	code := ctx.Params("code")
	token, err := gh.Exchange(code)
	if err != nil {
		ctx.Redirect("/")
	}

	a := models.NewGithubAccount(ctx.User.ID, token)
	a.Save()
}
