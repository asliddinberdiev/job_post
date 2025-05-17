package repository

import (
	"context"

	"github.com/asliddinberdiev/job_post/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Posts interface {
	CreatePost(ctx context.Context, req *models.Post) (primitive.ObjectID, error)
	GetPost(ctx context.Context, id primitive.ObjectID) (*models.Post, error)
	GetPosts(ctx context.Context, limit, skip int64) ([]models.Post, int64, error)
	UpdatePost(ctx context.Context, id primitive.ObjectID, req *models.UpdatePostRequest) error
	DeletePost(ctx context.Context, id primitive.ObjectID) error
}

type Repositories struct {
	Posts Posts
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Posts: NewPostsRepo(db),
	}
}
