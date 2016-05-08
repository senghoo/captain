package docker

import "github.com/senghoo/captain/web/middleware"

func Index(ctx *middleware.Context) {
	ctx.HTML(200, "docker/index")
}
