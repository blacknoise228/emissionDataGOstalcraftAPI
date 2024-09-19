package postgres

import (
	"stalcraftbot/internal/logs"

	"github.com/pressly/goose/v3"
)

func CheckAndMigrate() {
	migrateFile := "./00001_users_db.sql"
	db := InitDB()
	var exists bool
	err := db.QueryRow(`SELECT EXISTS (
        SELECT 1 FROM information_schema.tables WHERE table_name = 'users')`).Scan(&exists)
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
