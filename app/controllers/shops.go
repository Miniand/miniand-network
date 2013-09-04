package controllers

import (
	"fmt"
	"github.com/Miniand/miniand-network/app/models"
	"github.com/Miniand/miniand-network/app/routes"
	"github.com/robfig/revel"
)

type Shops struct {
	Application
}

func (c Shops) AdminIndex() revel.Result {
	shops, err := models.AllShops(c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not select shops: %s", err.Error())
	}
	return c.Render(shops)
}

func (c Shops) AdminNew() revel.Result {
	return c.Render()
}

func (c Shops) AdminShow(id int64) revel.Result {
	shop, err := models.FindShop(id, c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not load shop %d for viewing: %s",
			err.Error())
	}
	if shop == nil {
		return c.Redirect(routes.Shops.AdminIndex())
	}
	productShops, err := models.AllShopProductsForShop(id, c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not select product shops: %s", err.Error())
	}
	return c.Render(shop, productShops)
}

func (c Shops) Create(s models.Shop) revel.Result {
	err := models.CreateShop(&s, c.Validation, c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not create shop: %s", err.Error())
	}
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Shops.AdminNew())
	}
	return c.Redirect(ShopUrl(s.Identifier, routes.Shops.AdminShow(s.Id)))
}

func (c Shops) Delete(id int64) revel.Result {
	if err := models.DeleteShop(id, c.Txn); err != nil {
		revel.ERROR.Fatalf("Could not delete shop %d: %s", id, err.Error())
	}
	return c.Redirect(ShopUrl("", routes.Shops.AdminIndex()))
}

func (c Shops) AdminEdit(id int64) revel.Result {
	shop, err := models.FindShop(id, c.Txn)
	if err != nil {
		revel.ERROR.Fatalf("Could not load shop %d for editing: %s", err.Error())
	}
	if shop == nil {
		return c.Redirect(routes.Shops.AdminIndex())
	}
	// Set flash data to initialise form
	for key, val := range shop.ToStringMap() {
		prefixedKey := fmt.Sprintf("s.%s", key)
		if c.Flash.Data[prefixedKey] == "" {
			c.Flash.Data[prefixedKey] = val
		}
	}
	return c.Render(shop)
}

func (c Shops) Update(id int64, s models.Shop) revel.Result {
	if err := models.UpdateShop(&s, c.Validation, c.Txn); err != nil {
		revel.ERROR.Fatalf("Could not update shop %d: %s", id, err.Error())
	}
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Shops.AdminEdit(id))
	}
	return c.Redirect(ShopUrl(s.Identifier, routes.Shops.AdminIndex()))
}
