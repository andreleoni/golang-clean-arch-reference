package postgres

import (
	"fmt"
	"log"

	"github.com/go-pg/pg"
)

var PG *pg.DB

func PGSetup() {
	PG = pg.Connect(&pg.Options{
		User:     "gophercon_user",
		Password: "gophercon_password",
		Database: "gophercon_db",
		Addr:     "localhost:5432",
	})

	// Development only db seed to improve development agility
	customerSeed()

	fmt.Println("Seed completed...")
}

func PGClose() {
	PG.Close()
}

func customerSeed() {
	seeds := "DROP TABLE IF EXISTS customers;" +

		"CREATE TABLE IF NOT EXISTS customers (" +
		"id character varying," +
		"name character varying" +
		");"

	fmt.Println("Running customers seed...")

	_, err := PG.Exec(seeds)
	if err != nil {
		log.Fatal(err)
	}
}
