package user

import (
	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/web/middleware"
)

func SignIn(ctx *middleware.Context) {
	ctx.HTML(200, "user/sign_in")
}

func SignInPost(ctx *middleware.Context, form SignInForm) {
	if ctx.HasError() {
		ctx.Redirect("/user/sign_in", 302)
		return
	}

	u, err := models.UserSignIn(form.UserName, form.Password)
	if err != nil {
		if models.IsErrUserNotExist(err) {
			ctx.Flash.Error("User and password unmached")
		}
		ctx.Redirect("/user/sign_in", 302)
		return
	}

	ctx.SetUser(u)
	ctx.Redirect("/", 302)
}

type SignInForm struct {
	UserName string `binding:"Required;MaxSize(254)"`
	Password string `binding:"Required;MaxSize(254)"`
	Remember bool
}

func SignOut(ctx *middleware.Context) {
	ctx.Logout()
	ctx.Redirect("/", 302)
}
