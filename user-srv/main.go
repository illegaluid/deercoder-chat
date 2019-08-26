package main

import (
	"deercoder-chat/user-srv/handler"
	user "deercoder-chat/user-srv/proto"
	"github.com/dreamlu/go-tool"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"time"
)


func main() {

	//consul
	registry := consul.NewRegistry(consul.Config(
		&api.Config{
			Address: der.GetDevModeConfig("consul.address"),
			Scheme:  der.GetDevModeConfig("consul.scheme"),
		}))

	service := micro.NewService(
		micro.Name("deercoder-chat.user"),
		micro.Registry(registry),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Address(":"+der.GetDevModeConfig("http_port")),
	)

	// service init
	service.Init()

	// start DB
	der.NewDB()

	// Register Handlers
	// user register
	_ = user.RegisterUserServiceHandler(service.Server(), new(handler.UserService))
	// login register
	_ = user.RegisterLoginServiceHandler(service.Server(), new(handler.LoginService))

	// run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
