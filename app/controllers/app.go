package controllers

import (
	"fmt"

	"github.com/revel/revel"
	"gopkg.in/telegram-bot-api.v4"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	bot, err := tgbotapi.NewBotAPI(revel.Config.StringDefault("app.bot_api_key", ""))
	if err != nil {
		fmt.Println(err)
	}

	msg := tgbotapi.NewMessage(int64(revel.Config.IntDefault("app.bot_chat_id", 0)), revel.AppName)
	_, sendErr := bot.Send(msg)
	return c.RenderJSON(map[string]bool{"success": sendErr == nil})
}
