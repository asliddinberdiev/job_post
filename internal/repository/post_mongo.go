package repository

import "go.mongodb.org/mongo-driver/mongo"

type PostsRepo struct {
	db *mongo.Collection
}

func NewPostsRepo(db *mongo.Database) *PostsRepo {
	return &PostsRepo{
		db: db.Collection(postsCollection),
	}
}
