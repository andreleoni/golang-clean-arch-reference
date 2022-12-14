package main

import (
	"golang-clean-arch-reference/internal/infraestructure/api"
	"golang-clean-arch-reference/internal/infraestructure/database/postgres"
)

func main() {
	postgres.PGSetup()

	api.Start()
}
