package controllers

import (
	"Service-Content/errors"
	"Service-Content/models"
	"Service-Content/services"
	"Service-Content/vars"
)

type UserController struct {
	BaseController
}

func (c *UserController) Detail() {
	data, err := c.GetRequestParam().(*vars.DefaultQueryParam)
	if !err {
		c.ApiErrorReturn(errors.ErrParam)
		return
	}

	s := services.NewUserService(c.Ctx)
	user, errMsg := s.Detail(data.Id)
	if errMsg != nil {
		c.ApiErrorReturn(errMsg)
		return
	}

	c.ApiSuccessReturn(user)
	return
}

func (c *UserController) Insert() {
	user := new(models.UserModel)
	if err := c.ParseForm(user); err != nil {
		c.ApiErrorReturn(errors.ErrParam)
		return
	}

	s := services.NewUserService(c.Ctx)
	id, errMsg := s.Insert(user)
	if errMsg != nil {
		c.ApiErrorReturn(errMsg)
		return
	}
	data := map[string]int64{
		"id": id,
	}
	c.ApiSuccessReturn(data)
	return
}
