package types

import (
	"github.com/google/uuid"
)

type Chat struct {
	ID       uuid.UUID    `json:"id"`
	Name     string       `json:"name"`
	Icon     string       `json:"icon"`
	Type     ChatType     `json:"type"`
	Messages []Message    `json:"messages"`
	Members  []ChatMember `json:"members"`
}

type ChatMember struct {
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
	Owner ChatRole = iota
	Admin
	Member
)

type ChatType int

const (
	Group ChatType = iota
	Private
)

type MessageType int

const (
	Text MessageType = iota
	Image
	Sticker
)
