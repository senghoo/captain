package github

import (
	"errors"
	"strconv"

	"github.com/senghoo/captain/commands"
	"github.com/senghoo/captain/models"
	"github.com/senghoo/captain/web/middleware"
)

func Webhook(ctx *middleware.Context) {
	wh := new(models.GithubWebhook)
	id, _ := strconv.ParseInt(ctx.Params(":id"), 10, 32)
	has, err := models.GetByID(id, wh)
	if !has {
		ctx.NotFound("")
		return
	}

	if err != nil {
		ctx.HandleErr(err, "")
		return
	}
	signature := ctx.GetHeader("x-hub-signature")
	if len(signature) == 0 {
		ctx.HandleErr(errors.New("Signature not set"), "")
	}

	body, err := ctx.Req.Body().Bytes()
	if err != nil {
		ctx.HandleErr(err, "")
		return
	}

	if wh.VerifySignature(signature, body) {
		wf, err := wh.Workflow()
		if err != nil {
			ctx.HandleErr(err, "")
			return
		}
		err = command.RunWorkflow(wf)
		if err != nil {
			ctx.HandleErr(err, "")
			return
		}
	}
}
