package controllers

import (
	"Service-Content/errors"
	"Service-Content/services"
	"Service-Content/models"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	BaseController
}

func (c *UserController)Detail()  {
	id, err := c.GetInt64("id")
	if err != nil{
		c.ApiErrorReturn(errors.ErrParam)
		return
	}

	s := services.NewUserService(c.Ctx)
	user,errMsg := s.Detail(id)
	if errMsg != nil{
		c.ApiErrorReturn(errMsg)
		return
	}

	c.ApiSuccessReturn(user)
	return
}

func (c *UserController)Insert()  {
	user := new(models.UserModel)
	if err := c.ParseForm(user); err != nil{
		logs.Info(err)
		c.ApiErrorReturn(errors.ErrParam)
		return
	}

	s := services.NewUserService(c.Ctx)
	id,errMsg := s.Insert(user)
	if errMsg != nil{
		c.ApiErrorReturn(errMsg)
		return
	}
	data := map[string]int64{
		"id": id,
	}
	c.ApiSuccessReturn(data)
	return
}

func (c *UserController)Update()  {

}