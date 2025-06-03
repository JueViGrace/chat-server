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
	group.Post("/forgot/reset", handler.forgotPassReset)
	group.Post("/forgot/request", handler.forgotPassRequest)
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
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	tokens, err := h.db.LogIn(r)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)

	}

	res = types.RespondOk(tokens, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) signUp(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.SignUpRequest)

	err := types.ParseRequest(h.validator, c, r)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	tokens, err := h.db.SignUp(r)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)

	}

	res = types.RespondOk(tokens, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) refresh(c *fiber.Ctx, session *types.Session) error {
	res := new(types.APIResponse)

	tokens, err := h.db.Refresh(session)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)

	}

	res = types.RespondOk(tokens, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) forgotPassReset(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.RecoverPasswordRequest)

	err := types.ParseRequest(h.validator, c, r)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	msg, err := h.db.RecoverPassword(r)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(msg, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *authHandler) forgotPassRequest(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.RecoverPasswordRequest)

	err := types.ParseRequest(h.validator, c, r)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	msg, err := h.db.RecoverPassword(r)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondAccepted(msg, "Success")
	return c.Status(res.Status).JSON(res)
}
