package models

import (
	"github.com/Miniand/gorp"
	"time"
)

type ProductVariant struct {
	Id          int64
	CreatedAt   int64
	UpdatedAt   int64
	ProductId   int64
	Name        string
	Active      bool
	Price       int64
	Description string
	SKU         string
	Weight      int64
}

func (pv *ProductVariant) PreInsert(s gorp.SqlExecutor) error {
	pv.CreatedAt = time.Now().UnixNano()
	pv.UpdatedAt = pv.CreatedAt
	return nil
}

func (pv *ProductVariant) PreUpdate(s gorp.SqlExecutor) error {
	pv.UpdatedAt = time.Now().UnixNano()
	return nil
}
