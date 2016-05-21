package workspace

import (
	"github.com/go-macaron/csrf"
	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/web/middleware"
)

func Index(ctx *middleware.Context) {
	workspaces, _ := models.Workspaces()
	ctx.Data["Workspaces"] = workspaces
	ctx.HTML(200, "workspace/index")
}

func New(ctx *middleware.Context, x csrf.CSRF) {
	ctx.Data["csrf_token"] = x.GetToken()
	ctx.HTML(200, "workspace/new")
}

type NewForm struct {
	Name string `binding:"Required;MaxSize(254)"`
}

func NewPost(ctx *middleware.Context, form NewForm) {
	if ctx.HasError() {
		ctx.Redirect("/", 302)
		return
	}

	d := models.NewWorkspace(form.Name)
	d.Save()
	ctx.Redirect("/workspace")
}
