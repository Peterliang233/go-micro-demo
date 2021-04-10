package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/peterliang/demo/user/dao"
	"github.com/peterliang/demo/user/handler"
	"github.com/peterliang/demo/user/model"
	user "github.com/peterliang/demo/user/proto"
	"github.com/peterliang/demo/user/repository"
	"log"
)


func main(){
	err := config.LoadFile("./config.json")
	if err != nil {
		log.Fatalf("Could not load config file: %s",err.Error() )
		return
	}
	conf := config.Map()
	db, err := dao.CreateConnection(conf["mysql"].(map[string]interface{}))

	db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("connection error: %v\n", err)
	}
	repo := &repository.User{Db: db}
	service := micro.NewService(
		micro.Name("go.micro.srv.server"),
		micro.Version("v1.0.0"),
	)
	service.Init()

	_ = user.RegisterUserServiceHandler(service.Server(),&handler.UserService{Repo: repo})
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}