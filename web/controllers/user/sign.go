package user

import "gopkg.in/macaron.v1"

func SignIn(ctx *macaron.Context) {
	ctx.HTML(200, "user/sign_in")
}

func SignInPost(ctx *macaron.Context, form SignInForm) {
	ctx.HTML(200, "user/sign_in")
}

type SignInForm struct {
	UserName string `binding:"Required;MaxSize(254)"`
	Password string `binding:"Required;MaxSize(254)"`
	Remember bool
}
