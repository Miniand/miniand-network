package controllers

import (
	"fmt"
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

func (c Products) AdminNew() revel.Result {
	return c.Render()
}

func (c Products) Create(p models.Product) revel.Result {
	p.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Products.AdminNew())
	}
	err := c.Txn.Insert(&p)
	if err != nil {
		panic(err)
	}
	return c.Redirect(routes.Products.AdminIndex())
}

func (c Products) Delete(id int64) revel.Result {
	_, err := c.Txn.Delete(&models.Product{Id: id})
	if err != nil {
		revel.ERROR.Fatalf("Could not delete product %d: %s", id, err.Error())
	}
	return c.Redirect(routes.Products.AdminIndex())
}

func (c Products) Update(id int64, p models.Product) revel.Result {
	p.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Products.AdminEdit(id))
	}
	_, err := c.Txn.Update(&p)
	if err != nil {
		revel.ERROR.Fatalf("Could not update product %d: %s", id, err.Error())
	}
	return c.Redirect(routes.Products.AdminIndex())
}

func (c Products) AdminIndex() revel.Result {
	var products []*models.Product
	_, err := c.Txn.Select(&products, "select * from Product")
	if err != nil {
		revel.ERROR.Fatalf("Could not select products: %s", err.Error())
	}
	return c.Render(products)
}

func (c Products) AdminEdit(id int64) revel.Result {
	m, err := c.Txn.Get(models.Product{}, id)
	if err != nil {
		revel.ERROR.Fatalf("Could not load product %d for editing: %s",
			err.Error())
	}
	if m == nil {
		return c.Redirect(routes.Products.AdminIndex())
	}
	product := m.(*models.Product)
	// Set flash data to initialise form
	for key, val := range product.ToStringMap() {
		prefixedKey := fmt.Sprintf("p.%s", key)
		if c.Flash.Data[prefixedKey] == "" {
			c.Flash.Data[prefixedKey] = val
		}
	}
	return c.Render(product)
}

func (c Products) AdminShow(id int64) revel.Result {
	m, err := c.Txn.Get(models.Product{}, id)
	if err != nil {
		revel.ERROR.Fatalf("Could not load product %d for editing: %s",
			err.Error())
	}
	if m == nil {
		return c.Redirect(routes.Products.AdminIndex())
	}
	product := m.(*models.Product)
	return c.Render(product)
}
