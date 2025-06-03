package data

import (
	"context"
	"gps-tracker/internal/database"
)

type ChatStore interface {
}

func (s *storage) ChatStore() ChatStore {
	return NewChatStore(s.ctx, s.queries)
}

type chatStore struct {
	ctx     context.Context
	queries *database.Queries
}

func NewChatStore(ctx context.Context, queries *database.Queries) ChatStore {
	return &chatStore{
		ctx:     ctx,
		queries: queries,
	}
}
