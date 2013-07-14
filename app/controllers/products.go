package controllers

import (
	"github.com/robfig/revel"
)

type Products struct {
	*revel.Controller
}

func (c Products) Index() revel.Result {
	return c.Render()
}
