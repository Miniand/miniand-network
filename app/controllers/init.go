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
	revel.TemplateFuncs["postLink"] = func(url string, confirmation string) (
		template.HTML, error) {
		return template.HTML(fmt.Sprintf(
			`<form style="display:inline;" action="%s" method="POST"><a href="#" onclick="$(this).closest('form').submit();return false;">`,
			template.HTMLEscapeString(url))), nil
	}
	revel.TemplateFuncs["endPostLink"] = func() (template.HTML, error) {
		return template.HTML("</a></form>"), nil
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
