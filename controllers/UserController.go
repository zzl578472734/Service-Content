package controllers

import (
	"Service-Content/errors"
	"Service-Content/services"
)

type UserController struct {
	BaseController
}

func (c *UserController)Detail()  {

	id, err := c.GetInt64("id")
	if err != nil{
		c.ApiErrorReturn(errors.ErrQueryError)
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