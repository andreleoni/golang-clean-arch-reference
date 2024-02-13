package sqlite

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Sqlite *gorm.DB

func SQLiteSetup() {
	file, _ := os.OpenFile("gorm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	dbFilePath := "development.db"

	// If the file doesn't exist, create it

	// You can use a database creation command here if you're using a specific database
	// For example, if you're using SQLite, you can use the following command:
	// cmd := exec.Command("sqlite3", dbFilePath)

	// In this example, we'll just create an empty file for demonstration purposes
	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		file, err := os.Create(dbFilePath)
		if err != nil {
			fmt.Println("Error creating database file:", err)
			return
		}
		defer file.Close()

		fmt.Println("Database file created successfully.")
	}

	sqlite, err := gorm.Open(
		sqlite.Open(dbFilePath),
		&gorm.Config{
			Logger: logger.New(
				log.New(file, "\r\n", log.LstdFlags), // set log file and format
				logger.Config{
					LogLevel: logger.Info, // set log level
				},
			),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	Sqlite = sqlite

	// Development only db seed to improve development agility
	customerSeed()
	orderSeed()
	productSeed()

	fmt.Println("Seed completed...")
}

func customerSeed() {
	seeds := "DROP TABLE IF EXISTS customers;" +

		"CREATE TABLE IF NOT EXISTS customers (" +
		"id text," +
		"name text" +
		");"

	fmt.Println("Running customers seed...")

	result := Sqlite.Exec(seeds)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func orderSeed() {
	seeds := "DROP TABLE IF EXISTS orders;" +

		"CREATE TABLE IF NOT EXISTS orders (" +
		"id text," +
		"product_id text," +
		"customer_id text," +
		"quantity bigint," +
		"address_street text," +
		"address_number text," +
		"address_complement text," +
		"address_zipcode text," +
		"address_city text," +
		"address_province text" +
		");"

	fmt.Println("Running order seed...")

	result := Sqlite.Exec(seeds)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func productSeed() {
	seeds := "DROP TABLE IF EXISTS products;" +

		"CREATE TABLE IF NOT EXISTS products (" +
		"id text," +
		"name text," +
		"status text," +
		"price bigint" +
		");"

	fmt.Println("Running product seed...")

	result := Sqlite.Exec(seeds)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
