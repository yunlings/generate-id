package controllers

import (
	"context"
	"generate-id/internal/comm"
	"generate-id/internal/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

func Sonyflake(c context.Context, ctx *app.RequestContext) {
	ids, _ := utils.Sf.NextID()
	comm.ResSuccess(ctx, comm.SuccessCode, "success", ids)

}
