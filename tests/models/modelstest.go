package tests

import (
	"github.com/Miniand/miniand-network/app/models"
	"github.com/robfig/revel"
)

type ModelsTest struct {
	revel.TestSuite
}

func (t ModelsTest) TestThatDbConnects() {
	_, err := models.Db()
	t.AssertEqual(nil, err)
}
