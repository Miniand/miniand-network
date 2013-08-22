package tests

import (
	"github.com/robfig/revel"
	"net/url"
)

type ProductsTest struct {
	revel.TestSuite
}

func (t ProductsTest) TestThatIndexPageWorks() {
	t.Get("/products/")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t ProductsTest) TestThatNewPageWorks() {
	t.Get("/products/new/")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t ProductsTest) TestThatCreateWorks() {
	t.PostForm("/products/", url.Values{})
	t.AssertOk()
	t.AssertContentType("text/html")
}
