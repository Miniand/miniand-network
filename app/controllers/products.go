package controllers

import (
	"github.com/Miniand/miniand-network/app/models"
	"github.com/robfig/revel"
)

type Products struct {
	*revel.Controller
}

func (c Products) Index() revel.Result {
	p := &models.Product{
		Name: "Fart",
	}
	err := Dbm.Insert(p)
	if err != nil {
		revel.ERROR.Fatal(err)
	}
	return c.Render()
}
