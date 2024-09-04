package main

import (
	"stalcraftBot/cmd"
	"stalcraftBot/internal/tgBot"
)

func main() {

	cmd.Execute()
	tgBot.LoadChatID()

}
