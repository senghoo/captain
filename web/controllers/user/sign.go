package user

import (
	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/web/middleware"
	"gopkg.in/macaron.v1"
)

func SignIn(ctx *macaron.Context) {
	ctx.HTML(200, "user/sign_in")
}

func SignInPost(ctx *middleware.Context, form SignInForm) {
	u, err := models.UserSignIn(form.UserName, form.Password)
	if err != nil {
		ctx.HTML(200, "user/sign_in")
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
