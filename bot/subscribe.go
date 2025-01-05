package bot

import (
	_ "embed"
	"fmt"
	"log"
	"ouro-telegram-app/tools"

	tele "gopkg.in/telebot.v4"
)

//go:embed templates/subscribe.md
var subscribeMessage string

func subscribe(bot *tele.Bot) {
	telegram := tele.InlineButton{
		Text: "Telegram",
		URL:  "https://t.me/ouroguru",
	}

	twitter := tele.InlineButton{
		Text: "X (Twitter)",
		URL:  "https://x.com/OuroGuru",
	}

	checkResult := tele.InlineButton{
		Text:   "Check Result",
		Unique: "check_result",
	}

	inlineKeys := [][]tele.InlineButton{
		{telegram, twitter},
		{checkResult},
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

	bot.Handle(&checkResult, func(c tele.Context) error {
		var (
			channelID      = ChannelID(tools.GetEnv("CHANNEL_ID", "@ouroguru"))
			telegramResult = "❌"
			twitterResult  = "❌"
		)

		if member, err := bot.ChatMemberOf(channelID, c.Sender()); err != nil {
			log.Println(err)
		} else if member.Role != tele.Left && member.Role != tele.Kicked {
			telegramResult = "✅"
		}

		return c.Send(fmt.Sprintf("Telegram subscription: %s\n", telegramResult) +
			fmt.Sprintf("Twitter subscription: %s\n", twitterResult))
	})
}
