package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	user "github.com/peterliang/demo/user/proto"
)

func main(){
	service := micro.NewService(
		micro.Name("go.micro.srv.client"),
		micro.Version("v1.0.0"),
	)
	service.Init()

	cl := user.NewUserService("go.micro.srv.server", service.Client())

	resq, err := cl.Register(context.Background(), &user.RegisterRequest{
		User: &user.User{
			Id: 2,
			Name: "peter",
			Phone: "123",
			Password: "123",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v",resq)
}
