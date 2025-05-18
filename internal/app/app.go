package app

import (
	"context"
	"log/slog"

	"github.com/asliddinberdiev/job_post/internal/config"
	"github.com/asliddinberdiev/job_post/internal/handler"
	"github.com/asliddinberdiev/job_post/internal/repository"
	"github.com/asliddinberdiev/job_post/pkg/db"
	"github.com/go-playground/validator/v10"
)

func Run() {
	logger := slog.Default()

	cfg, err := config.Init()
	if err != nil {
		logger.Error("failed to load configs", "error", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoClient, err := db.NewMongoDBClient(ctx, cfg.GetMongoURL(), cfg.MongoDB.Username, cfg.MongoDB.Password)
	if err != nil {
		logger.Error("failed to connect to mongo", "error", err)
		return
	}
	defer mongoClient.Disconnect(ctx)

	db := mongoClient.Database(cfg.MongoDB.Database)
	repo := repository.NewRepositories(db)
	valid := validator.New()

	handler := handler.NewHandler(logger, cfg, valid, repo)

	app := handler.CreateApp()

	logger.Info("Server started", "address", cfg.GetAppAddress())
	if err := app.Listen(cfg.GetAppAddress()); err != nil {
		logger.Error("failed to start server", "error", err)
	}
}
