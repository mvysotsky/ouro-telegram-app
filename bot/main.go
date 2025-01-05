package bot

import (
	"log"
	"ouro-telegram-app/tools"
	"time"

	tele "gopkg.in/telebot.v4"
)

func Start() {
	var (
		bot *tele.Bot
		err error
	)

	pref := tele.Settings{
		Token:  tools.GetEnv("BOT_TOKEN", ""),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	if bot, err = tele.NewBot(pref); err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello Ruslan!")
	})

	log.Println("Bot started")
	bot.Start()
}
