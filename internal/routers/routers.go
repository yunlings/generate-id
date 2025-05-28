package routers

import (
	"generate-id/internal/controllers"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func CustomizedRegister(h *server.Hertz) {
	openapi := h.Group("/openapi")
	{
		openapi.GET("/uuid4", controllers.Uuid4)
		openapi.GET("/generate", controllers.GenerateUUIDWithoutHyphen)
		openapi.GET("/sonyflake", controllers.Sonyflake)
	}

	h.GET("/system/checkOnline", controllers.CheckOnline)
	h.GET("/system/offline", controllers.Offline)
}
