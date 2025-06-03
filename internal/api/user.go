package api

import (
	"gps-tracker/internal/data"
	"gps-tracker/internal/types"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type userHandler struct {
	db        data.UserStore
	validator *types.XValidator
}

func (a *api) userRoutes(api fiber.Router) {
	group := api.Group("/user")
	handler := newUserHandler(a.db.UserStore(), a.validator)

	group.Get("/check", handler.CheckUserFields)
	group.Get("/:id", a.sessionMiddleware, handler.GetUser)
	group.Get("/me", a.authenticatedHandler(handler.GetMyself))
	group.Patch("/me", a.authenticatedHandler(handler.UpdateUser))
	group.Delete("/me", a.authenticatedHandler(handler.DeleteUser))
}

func newUserHandler(db data.UserStore, validator *types.XValidator) *userHandler {
	return &userHandler{
		db:        db,
		validator: validator,
	}
}

func (h *userHandler) CheckUserFields(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	fields := new(types.CheckUserFieldsRequest)

	email := c.Query("email", "")
	username := c.Query("username", "")
	phone := c.Query("phone", "")

	if email != "" {
		usedEmail, err := h.db.CheckUsedEmail(email)
		if err != nil {
			res = types.RespondBadRequest(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		}
		fields.UsedEmail = usedEmail
	}

	if username != "" {
		usedUsername, err := h.db.CheckUsedUsername(username)
		if err != nil {
			res = types.RespondBadRequest(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		}
		fields.UsedUsername = usedUsername
	}

	if phone != "" {
		usedPhone, err := h.db.CheckUsedPhoneNumber(phone)
		if err != nil {
			res = types.RespondBadRequest(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		}
		fields.UsedPhoneNumber = usedPhone
	}

	res = types.RespondOk(fields, "Success")
	return c.Status(res.Status).JSON(res)

}

func (h *userHandler) GetUser(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	user := new(types.User)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	user, err = h.db.GetUser(id)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	user.PhoneNumber = ""
	user.Email = ""

	res = types.RespondOk(user, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *userHandler) GetMyself(c *fiber.Ctx, session *types.Session) error {
	res := new(types.APIResponse)
	user := new(types.User)

	user, err := h.db.GetUser(session.UserID)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(user, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *userHandler) UpdateUser(c *fiber.Ctx, session *types.Session) error {
	res := new(types.APIResponse)
	r := new(types.UpdateUser)

	err := types.ParseRequest(h.validator, c, r)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	err = h.db.UpdateUser(r)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk("Updated!", "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *userHandler) DeleteUser(c *fiber.Ctx, session *types.Session) error {
	res := new(types.APIResponse)

	err := h.db.DeleteUser(session.ID)
	if err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk("User deleted!", "Success")
	return c.Status(res.Status).JSON(res)
}
