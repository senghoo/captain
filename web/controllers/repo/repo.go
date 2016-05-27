package repo

import (
	"strconv"

	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/web/middleware"
)

func Clone(ctx *middleware.Context) {
	repo := new(models.Repository)
	id, _ := strconv.ParseInt(ctx.Params(":id"), 10, 32)

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
}
