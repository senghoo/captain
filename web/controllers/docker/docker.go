package docker

import "github.com/senghoo/captain/web/middleware"

func Index(ctx *middleware.Context) {
	ctx.HTML(200, "docker/index")
}

func New(ctx *middleware.Context) {
	ctx.HTML(200, "docker/new")
}

func Create(ctx *middleware.Context) {
}
