package controllers

import (
	"github.com/Miniand/miniand-network/app/models"
	"github.com/robfig/revel"
	"regexp"
)

type Application struct {
	GorpController
	Shop *models.Shop
}

func (c *Application) DetectShopNameInHost() revel.Result {
	r := regexp.MustCompile(`\.`)
	hostParts := r.Split(c.Request.Host, 2)
	if len(hostParts) > 1 {
		c.Shop = models.FindShopByName(hostParts[0], c.Txn)
	}
	return nil
}

func (c Application) Index() revel.Result {
	return c.Render()
}
