package controllers

import (
	"github.com/Miniand/miniand-network/app/models"
	"github.com/Miniand/miniand-network/app/routes"
	"github.com/robfig/revel"
)

type Products struct {
	Application
}

func (c Products) Index() revel.Result {
	var products []*models.Product
	_, err := c.Txn.Select(&products, "select * from Product")
	if err != nil {
		revel.ERROR.Fatalf("Could not select products: %s", err.Error())
	}
	return c.Render(products)
}

func (c Products) New() revel.Result {
	return c.Render()
}

func (c Products) Create(p models.Product) revel.Result {
	p.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Products.New())
	}
	err := c.Txn.Insert(&p)
	if err != nil {
		panic(err)
	}
	return c.Redirect(routes.Products.Index())
}
