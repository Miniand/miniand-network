package jobs

import (
	"github.com/Miniand/miniand-network/app/models"
	"log"
)

type InitialiseDb struct{}

func (d InitialiseDb) Run() {
	db, err := models.Db()
	if err != nil {
		panic(err.Error())
	}
	// Make sure we can connect
	err = db.Db.Ping()
	if err != nil {
		panic(err.Error())
	}
	log.Println("Can connect to database without issue")
	db.CreateTablesIfNotExists()
	log.Println("Database tables created")
}
