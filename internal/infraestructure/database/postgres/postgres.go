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
	orderSeed()
	productSeed()

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

func orderSeed() {
	seeds := "DROP TABLE IF EXISTS orders;" +

		"CREATE TABLE IF NOT EXISTS orders (" +
		"id character varying," +
		"product_id character varying," +
		"customer_id character varying," +
		"quantity bigint," +
		"address_street character varying," +
		"address_number character varying," +
		"address_complement character varying," +
		"address_zipcode character varying," +
		"address_city character varying," +
		"address_province character varying" +
		");"

	fmt.Println("Running order seed...")

	_, err := PG.Exec(seeds)
	if err != nil {
		log.Fatal(err)
	}
}

func productSeed() {
	seeds := "DROP TABLE IF EXISTS products;" +

		"CREATE TABLE IF NOT EXISTS products (" +
		"id character varying," +
		"name character varying," +
		"status character varying," +
		"price bigint" +
		");"

	fmt.Println("Running product seed...")

	_, err := PG.Exec(seeds)
	if err != nil {
		log.Fatal(err)
	}
}
