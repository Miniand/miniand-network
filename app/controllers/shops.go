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
	var shops []*models.Shop
	_, err := c.Txn.Select(&shops, "select * from Shop")
	if err != nil {
		revel.ERROR.Fatalf("Could not select shops: %s", err.Error())
	}
	return c.Render(shops)
}

func (c Shops) AdminNew() revel.Result {
	return c.Render()
}

func (c Shops) Create(s models.Shop) revel.Result {
	s.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Shops.AdminNew())
	}
	err := c.Txn.Insert(&s)
	if err != nil {
		revel.ERROR.Fatalf("Could not create shop: %s", err.Error())
	}
	return c.Redirect(ShopUrl(s.Identifier, routes.Shops.AdminIndex()))
}

func (c Shops) Delete(id int64) revel.Result {
	_, err := c.Txn.Delete(&models.Shop{Id: id})
	if err != nil {
		revel.ERROR.Fatalf("Could not delete shop %d: %s", id, err.Error())
	}
	return c.Redirect(ShopUrl("", routes.Shops.AdminIndex()))
}

func (c Shops) AdminEdit(id int64) revel.Result {
	m, err := c.Txn.Get(models.Shop{}, id)
	if err != nil {
		revel.ERROR.Fatalf("Could not load shop %d for editing: %s", err.Error())
	}
	if m == nil {
		return c.Redirect(routes.Shops.AdminIndex())
	}
	shop := m.(*models.Shop)
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
	s.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Shops.AdminEdit(id))
	}
	_, err := c.Txn.Update(&s)
	if err != nil {
		revel.ERROR.Fatalf("Could not update shop %d: %s", id, err.Error())
	}
	return c.Redirect(ShopUrl(s.Identifier, routes.Shops.AdminIndex()))
}
