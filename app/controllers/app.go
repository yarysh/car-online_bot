package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.RenderJSON(map[string]string{"app": "car-online_bot"})
}
