package repository

import (
	"context"
	"time"

	"github.com/asliddinberdiev/job_post/internal/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostsRepo struct {
	db *mongo.Collection
}

func NewPostsRepo(db *mongo.Database) *PostsRepo {
	return &PostsRepo{
		db: db.Collection(postsCollection),
	}
}

func (p *PostsRepo) CreatePost(ctx context.Context, req *models.Post) (primitive.ObjectID, error) {
	result, err := p.db.InsertOne(ctx, req)
	if err != nil {
		return primitive.NilObjectID, errors.Wrap(err, "failed to insert post")
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (p *PostsRepo) GetPost(ctx context.Context, id primitive.ObjectID) (*models.Post, error) {
	var post models.Post
	err := p.db.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&post)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find post")
	}

	return &post, nil
}

func (p *PostsRepo) GetPosts(ctx context.Context, limit, skip int64) ([]models.Post, int64, error) {
	var posts = make([]models.Post, 0)

	total, err := p.db.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to count posts")
	}
	cursor, err := p.db.Find(ctx, bson.D{}, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	})
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to find posts")
	}

	if err := cursor.All(ctx, &posts); err != nil {
		return nil, 0, errors.Wrap(err, "failed to decode posts")
	}

	return posts, total, nil
}

func (p *PostsRepo) UpdatePost(ctx context.Context, id primitive.ObjectID, req *models.UpdatePostRequest) error {
	update := bson.M{}
	now := time.Now()

	if req.Title != nil {
		update["title"] = *req.Title
	}
	if req.CompanyName != nil {
		update["company_name"] = *req.CompanyName
	}
	if req.Description != nil {
		update["description"] = *req.Description
	}
	if req.Experience != nil {
		update["experience"] = *req.Experience
	}
	if req.JobType != nil {
		update["job_type"] = *req.JobType
	}
	if req.EmploymentType != nil {
		update["employment_type"] = *req.EmploymentType
	}
	if req.Salary != nil {
		update["salary"] = req.Salary
	}
	if req.Location != nil {
		update["location"] = req.Location
	}
	if req.Contact != nil {
		update["contact"] = req.Contact
	}
	if req.Tags != nil {
		update["tags"] = req.Tags
	}
	if req.Responsibilities != nil {
		update["responsibilities"] = req.Responsibilities
	}
	if req.Requirements != nil {
		update["requirements"] = req.Requirements
	}
	if req.Benefits != nil {
		update["benefits"] = req.Benefits
	}
	if req.Deadline != nil {
		update["deadline"] = req.Deadline
	}

	update["updated_at"] = now

	if len(update) == 1 {
		return nil
	}

	result, err := p.db.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	if err != nil {
		return errors.Wrap(err, "failed to update post")
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (p *PostsRepo) DeletePost(ctx context.Context, id primitive.ObjectID) error {
	result, err := p.db.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return errors.Wrap(err, "failed to delete post")
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
