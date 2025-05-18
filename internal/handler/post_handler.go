package handler

import (
	"time"

	"github.com/asliddinberdiev/job_post/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary		Create post
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		request body	models.CreatePostRequest	true	"Create post"
// @Success		201	{object} models.ResponseID
// @Failure		400	{object} models.ResponseMessage
// @Failure		500	{object} models.ResponseMessage
// @Router		/posts [post]
func (h *Handler) CreatePost(c *fiber.Ctx) error {
	var req models.CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	if err := h.valid.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	post := models.Post{
		ID:               primitive.NewObjectID(),
		Title:            req.Title,
		CompanyName:      req.CompanyName,
		Description:      req.Description,
		Experience:       req.Experience,
		JobType:          req.JobType,
		EmploymentType:   req.EmploymentType,
		Salary:           req.Salary,
		Location:         req.Location,
		Contact:          req.Contact,
		Status:           "active",
		Tags:             req.Tags,
		Responsibilities: req.Responsibilities,
		Requirements:     req.Requirements,
		Benefits:         req.Benefits,
		Deadline:         req.Deadline,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	id, err := h.repo.Posts.CreatePost(c.Context(), &post)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(&models.ResponseID{ID: id})
}

// @Summary		Get post
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		id	path	string	true	"Post ID"
// @Success		200	{object} models.ResponseData
// @Failure		400	{object} models.ResponseMessage
// @Failure		404	{object} models.ResponseMessage
// @Failure		500	{object} models.ResponseMessage
// @Router		/posts/{id} [get]
func (h *Handler) GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	jobPost, err := h.repo.Posts.GetPost(c.Context(), objID)
	if err != nil {
		return err
	}

	return c.JSON(&models.ResponseData{
		Data:    jobPost,
		Code:    fiber.StatusOK,
		Message: "success",
	})
}

// @Summary		Get posts
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		limit	query	int	false	"Limit"
// @Param		skip	query	int	false	"Skip"
// @Success		200	{object} models.ResponseList
// @Failure		400	{object} models.ResponseMessage
// @Failure		500	{object} models.ResponseMessage
// @Router		/posts [get]
func (h *Handler) GetPosts(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	skip := c.QueryInt("skip", 0)

	jobPosts, total, err := h.repo.Posts.GetPosts(c.Context(), int64(limit), int64(skip))
	if err != nil {
		return err
	}

	return c.JSON(&models.ResponseList{
		Data:    jobPosts,
		Total:   total,
		Code:    fiber.StatusOK,
		Message: "success",
	})
}

// @Summary		Update post
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		id	path	string	true	"Post ID"
// @Param		request	body	models.UpdatePostRequest	true	"Update post"
// @Success		200	{object} models.ResponseMessage
// @Failure		400	{object} models.ResponseMessage
// @Failure		404	{object} models.ResponseMessage
// @Failure		500	{object} models.ResponseMessage
// @Router		/posts/{id} [patch]
func (h *Handler) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	var req models.UpdatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	if err := h.valid.Struct(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.repo.Posts.UpdatePost(c.Context(), objID, &req); err != nil {
		return err
	}

	return c.JSON(&models.ResponseMessage{
		Code:    fiber.StatusOK,
		Message: "success",
	})
}

// @Summary		Delete post
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		id	path	string	true	"Post ID"
// @Success		200	{object} models.ResponseMessage
// @Failure		400	{object} models.ResponseMessage
// @Failure		404	{object} models.ResponseMessage
// @Failure		500	{object} models.ResponseMessage
// @Router		/posts/{id} [delete]
func (h *Handler) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	if err := h.repo.Posts.DeletePost(c.Context(), objID); err != nil {
		return err
	}

	return c.JSON(&models.ResponseMessage{
		Code:    fiber.StatusOK,
		Message: "success",
	})
}
