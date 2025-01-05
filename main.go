package main

import (
	"ouro-telegram-app/bot"
	"ouro-telegram-app/tools"
)

func main() {
	tools.LoadEnv()
	bot.Start()
}
