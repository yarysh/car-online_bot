package bot

// Partial representation of a Chat object: https://core.telegram.org/bots/api#chat
type chat struct {
	Id        int64
	Type      string
	Username  string
	FirstName string
	LastName  string
}

// Partial representation of a Message object: https://core.telegram.org/bots/api#message
type message struct {
	Id       int64
	From     user
	Chat     chat
	Text     string
	Entities []messageEntity
}

// Partial representation of a MessageEntity object: https://core.telegram.org/bots/api#messageentity
type messageEntity struct {
	Type   string
	Offset int
	Length int
	Url    string
}

// Representation of a User object: https://core.telegram.org/bots/api#user
type user struct {
	Id           int64
	IsBot        bool
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
}
