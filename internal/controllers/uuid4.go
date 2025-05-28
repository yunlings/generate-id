package controllers

import (
	"context"
	"generate-id/internal/comm"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
)

func Uuid4(c context.Context, ctx *app.RequestContext) {
	ids := uuid.New()
	comm.ResSuccess(ctx, comm.SuccessCode, "success", ids)
}
