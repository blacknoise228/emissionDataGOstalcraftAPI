package jsWorker

type EmissionInfo struct {
	CurrentStart  string `json:"currentStart"`  // last emission time
	PreviousStart string `json:"previousStart"` // preview emission time
	PreviousEnd   string `json:"previousEnd"`   // preview emission end
	Status        int    `json:"status"`        // status normal = 0, if status = 401, recreate token auth
}
