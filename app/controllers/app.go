package controllers

import (
	"github.com/revel/revel"
	"github.com/yarysh/car-online_bot/app/libs/bot"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var update map[string]interface{}
	c.Params.BindJSON(&update)
	bot.ProcessUpdate(update)
	return c.RenderText("")
}
