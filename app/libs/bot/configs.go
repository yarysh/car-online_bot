package bot

// Outgoing message https://core.telegram.org/bots/api#sendmessage
type Message struct {
	ChatId                int64       `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWepPagePreview bool        `json:"disable_wep_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId      int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
}

type InlineKeyboard struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	Url          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

type ReplyKeyboard struct {
	Keyboard        [][]ReplyKeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                    `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool                    `json:"one_time_keyboard,omitempty"`
	Selective       bool                    `json:"selective,omitempty"`
}

type ReplyKeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact,omitempty"`
	RequestLocation bool   `json:"request_location,omitempty"`
}

// Partial representation of an incoming update Message object: https://core.telegram.org/bots/api#message
type updateMessage struct {
	Id       int64
	From     from
	Chat     chat
	Text     string
	Entities []messageEntity
}

// Representation of a User object: https://core.telegram.org/bots/api#user
type from struct {
	Id           int64
	IsBot        bool
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
}

// Partial representation of a Chat object: https://core.telegram.org/bots/api#chat
type chat struct {
	Id        int64
	Type      string
	Username  string
	FirstName string
	LastName  string
}

// Partial representation of a MessageEntity object: https://core.telegram.org/bots/api#messageentity
type messageEntity struct {
	Type   string
	Offset int
	Length int
	Url    string
}
