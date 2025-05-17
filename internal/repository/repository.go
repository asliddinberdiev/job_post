package repository

import "go.mongodb.org/mongo-driver/mongo"

type Posts interface {
}

type Repositories struct {
	Posts Posts
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Posts: NewPostsRepo(db),
	}
}
