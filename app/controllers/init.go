package controllers

import (
	"fmt"
	"github.com/robfig/revel"
	"html/template"
)

func init() {
	revel.OnAppStart(Init)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
	revel.InterceptMethod((*Application).DetectShopIdentifierInHost,
		revel.BEFORE)
	revel.InterceptMethod((*Application).FetchShopList, revel.BEFORE)
	revel.TemplateFuncs["hue"] = func(hue int) template.CSS {
		return template.CSS(fmt.Sprintf("hsl(%d,61%%,55%%)", hue))
	}
	revel.TemplateFuncs["shopUrl"] = func(identifier string,
		args ...interface{}) (string, error) {
		path, err := revel.ReverseUrl(args...)
		if err != nil {
			return "", err
		}
		return ShopUrl(identifier, path), nil
	}
}

func ShopUrl(identifier string, path string) string {
	addr, _ := revel.Config.String("http.addr")
	port, _ := revel.Config.String("http.port")
	host := addr
	if identifier != "" {
		host = fmt.Sprintf("%s.%s", identifier, addr)
	}
	return fmt.Sprintf("//%s:%s%s", host, port, path)
}
