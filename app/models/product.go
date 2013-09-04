package models

import (
	"github.com/Miniand/gorp"
	"github.com/robfig/revel"
	"strconv"
	"time"
)

type Product struct {
	Id          int64
	CreatedAt   int64
	UpdatedAt   int64
	Name        string
	Active      bool
	Description string
}

func (p *Product) PreInsert(s gorp.SqlExecutor) error {
	p.CreatedAt = time.Now().UnixNano()
	p.UpdatedAt = p.CreatedAt
	return nil
}

func (p *Product) PreUpdate(s gorp.SqlExecutor) error {
	p.UpdatedAt = time.Now().UnixNano()
	return nil
}

func (p *Product) Validate(v *revel.Validation) {
	v.Required(p.Name)
	v.Required(p.Description)
}

func (p *Product) ToStringMap() map[string]string {
	activeText := ""
	if p.Active {
		activeText = "true"
	}
	return map[string]string{
		"Id":          strconv.Itoa(int(p.Id)),
		"Name":        p.Name,
		"Active":      activeText,
		"Description": p.Description,
	}
}

func AllProducts(txn *gorp.Transaction) (products []*Product, err error) {
	_, err = txn.Select(&products, "SELECT * FROM Product")
	return
}

func CreateProduct(p *Product, v *revel.Validation, txn *gorp.Transaction) error {
	p.Validate(v)
	if v.HasErrors() {
		return nil
	}
	return txn.Insert(p)
}

func DeleteProduct(id int64, txn *gorp.Transaction) error {
	_, err := txn.Delete(&Product{Id: id})
	return err
}

func UpdateProduct(p *Product, v *revel.Validation, txn *gorp.Transaction) error {
	p.Validate(v)
	if v.HasErrors() {
		return nil
	}
	_, err := txn.Update(p)
	return err
}

func FindProduct(id int64, txn *gorp.Transaction) (*Product, error) {
	m, err := txn.Get(Product{}, id)
	if err != nil {
		revel.ERROR.Fatalf("Could not find product %d: %s", id, err.Error())
	}
	if m == nil {
		return nil, nil
	}
	return m.(*Product), nil
}
