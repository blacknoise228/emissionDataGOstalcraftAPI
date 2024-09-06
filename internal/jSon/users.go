package jSon

type User struct {
	ID     int    `json:"id"`
	UserID int64  `json:"userid"`
	Name   string `json:"username"`
}

var Users []User
