package main

import (
	"generate-id/internal/routers"
	"generate-id/internal/utils"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	//获取机器machine-id
	utils.GetMachineID()
	//初始化sonyflake实例
	utils.SF()
	h := server.Default(server.WithExitWaitTime(10*time.Second), server.WithHostPorts(":9000"))

	routers.CustomizedRegister(h)
	h.Spin()
}
