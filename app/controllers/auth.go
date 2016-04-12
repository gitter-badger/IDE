package controllers

import "github.com/revel/revel"

type Auth struct {
     *revel.Controller
}

func (c Auth) Index() revel.Result {
     return c.Render()
}

func (c Auth) Register() revel.Result {
	return c.Render()
}

