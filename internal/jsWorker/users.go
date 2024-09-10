package jsWorker

type User struct {
	ID      int    `json:"id"`       // ID in the database
	Blocked bool   `json:"blocked"`  // If user in blacklist is true
	UserID  int64  `json:"userid"`   // UserID in telegram
	Name    string `json:"username"` // UserName in telegram
}

var Users []User
