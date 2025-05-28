package controllers

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// 自定义Http状态码
var statusCode = 200

// 设定Http状态码
func changeStatusCode() {
	statusCode = 555
}

func CheckOnline(c context.Context, ctx *app.RequestContext) {

	ctx.JSON(statusCode, utils.H{
		"message": "online",
	})
}

func Offline(c context.Context, ctx *app.RequestContext) {
	changeStatusCode()

	ctx.JSON(200, utils.H{
		"message": "offline",
	})
	time.Sleep(1 * time.Second)
}
