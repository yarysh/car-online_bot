package models

import (
	"database/sql"

	"github.com/yarysh/car-online_bot/app"
)

type User struct {
	Id        int
	Username  string
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	ChatId    int64  `db:"chat_id"`
	ApiKey    string `db:"api_key"`
}

func (u User) UpdateApiKey() (sql.Result, error) {
	return app.DbMap.Exec(
		"INSERT INTO users (username, first_name, last_name, chat_id, api_key) VALUES (:username, :firstName, :lastName, :chatId, :apiKey) "+
			"ON CONFLICT (chat_id) DO UPDATE SET "+
			"(username, first_name, last_name, api_key) = (:username, :firstName, :lastName, :apiKey)",
		map[string]interface{}{"username": u.Username, "firstName": u.FirstName, "lastName": u.LastName, "chatId": u.ChatId, "apiKey": u.ApiKey},
	)
}
