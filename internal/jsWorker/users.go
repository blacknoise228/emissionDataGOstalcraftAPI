package jsWorker

type User struct {
	ID      int    `json:"id"`
	Blocked bool   `json:"blocked"`
	UserID  int64  `json:"userid"`
	Name    string `json:"username"`
}

var Users []User
