package docker

import (
	"github.com/go-macaron/csrf"
	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/web/middleware"
)

func Index(ctx *middleware.Context) {
	servers, err := models.DockerServers()
	if err != nil {
		return
	}

	ctx.Data["Servers"] = servers
	ctx.HTML(200, "docker/index")
}

func New(ctx *middleware.Context, x csrf.CSRF) {
	ctx.Data["csrf_token"] = x.GetToken()
	ctx.HTML(200, "docker/new")
}

type NewForm struct {
	Name     string `binding:"Required;MaxSize(254)"`
	Endpoint string `binding:"Required;MaxSize(254)"`
}

func NewPost(ctx *middleware.Context, form NewForm) {
	if ctx.HasError() {
		ctx.Redirect("/docker", 302)
		return
	}

	d := models.NewDockerServer(form.Name, form.Endpoint)
	d.Save()
	ctx.Redirect("/docker")
}
