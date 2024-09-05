package main

import (
	"stalcraftBot/cmd"
	"stalcraftBot/internal/logs"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	logs.StartLogger()
	cmd.Execute()

	wg.Wait()
}
