package main

//go:generate ./scripts/generate.sh

import (
	"github.com/micro/micro/v3/cmd"
	// load packages so they can register commands
	_ "github.com/micro/micro/v3/client/cli"
	_ "github.com/micro/micro/v3/cmd/server"
	_ "github.com/micro/micro/v3/cmd/service"
	_ "github.com/micro/micro/v3/cmd/usage"
)

func main() {
	cmd.Run()
}

/*
		micro server 提供micro platform功能，启动register broker network runtime config store event auth proxy api web 服务实现平台功能。
	命令入口是cmd包
		micro run . 部署微服务到平台中后能通过http api和grpc访问，是因为api组件在启动的时候就run了一个8080的http server去转发http请求到微服务中(endpoint)的grpc

*/
