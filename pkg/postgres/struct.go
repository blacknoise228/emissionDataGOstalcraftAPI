package postgres

type User struct {
	ID     int    `db:"id"`       // ID in the database
	UserID int64  `db:"userid"`   // UserID in telegram
	Name   string `db:"username"` // UserName in telegram
}

var Users []User
