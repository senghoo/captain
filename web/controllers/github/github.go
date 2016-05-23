package github

import (
	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/web/middleware"
)

const accountPerPage = 30

func Auth(ctx *middleware.Context) {
	url, state := models.GithubAuthCodeURL()
	ctx.Session.Set("github_state", state)
	ctx.Redirect(url)
}

func Callback(ctx *middleware.Context) {
	state, ok := ctx.Session.Get("github_state").(string)
	if !ok {
		ctx.Redirect("/github")
		return
	}

	if state != ctx.Query("state") {
		ctx.Redirect("/github")
		return
	}

	code := ctx.Query("code")
	token, err := models.GithubTokenExchange(code)
	if err != nil {
		ctx.Redirect("/github")
		return
	}

	a := models.NewGithubAccount(token)
	a.Save()
	ctx.Session.Delete("github_state")
	ctx.Redirect("/github")
}

func Info(ctx *middleware.Context) {
	a, err := models.GetGithubAccount()
	if err != nil {
		return
	}
	ctx.Data["Account"] = a
	ctx.Data["Repos"], _ = a.Repos()
	ctx.HTML(200, "github/info")
}
