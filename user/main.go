package main

import (
	"fmt"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	user_handler "user-srv/handler"
	user_proto "user-srv/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user"),
	)

	// Register handler
	err := user_proto.RegisterUserHandler(srv.Server(), user_handler.New())

	if err != nil {
		fmt.Println("错误启动")
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}