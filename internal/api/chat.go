package api

import (
	"gps-tracker/internal/data"
	"gps-tracker/internal/types"

	"github.com/gofiber/fiber/v2"
)

type chatHandler struct {
	db        data.ChatStore
	validator *types.XValidator
}

func (a *api) chatRoutes(api fiber.Router) {
	group := api.Group("/chats")
	handler := newChatHandler(a.db.ChatStore(), a.validator)

	group.Get("/", a.authenticatedHandler(handler.getChats))
	group.Get("/:id", a.authenticatedHandler(handler.getChat))

}

func newChatHandler(db data.ChatStore, validator *types.XValidator) *chatHandler {
	return &chatHandler{
		db:        db,
		validator: validator,
	}
}

func (h *chatHandler) getChats(c *fiber.Ctx, data *types.AuthData) error {
	return nil
}

func (h *chatHandler) getChat(c *fiber.Ctx, data *types.AuthData) error {
	return nil
}
