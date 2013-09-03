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

func (sp *ShopProduct) Validate(v *revel.Validation, txn *gorp.Transaction) {
	var shopProducts []*ShopProduct
	if sp.ShopId == 0 {
		v.Error("You must choose a shop").Key("sp.ShopId")
	}
	if sp.ProductId == 0 {
		v.Error("You must choose a product").Key("sp.ProductId")
	}
	if _, err := txn.Select(&shopProducts, `
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
