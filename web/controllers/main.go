package controllers

import "gopkg.in/macaron.v1"

func Main(ctx *macaron.Context) {
	ctx.HTML(200, "index")
}
