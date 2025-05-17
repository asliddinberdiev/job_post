package handler

import (
	"errors"
	"log/slog"

	"github.com/asliddinberdiev/job_post/internal/config"
	"github.com/asliddinberdiev/job_post/internal/models"
	"github.com/asliddinberdiev/job_post/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Handler struct {
	log   *slog.Logger
	cfg   *config.Config
	repo  *repository.Repositories
	valid *validator.Validate
}

func NewHandler(log *slog.Logger, cfg *config.Config, valid *validator.Validate, repo *repository.Repositories) *Handler {
	return &Handler{
		log:   log,
		cfg:   cfg,
		valid: valid,
		repo:  repo,
	}
}

func (h *Handler) CreateApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: h.errorHandler,
	})

	app.Use(recover.New())

	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("started")
	})

	return app
}

func (h *Handler) errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "something went wrong"

	var fe *fiber.Error
	if errors.As(err, &fe) {
		code = fe.Code
		message = fe.Message
	}

	if h.cfg.App.Environment != "prod" {
		message = err.Error()
	}

	h.log.Error("error", "error", err)

	return ctx.Status(code).JSON(&models.MessageResponse{
		Code:    code,
		Message: message,
	})
}
