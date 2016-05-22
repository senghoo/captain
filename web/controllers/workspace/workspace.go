package workspace

import (
	"strconv"

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

func Info(ctx *middleware.Context) {
	ws := new(models.Workspace)
	id, _ := strconv.ParseInt(ctx.Params(":id"), 10, 32)
	has, err := models.GetByID(id, ws)
	if !has {
		ctx.NotFound("")
		return
	}

	if err != nil {
		ctx.HandleErr(err, "")
		return
	}
	ctx.Data["Workspace"] = ws
	ctx.HTML(200, "workspace/info")
}

func AddRepository(ctx *middleware.Context) {
	accounts, err := models.GithubAccounts()
	if err != nil {
		return
	}
	ctx.Data["GithubAccounts"] = accounts

	ctx.HTML(200, "workspace/repo_list")
}

type AddRepositoryForm struct {
	FullName string `binding:"Required;MaxSize(254)"`
	Source   string `binding:"Required;MaxSize(254)"`
}

func PostAddRepository(ctx *middleware.Context, form AddRepositoryForm) {

}
