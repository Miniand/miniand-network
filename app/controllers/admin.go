package controllers

import (
	"github.com/Miniand/miniand-network/app/routes"
	"github.com/robfig/revel"
)

type Admin struct {
	Application
}

func (c Admin) Index() revel.Result {
	return c.Redirect(routes.Shops.AdminIndex())
}
