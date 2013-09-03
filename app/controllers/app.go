package controllers

import (
	"bytes"
	"github.com/Miniand/miniand-network/app/models"
	"github.com/robfig/revel"
	"image"
	"image/color"
	"image/png"
	"regexp"
	"time"
)

const (
	FAVICON_SIZE = 64
	FAVICON_S    = 0.61
	FAVICON_L    = 0.55
)

// The last time the favicon logic was changed
var faviconTime = time.Date(2013, 8, 24, 0, 0, 0, 0, time.UTC)

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

func (c Application) Favicon(hue int) revel.Result {
	buf := bytes.NewBuffer([]byte{})
	icon := image.NewRGBA(image.Rect(0, 0, FAVICON_SIZE, FAVICON_SIZE))
	rgba := HSLToRGB(float64(hue)/360, FAVICON_S, FAVICON_L)
	faviconHalf := float64(FAVICON_SIZE / 2)
	for x := 0; x < FAVICON_SIZE; x++ {
		for y := 0; y < FAVICON_SIZE; y++ {
			pointRgba := rgba
			pointRgba.A = 0
			xx, yy, rr := float64(x)-faviconHalf+0.5,
				float64(y)-faviconHalf+0.5,
				faviconHalf
			if xx*xx+yy*yy < rr*rr {
				pointRgba.A = 255
			}
			icon.SetRGBA(x, y, pointRgba)
		}
	}
	png.Encode(buf, icon)
	return c.RenderBinary(buf, "favicon.png", revel.Inline, time.Now())
}

func HSLToRGB(h, s, l float64) (rgba color.RGBA) {
	var fR, fG, fB float64
	if s == 0 {
		fR, fG, fB = l, l, l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - s*l
		}
		p := 2*l - q
		fR = hueToRGB(p, q, h+1.0/3)
		fG = hueToRGB(p, q, h)
		fB = hueToRGB(p, q, h-1.0/3)
	}
	rgba.R = uint8((fR * 255) + 0.5)
	rgba.G = uint8((fG * 255) + 0.5)
	rgba.B = uint8((fB * 255) + 0.5)
	return
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6 {
		return p + (q-p)*6*t
	}
	if t < 0.5 {
		return q
	}
	if t < 2.0/3 {
		return p + (q-p)*(2.0/3-t)*6
	}
	return p
}
