package repo

import (
	"fmt"

	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/web/middleware"
)

func Clone(ctx *middleware.Context) {
	repo := new(models.Repository)
	id := ctx.ParamsInt64(":id")

	has, err := models.GetByID(id, repo)
	if !has {
		ctx.NotFound("")
		return
	}
	if err != nil {
		ctx.HandleErr(err, "")
		return
	}

	repo.Clone()

	ctx.Flash.Info("Clone processing")
	ctx.Redirect(fmt.Sprintf("/workspace/%d", id))
}
