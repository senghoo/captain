package github

import (
	gh "github.com/senghoo/captain/modules/github"
	"github.com/senghoo/captain/web/middleware"
)

func Auth(ctx *middleware.Context) {
	url, state := gh.AuthCodeURL()
	ctx.Session.Set("github_state", state)
	ctx.Redirect(url)
}
