package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobPost struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title       string             `json:"title" bson:"title" validate:"required"`
	CompanyName string             `json:"company_name" bson:"company_name" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"omitempty,min=10"`
	JobType     string             `json:"job_type" bson:"job_type" validate:"required,oneof=full-time part-time"`
	Salary      Salary             `json:"salary" bson:"salary" validate:"required"`
	Location    Location           `json:"location" bson:"location" validate:"required"`
	Contact     Contact            `json:"contact" bson:"contact" validate:"required"`
	Status      string             `json:"status" bson:"status" validate:"oneof=active closed draft"`
	Views       uint16             `json:"views" bson:"views"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type Salary struct {
	Min      float64 `json:"min" bson:"min" validate:"required,number,gte=0"`
	Max      float64 `json:"max" bson:"max" validate:"required,number,gte=0"`
	Currency string  `json:"currency" bson:"currency" validate:"required,oneof=UZS RUB USD"`
}

type Location struct {
	Longitude float64 `json:"longitude" bson:"longitude" validate:"required,number"`
	Latitude  float64 `json:"latitude" bson:"latitude" validate:"required,number"`
	Address   string  `json:"address" bson:"address"`
}

type Contact struct {
	Phone    string `json:"phone" bson:"phone" validate:"required,e164"`
	Telegram string `json:"telegram" bson:"telegram" validate:"omitempty,url"`
	LinkedIn string `json:"linkedin" bson:"linkedin" validate:"omitempty,url"`
}
