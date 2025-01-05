package bot

import (
	_ "embed"

	tele "gopkg.in/telebot.v4"
)

//go:embed templates/subscribe.md
var subscribeMessage string

func subscribe(bot *tele.Bot) {
	telegram := tele.InlineButton{
		Text:   "Telegram",
		Unique: "telegram",
		URL:    "https://t.me/ouroguru",
	}

	twitter := tele.InlineButton{
		Text:   "X (Twitter)",
		Unique: "twitter",
		URL:    "https://x.com/OuroGuru",
	}

	inlineKeys := [][]tele.InlineButton{
		{telegram, twitter},
	}
	replyMarkup := &tele.ReplyMarkup{InlineKeyboard: inlineKeys}

	for _, cmd := range []string{"/start", "/START"} {
		bot.Handle(cmd, func(c tele.Context) error {
			return c.Send(subscribeMessage, &tele.SendOptions{
				ParseMode:   tele.ModeMarkdown,
				ReplyMarkup: replyMarkup,
			})
		})
	}
}
