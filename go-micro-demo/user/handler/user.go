package handler

import (
	"context"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/peterliang/demo/user/model"
	user "github.com/peterliang/demo/user/proto"
	"github.com/peterliang/demo/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.User
}

func (e *UserService) Register(ctx context.Context, req *user.RegisterRequest, resp *user.Response) error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Name: req.User.Name,
		Phone: req.User.Phone,
		Password: string(hashedPwd),
	}

	if err := e.Repo.Create(user); err != nil {
		log.Log("create error")
		return err
	}
	resp.Msg = "create success"
	resp.Code = "200"
	return nil
}

func (e *UserService) Login(ctx context.Context, req *user.LoginRequest, resq *user.Response) error {
	user, err := e.Repo.FindByField("phone", req.Phone, "id, password")
	if err != nil {
		return err
	}

	if user == nil {
		return errors.Unauthorized("go.micro.srv.user, login", "该手机号码不存在")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.Unauthorized("go.micro.srv.micro.login", "用户密码错误")
	}

	resq.Code = "200"
	resq.Msg = "登录成功"
	return nil
}

func (e *UserService) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest, resq *user.Response) error {
	user, err := e.Repo.Find(int32(req.Uid))
	if user == nil {
		return errors.Unauthorized("go.micro.srv.micro.updatePassword", "用户不存在")
	}
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return errors.Unauthorized("go.micro.srv.UpdatePassword", "用户旧密码验证失败")
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost);
	if err != nil {
		log.Fatal("生成密码哈希值错误")
		return err
	}
	user.Password = string(hashedPwd);
	e.Repo.Update(user)
	resq.Msg = "密码修改成功"
	resq.Code = "200"
	return nil
}