package api

import (
	"context"
	"fmt"
	"gps-tracker/internal/data"
	"gps-tracker/internal/types"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

type Api interface {
	Start() error
	ShutDownWithContext(ctx context.Context) error
}

type api struct {
	*fiber.App
	db        data.Storage
	validator *types.XValidator
}

func New() Api {
	return &api{
		App: fiber.New(fiber.Config{
			ServerHeader: "gps-tracker-server",
			AppName:      "gps-tracker-server",
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				res := types.RespondInternalServerError(nil, err.Error())
				return c.Status(res.Status).JSON(res)
			},
		}),
		db: data.NewStorage(),
		validator: &types.XValidator{
			Validator: validator.New(),
		},
	}
}

func (a *api) Start() error {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return err
	}

	a.App.Use(logger.New())
	a.App.Use(cors.New())

	a.registerRoutes()

	// this should go after registering app routes
	a.App.Use(func(c *fiber.Ctx) error {
		res := types.RespondNotFound(nil, "Route not found")
		return c.Status(res.Status).JSON(res)
	})

	err = a.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	return nil
}

func (a *api) ShutDownWithContext(ctx context.Context) error {
	return a.ShutDownWithContext(ctx)
}
