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
	products, err := models.AllProducts(c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not get products: %s", err.Error())
	}
	return c.Render(products)
}

func (c Products) AdminNew() revel.Result {
	return c.Render()
}

func (c Products) Create(p models.Product) revel.Result {
	err := models.CreateProduct(&p, c.Validation, c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not create product: %s", err.Error())
	}
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Products.AdminNew())
	}
	return c.Redirect(routes.Products.AdminShow(p.Id))
}

func (c Products) Delete(id int64) revel.Result {
	err := models.DeleteProduct(id, c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not delete product %d: %s", id, err.Error())
	}
	return c.Redirect(routes.Products.AdminIndex())
}

func (c Products) Update(id int64, p models.Product) revel.Result {
	err := models.UpdateProduct(&p, c.Validation, c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not update product: %s", err.Error())
	}
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Products.AdminEdit(id))
	}
	return c.Redirect(routes.Products.AdminShow(id))
}

func (c Products) AdminIndex() revel.Result {
	products, err := models.AllProducts(c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not get products: %s", err.Error())
	}
	return c.Render(products)
}

func (c Products) AdminEdit(id int64) revel.Result {
	product, err := models.FindProduct(id, c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not load product %d for editing: %s",
			err.Error())
	}
	if product == nil {
		return c.Redirect(routes.Products.AdminIndex())
	}
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
	product, err := models.FindProduct(id, c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not load product %d for viewing: %s",
			err.Error())
	}
	if product == nil {
		return c.Redirect(routes.Products.AdminIndex())
	}
	productShops, err := models.AllShopProductsForProduct(id, c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not select product shops: %s", err.Error())
	}
	return c.Render(product, productShops)
}
