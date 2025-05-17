package handler

import (
	"errors"

	"github.com/asliddinberdiev/job_post/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "something went wrong"

	var fe *fiber.Error
	if errors.As(err, &fe) {
		code = fe.Code
		message = fe.Message
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		code = fiber.StatusNotFound
		message = "not found"
	}

	if h.cfg.App.Environment != "prod" {
		message = err.Error()
	}

	h.log.Error("error", "error", err)

	return ctx.Status(code).JSON(&models.ResponseMessage{
		Code:    code,
		Message: message,
	})
}
