package models

import (
	"github.com/Miniand/gorp"
	"github.com/robfig/revel"
	"time"
)

type ShopProduct struct {
	Id        int64
	CreatedAt int64
	UpdatedAt int64
	ShopId    int64
	ProductId int64
}

type shopProductForProduct struct {
	Id     int64
	ShopId int64
	Hue    int
	Name   string
}

type shopProductForShop struct {
	Id        int64
	ProductId int64
	Name      string
}

func (sp *ShopProduct) Validate(v *revel.Validation, exe gorp.SqlExecutor) {
	var shopProducts []*ShopProduct
	if sp.ShopId == 0 {
		v.Error("You must choose a shop").Key("sp.ShopId")
	}
	if sp.ProductId == 0 {
		v.Error("You must choose a product").Key("sp.ProductId")
	}
	if _, err := exe.Select(&shopProducts, `
SELECT * from ShopProduct
WHERE Id <> ?
AND ShopId = ?
AND ProductId = ?
	`, sp.Id, sp.ShopId, sp.ProductId); err != nil {
		revel.ERROR.Fatalf("Could not select products: %s", err.Error())
	}
	if len(shopProducts) > 0 {
		v.Error("This product already exists in this shop").Key("sp.ProductId")
	}
}

func (sp *ShopProduct) PreInsert(s gorp.SqlExecutor) error {
	sp.CreatedAt = time.Now().UnixNano()
	sp.UpdatedAt = sp.CreatedAt
	return nil
}

func (sp *ShopProduct) PreUpdate(s gorp.SqlExecutor) error {
	sp.UpdatedAt = time.Now().UnixNano()
	return nil
}

func AllShopProductsForProduct(id int64, exe gorp.SqlExecutor) (
	shopProducts []*shopProductForProduct, err error) {
	_, err = exe.Select(&shopProducts, `
SELECT sp.Id, sp.ShopId, s.Hue, s.Name
FROM ShopProduct sp
INNER JOIN Shop s
ON sp.ShopId = s.Id
WHERE sp.ProductId = ?
	`, id)
	return
}

func AllShopProductsForShop(id int64, exe gorp.SqlExecutor) (
	shopProducts []*shopProductForShop, err error) {
	_, err = exe.Select(&shopProducts, `
SELECT sp.Id, sp.ProductId, p.Name
FROM ShopProduct sp
INNER JOIN Product p
ON sp.ProductId = p.Id
WHERE sp.ShopId = ?
	`, id)
	return
}

func CreateShopProduct(sp *ShopProduct, v *revel.Validation,
	exe gorp.SqlExecutor) error {
	sp.Validate(v, exe)
	if v.HasErrors() {
		return nil
	}
	return exe.Insert(sp)
}

func DeleteShopProduct(id int64, exe gorp.SqlExecutor) error {
	_, err := exe.Delete(&ShopProduct{Id: id})
	return err
}
