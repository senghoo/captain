package repo

import (
	"fmt"

	"github.com/senghoo/captain/commands"
	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/web/middleware"
)

func Build(ctx *middleware.Context) {
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

	workspace := repo.Workspace()
	build, _ := workspace.NewBuild("abc")

	node := command.AutoBuildTree(id, 1, "master", "test")
	go command.RunNode(node, build)

	ctx.Flash.Info("Build processing")
	ctx.Redirect(fmt.Sprintf("/workspace/%d", repo.Workspace().ID))
}
