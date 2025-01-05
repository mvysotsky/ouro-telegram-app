package bot

import (
	_ "embed"
	"log"
	"ouro-telegram-app/tools"
	"time"

	tele "gopkg.in/telebot.v4"
)

//go:embed templates/welcome.md
var welcomeMessage string

//go:embed templates/subscribe.md
var subscribeMessage string

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

	for _, cmd := range []string{"/start", "/START"} {
		bot.Handle(cmd, func(c tele.Context) error {
			return c.Send(subscribeMessage, &tele.SendOptions{ParseMode: tele.ModeMarkdown})
		})
	}

	log.Println("Bot started")
	bot.Start()
}
