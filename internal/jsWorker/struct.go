package jsWorker

type EmissionInfo struct {
	CurrentStart  string `json:"currentStart"`  // last emission time
	PreviousStart string `json:"previousStart"` // preview emission time
	PreviousEnd   string `json:"previousEnd"`   // preview emission end
	Status        int    `json:"status"`        // status normal = 0, if status = 401, recreate token auth
}

type User struct {
	ID      int    `json:"id"`       // ID in the database
	Blocked bool   `json:"blocked"`  // If user in blacklist is true
	UserID  int64  `json:"userid"`   // UserID in telegram
	Name    string `json:"username"` // UserName in telegram
}

var Users []User
