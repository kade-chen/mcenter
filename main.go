package main

import (
	_ "github.com/kade-chen/mcenter/apps" //registry all impl/api

	// 开启Health健康检查
	_ "github.com/kade-chen/library/ioc/apps/health/restful"

	"github.com/kade-chen/mcenter/cmd"
)

func main() {
	cmd.Execute()
}
