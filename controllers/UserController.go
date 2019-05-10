package controllers

import (
	"Service-Content/errors"
	"Service-Content/services"
	"Service-Content/vars"
)

type UserController struct {
	BaseController
}

/**
 *
 */
func (c *UserController) Detail() {
	data, err := c.GetRequestParam().(*vars.DefaultIdQueryParam)
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

/**
 *
 */
func (c *UserController) Login() {
	param, err := c.GetRequestParam().(*vars.UserLoginParam)

	if !err {
		c.ApiErrorReturn(errors.ErrParam)
		return
	}

	service := services.NewUserService(c.Ctx)
	user, errMsg := service.Login(param)
	if errMsg != nil {
		c.ApiErrorReturn(errMsg)
		return
	}
	c.ApiSuccessReturn(user)
	return
}

func (c *UserController) Insert() {
	param, err := c.GetRequestParam().(*vars.UserInsertParam)
	if !err {
		c.ApiErrorReturn(errors.ErrParam)
		return
	}

	s := services.NewUserService(c.Ctx)
	id, errMsg := s.Insert(param)
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

func (c *UserController) Search() {
	param, err := c.GetRequestParam().(*vars.UserSearchParam)
	if !err {
		c.ApiErrorReturn(errors.ErrParam)
		return
	}

	service := services.NewUserService(c.Ctx)
	service.Search(param)
}

