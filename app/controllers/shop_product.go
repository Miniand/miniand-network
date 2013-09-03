package controllers

import (
	"github.com/Miniand/miniand-network/app/models"
	"github.com/Miniand/miniand-network/app/routes"
	"github.com/robfig/revel"
	"strconv"
)

type ShopProducts struct {
	Application
}

func (c ShopProducts) AdminNew(productId int, shopId int,
	returnUrl string) revel.Result {
	var (
		products []*models.Product
		shops    []*models.Shop
	)
	if _, err := c.Txn.Select(&products, "select * from Product"); err != nil {
		revel.ERROR.Fatalf("Could not select products: %s", err.Error())
	}
	if _, err := c.Txn.Select(&shops, "select * from Shop"); err != nil {
		revel.ERROR.Fatalf("Could not select shops: %s", err.Error())
	}
	if productId > 0 {
		c.Flash.Data["sp.ProductId"] = strconv.Itoa(productId)
	}
	if shopId > 0 {
		c.Flash.Data["sp.ShopId"] = strconv.Itoa(shopId)
	}
	return c.Render(products, shops, returnUrl)
}

func (c ShopProducts) Create(sp models.ShopProduct) revel.Result {
	returnUrl := c.Params.Get("returnUrl")
	sp.Validate(c.Validation, c.Txn)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.ShopProducts.AdminNew(0, 0, returnUrl))
	}
	err := c.Txn.Insert(&sp)
	if err != nil {
		panic(err)
	}
	if returnUrl != "" {
		return c.Redirect(returnUrl)
	}
	return c.Redirect(routes.Products.AdminShow(sp.ProductId))
}

func (c ShopProducts) Delete(id int64) revel.Result {
	_, err := c.Txn.Delete(&models.ShopProduct{Id: id})
	if err != nil {
		revel.ERROR.Fatalf("Could not delete product from shop %d: %s", id,
			err.Error())
	}
	returnUrl := c.Params.Get("returnUrl")
	if returnUrl != "" {
		return c.Redirect(returnUrl)
	}
	return c.Redirect(routes.Products.AdminIndex())
}
