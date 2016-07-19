package workspace

import (
	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/web/middleware"
)

func BuildStatus(ctx *middleware.Context) {
	build := new(models.Build)
	id := ctx.ParamsInt64(":id")
	has, err := models.GetByID(id, build)
	if !has {
		ctx.NotFound("")
		return
	}

	if err != nil {
		ctx.HandleErr(err, "")
		return
	}

	ctx.Data["Build"] = build
	ctx.HTML(200, "workspace/build_status")
}
