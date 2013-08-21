package controllers

import (
	"database/sql"
	"github.com/Miniand/miniand-network/app/models"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/revel"
	"github.com/robfig/revel/modules/db/app"
	_ "github.com/ziutek/mymysql/mysql"
)

var (
	Dbm *gorp.DbMap
)

func Init() {
	var dialect gorp.Dialect

	db.Init()

	driverName, found := revel.Config.String("db.driver")
	if !found {
		revel.ERROR.Fatal("No db.driver found.")
	}
	switch driverName {
	case "sqlite3":
		dialect = gorp.SqliteDialect{}
	case "mysql":
		dialect = gorp.MySQLDialect{}
	case "postgres":
		dialect = gorp.PostgresDialect{}
	}
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: dialect}

	Dbm.AddTable(models.Permission{}).SetKeys(true, "Id")
	Dbm.AddTable(models.Product{}).SetKeys(true, "Id")
	Dbm.AddTable(models.ProductVariant{}).SetKeys(true, "Id")
	Dbm.AddTable(models.Shop{}).SetKeys(true, "Id")
	Dbm.AddTable(models.ShopProduct{}).SetKeys(true, "Id")
	Dbm.AddTable(models.User{}).SetKeys(true, "Id")

	Dbm.TraceOn("[gorp]", revel.INFO)
	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		revel.ERROR.Fatal("Unable to create tables: " + err.Error())
	}
}

type GorpController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
