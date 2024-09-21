package postgres

import (
	"database/sql"
	"stalcraftbot/internal/logs"

	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
)

func CheckAndMigrate() {
	migrateFile := "./"

	connStr := viper.GetString("databaseurl")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logs.Logger.Fatal().Err(err).Msg("Open DB Error:")
	}
	var exists bool

	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users')").Scan(&exists)

	if err != nil {
		logs.Logger.Fatal().Err(err).Msg("Check table DB Error:")
	}

	if !exists {

		if err := goose.Up(db, migrateFile); err != nil {
			logs.Logger.Fatal().Err(err).Msg("Failed to run migrations:")
		}

	} else {
		logs.Logger.Debug().Msg("Table exists. Migration skip")
	}
}
