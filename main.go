package main

import (
	"stalcraftbot/cmd"
	"stalcraftbot/pkg/postgres"
)

func main() {

	cmd.Execute()
	postgres.CheckAndMigrate()

}
