package api

import (
	"gps-tracker/internal/data"
	"gps-tracker/internal/types"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	db        data.AuthStore
	validator *types.XValidator
}

func (a *api) authRoutes(api fiber.Router) {
	group := api.Group("/auth")
	handler := newAuthHandler(a.db.AuthStore(), a.validator)

	group.Post("/login", handler.login)
	group.Post("/signup", handler.signUp)
	group.Post("/forgot", handler.forgotPass)
	group.Post("/refresh", a.authenticatedHandler(handler.refresh))

}

func newAuthHandler(db data.AuthStore, validator *types.XValidator) *authHandler {
	return &authHandler{
		db:        db,
		validator: validator,
	}
}

func (h *authHandler) login(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.SignInRequest)

	err := types.ParseRequest(h.validator, c, r)

	session, err := h.db.CreateSession()
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)

	}

	res = types.RespondOk(session, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) signUp(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	session, err := h.db.CreateSession()
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)

	}

	res = types.RespondOk(session, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) refresh(c *fiber.Ctx, data *types.AuthData) error {
	res := new(types.APIResponse)

	session, err := h.db.Refresh(data.SessionID)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)

	}

	res = types.RespondOk(session, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) forgotPass(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	res = types.RespondNotFound("end point not implemented", "")
	return c.Status(res.Status).JSON(res)
}
