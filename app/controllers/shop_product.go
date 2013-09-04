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
	products, err := models.AllProducts(c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not get products: %s", err.Error())
	}
	shops, err := models.AllShops(c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not get shops: %s", err.Error())
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
	if err := models.CreateShopProduct(&sp, c.Validation, c.Txn); err != nil {
		revel.ERROR.Fatalf("Could not create shop product: %s", err.Error())
	}
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.ShopProducts.AdminNew(0, 0, returnUrl))
	}
	return c.Redirect(returnUrl)
}

func (c ShopProducts) Delete(id int64) revel.Result {
	if err := models.DeleteShopProduct(id, c.Txn); err != nil {
		revel.ERROR.Fatalf("Could not delete product from shop %d: %s", id,
			err.Error())
	}
	returnUrl := c.Params.Get("returnUrl")
	if returnUrl != "" {
		return c.Redirect(returnUrl)
	}
	return c.Redirect(routes.Products.AdminIndex())
}
