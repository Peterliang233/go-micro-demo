package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
	"github.com/peterliang/demo/user/model"
	user "github.com/peterliang/demo/user/proto"
	"net/http"
)

var (
	cl user.UserService
)

func Registry(c *gin.Context) {
	var u = model.User{}
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": u,
			"err": err,
		})
		return
	}
	response, err := cl.Register(context.TODO(), &user.RegisterRequest{
		User: &user.User{
			Id: uint32(u.ID),
			Name: u.Name,
			Password: u.Password,
			Phone: u.Phone,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": u,
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func Login(c *gin.Context) {
	var login model.Loginer
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 404,
			"detail": err,
		})
		return
	}
	resq, err := cl.Login(context.TODO(), &user.LoginRequest{
		Phone: login.Phone,
		Password: login.Password,
	})

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": login,
		"detail": resq,
	})
}

func UpdatePassword(c *gin.Context) {
	var u model.UpdatePassword
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail" : err,
		})
		return
	}
	resq, err := cl.UpdatePassword(context.TODO(), &user.UpdatePasswordRequest{
		Uid: u.Uid,
		OldPassword: u.OldPassword,
		NewPassword: u.NewPassword,
		ConfirmPassword: u.ConfirmPassword,
	})
	c.JSON(http.StatusOK, gin.H{
		"data": u,
		"detail" : resq,
	})
}

func main(){
	service := web.NewService(
		web.Name("go.micro.srv.client"),
		web.Version("v1.0.0"),

	)
	service.Init()

	cl = user.NewUserService("go.micro.srv.server", client.DefaultClient)
	router := gin.Default()
	router.POST("/registry", Registry)
	router.POST("/login", Login)
	router.POST("/update", UpdatePassword)
	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v",resq)
}
