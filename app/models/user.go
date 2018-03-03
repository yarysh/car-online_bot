package models

type User struct {
	Id        int
	Username  string
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	ChatId    int    `db:"chat_id"`
	ApiKey    string `db:"api_key"`
}
