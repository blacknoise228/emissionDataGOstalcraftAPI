package postgres

import (
	"database/sql"
	"fmt"
	"stalcraftbot/internal/logs"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitDB() *sql.DB {
	connStr := viper.GetString("databaseurl")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logs.Logger.Fatal().Err(err).Msg("Open DB Error:")
	}
	if err := db.Ping(); err != nil {
		logs.Logger.Err(err).Msg("Connect DB Error:")
		return nil
	}
	logs.Logger.Info().Msg("DB Connected")
	return db
}

// Save user to db
func SaveChatID(db *sql.DB, user User) {
	_, err := db.Exec(
		"insert into users (userid, username) values ($1, $2)",
		user.UserID, user.Name)
	if err != nil {
		logs.Logger.Error().Err(err).Msg("Adding user to database error:")
	}
	logs.Logger.Debug().Msg("Adding user to database done")
}

// Load users from db
func LoadChatID(db *sql.DB) error {
	Users = nil
	rows, err := db.Query("select * from users")
	if err != nil {
		logs.Logger.Error().Err(err).Msg("Load chat id from db error")
		return err
	}
	defer rows.Close()
	logs.Logger.Debug().Msg("Open chat ids db done")

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.UserID)
		if err != nil {
			logs.Logger.Err(err).Msg("Load users from db error:")
			continue
		}
		Users = append(Users, user)
	}
	logs.Logger.Debug().Msg("Load chat ids from db done")
	return nil
}
func SearchID(db *sql.DB, num int64) bool {
	LoadChatID(db)
	for _, v := range Users {
		if v.UserID == num {
			logs.Logger.Info().Msg(fmt.Sprint("Request ID: ", v.UserID))
			return true
		}
	}
	return false
}
func DeleteUserInDB(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		logs.Logger.Err(err).Msg("Deleting user Error:")
		return
	}
	logs.Logger.Info().Msg("Deleting user done")
}
func QuantityUsers() int {
	var db = InitDB()
	LoadChatID(db)
	return len(Users)
}
