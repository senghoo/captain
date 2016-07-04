package workspace

import (
	"fmt"
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
	ctx.Data["Repos"], _ = ws.Repositories()
	ctx.Data["Builds"], _ = ws.Builds()
	ctx.Data["Workflows"], _ = ws.Workflows()
	ctx.HTML(200, "workspace/info")
}

func AddRepository(ctx *middleware.Context, x csrf.CSRF) {
	id := ctx.ParamsInt64(":id")
	github, err := models.GetGithubAccount()

	if err != nil {
		ctx.HandleErr(err, fmt.Sprintf("/workspace/%d", id))
		return
	}
	if github != nil {
		ctx.Data["GithubAccounts"] = github
		ctx.Data["GithubRepos"], _ = github.Repos()
	}

	ctx.Data["csrf_token"] = x.GetToken()

	ctx.HTML(200, "workspace/new_repo")
}

type AddRepositoryForm struct {
	RepoIdentify string `binding:"Required;MaxSize(254)"`
}

func PostAddRepository(ctx *middleware.Context, form AddRepositoryForm) {
	repo, err := models.GetRepositoryByIdentify(form.RepoIdentify)

	if err != nil {
		ctx.HandleErr(err, "/workspace")
		return
	}

	ws := new(models.Workspace)
	id := ctx.ParamsInt64(":id")
	has, err := models.GetByID(id, ws)
	if !has {
		ctx.NotFound("")
		return
	}

	if err != nil {
		ctx.HandleErr(err, "")
		return
	}

	ws.AddRepository(repo)
	ctx.Redirect(fmt.Sprintf("/workspace/%d", id), 302)
}

func AddWorkflow(ctx *middleware.Context, x csrf.CSRF) {
	ctx.Data["csrf_token"] = x.GetToken()
	ctx.HTML(200, "workspace/new_workflow")
}

type AddWorkflowForm struct {
	Name       string `binding:"Required;MaxSize(254)"`
	Config     string `binding:"Required;MaxSize(4096)"`
	ConfigType string `binding:"Required;MaxSize(64)"`
}

func PostAddWorkflow(ctx *middleware.Context, form AddWorkflowForm) {
	wsID := ctx.ParamsInt64(":id")
	wf := models.NewWorkflow(wsID, form.Name, form.Config, form.ConfigType)
	models.Insert(wf)
	ctx.Redirect(fmt.Sprintf("/workspace/%d", wsID), 302)
}
