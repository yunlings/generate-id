package controllers

import (
	"context"
	"generate-id/internal/comm"
	"generate-id/internal/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

func GenerateUUIDWithoutHyphen(c context.Context, ctx *app.RequestContext) {
	ids := utils.GenerateUUIDWithoutHyphen()
	comm.ResSuccess(ctx, comm.SuccessCode, "success", ids)
}
