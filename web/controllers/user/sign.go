package user

import "gopkg.in/macaron.v1"

func SignIn(ctx *macaron.Context) {
	ctx.HTML(200, "user/sign_in")
}
