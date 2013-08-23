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

func (c *Application) DetectShopIdentifierInHost() revel.Result {
	// Defaults
	c.RenderArgs["shopHue"] = 125
	r := regexp.MustCompile(`\.`)
	hostParts := r.Split(c.Request.Host, 2)
	if len(hostParts) > 1 {
		c.Shop = models.FindShopByIdentifier(hostParts[0], c.Txn)
		c.RenderArgs["shop"] = c.Shop
		if c.Shop != nil {
			c.RenderArgs["shopHue"] = c.Shop.Hue
		}
	}
	return nil
}

func (c *Application) FetchShopList() revel.Result {
	var shops []*models.Shop
	_, err := c.Txn.Select(&shops, "select * from Shop")
	if err != nil {
		revel.ERROR.Fatalf("Could not select shops: %s", err.Error())
	}
	c.RenderArgs["shops"] = shops
	return nil
}

func (c Application) Index() revel.Result {
	return c.Render()
}
