package workspace

import (
	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/web/middleware"
)

func Index(ctx *middleware.Context) {
	workspaces, _ := models.Workspaces()
	ctx.Data["Workspaces"] = workspaces
	ctx.HTML(200, "workspace/index")
}
