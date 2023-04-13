package main

import (
	"golang-clean-arch-reference/internal/infraestructure/api"
	"golang-clean-arch-reference/internal/infraestructure/database/sqlite"
)

func main() {
	sqlite.SQLiteSetup()

	api.Start()
}
