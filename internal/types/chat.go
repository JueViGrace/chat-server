package types

import (
	"github.com/google/uuid"
)

type Chat struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Icon     string    `json:"icon"`
	Type     ChatType  `json:"type"`
	Messages []Message `json:"messages"`
	Members  []Member  `json:"members"`
}

type Member struct {
	User User     `json:"user"`
	Role ChatRole `json:"role"`
}

type Message struct {
	ID     uuid.UUID   `json:"id"`
	Type   MessageType `json:"type"`
	Text   string      `json:"text"`
	Time   string      `json:"time"`
	ChatID uuid.UUID   `json:"chat_id"`
	UserID uuid.UUID   `json:"user_id"`
}

type ChatRole int

const (
	ChatMember ChatRole = iota
	ChatAdmin
	ChatOwner
)

func (c ChatRole) String() string {
	switch c {
	case ChatOwner:
		return "ChatOwner"
	case ChatAdmin:
		return "ChatAdmin"
	case ChatMember:
		return "ChatMember"
	default:
		return "Unknown"
	}
}

type ChatType int

const (
	PrivateChat ChatType = iota
	GroupChat
)

func (c ChatType) String() string {
	switch c {
	case GroupChat:
		return "GroupChat"
	case PrivateChat:
		return "PrivateChat"
	default:
		return "Unknown"
	}
}

type MessageType int

const (
	Text MessageType = iota
	Image
	Sticker
)

func (c MessageType) String() string {
	switch c {
	case Text:
		return "Text"
	case Image:
		return "Image"
	case Sticker:
		return "Sticker"
	default:
		return "Unknown"
	}
}
