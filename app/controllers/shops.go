package controllers

import (
	"github.com/Miniand/miniand-network/app/models"
	"github.com/Miniand/miniand-network/app/routes"
	"github.com/robfig/revel"
)

type Shops struct {
	Application
}

func (c Shops) Index() revel.Result {
	var shops []*models.Shop
	_, err := c.Txn.Select(&shops, "select * from Shop")
	if err != nil {
		revel.ERROR.Fatalf("Could not select shops: %s", err.Error())
	}
	return c.Render(shops)
}

func (c Shops) New() revel.Result {
	return c.Render()
}

func (c Shops) Create(s models.Shop) revel.Result {
	s.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Shops.New())
	}
	err := c.Txn.Insert(&s)
	if err != nil {
		panic(err)
	}
	return c.Redirect(routes.Shops.Index())
}
