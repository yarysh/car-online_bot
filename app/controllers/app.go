package controllers

import (
	"fmt"

	"github.com/revel/revel"
	"github.com/yarysh/car-online_bot/app"
	"github.com/yarysh/car-online_bot/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	var users []models.User

	_, err := app.DbMap.Select(&users, "select * from users")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)

	return c.RenderJSON(map[string]bool{"success": true})
}
