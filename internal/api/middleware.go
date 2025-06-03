package api

import (
	"gps-tracker/internal/data"
	"gps-tracker/internal/types"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (a *api) adminAuthMiddleware(c *fiber.Ctx) error {
	session, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	if session.Role != types.SystemAdmin {
		res := types.RespondForbbiden(nil, "forbbiden resource")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) sessionMiddleware(c *fiber.Ctx) error {
	_, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) authenticatedHandler(handler types.AuthDataHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := getUserDataForReq(c, a.db)
		if err != nil {
			res := types.RespondUnauthorized(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		}

		return handler(c, data)
	}
}

func getUserDataForReq(c *fiber.Ctx, db data.Storage) (*types.Session, error) {
	jwt, err := types.ExtractJWTFromHeader(c, func(s string, err error) {
		log.Errorf("Error: %v", err)
		db.AuthStore().DeleteSessionByToken(s)
	})
	if err != nil {
		return nil, err
	}

	session, err := db.AuthStore().GetSessionById(jwt.Claims.SessionID)
	if err != nil {
		return nil, err
	}

	return session, nil
}
