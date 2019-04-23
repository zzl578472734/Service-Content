package controllers

import "Service-Content/errors"

type ErrorController struct {
	BaseController
}

func (c *ErrorController)Error404()  {
	c.ApiErrorReturn(errors.ErrUrlError)
	return
}

func (c *ErrorController)Error500()  {
	c.ApiErrorReturn(errors.ErrSysBusy)
	return
}

func (c *ErrorController)Error501()  {
	c.ApiErrorReturn(errors.ErrSysBusy)
	return
}


