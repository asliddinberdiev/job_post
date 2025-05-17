package handler

import (
	"log/slog"

	"github.com/asliddinberdiev/job_post/internal/config"
	"github.com/asliddinberdiev/job_post/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	corsConfig := cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Authorization, Origin, Content-Type, Accept, Content-Language, Accept-Language, Access-Control-Allow-Headers",
	}

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(corsConfig))

	api := app.Group("/api")

	posts := api.Group("/posts")
	{
		posts.Post("/", h.CreatePost)
		posts.Get("/", h.GetPosts)
		posts.Get("/:id", h.GetPost)
		posts.Put("/:id", h.UpdatePost)
		posts.Delete("/:id", h.DeletePost)
	}

	return app
}
