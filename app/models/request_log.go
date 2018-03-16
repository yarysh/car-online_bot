package models

import (
	"database/sql"

	"github.com/revel/revel"
	"github.com/yarysh/car-online_bot/app"
)

type RequestLog struct {
	Id       int
	Url      string
	Response string
	Error    string
}

func (rl RequestLog) Save() (sql.Result, error) {
	res, err := app.DbMap.Exec(
		"INSERT1 INTO request_logs (url, response, error) VALUES (:url, :response, :error) ",
		map[string]interface{}{"url": rl.Url, "response": rl.Response, "error": rl.Error},
	)
	if err != nil {
		revel.AppLog.Error(err.Error())
	}
	return res, err
}
