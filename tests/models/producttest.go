package tests

import (
	"github.com/Miniand/miniand-network/app/models"
	"github.com/robfig/revel"
)

type ProductTest struct {
	revel.TestSuite
}

func (t ProductTest) TestThatProductSaves() {
	p := models.Product{}
	t.AssertEqual(0, p.Id)
}
