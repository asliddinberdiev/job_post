package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID           primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title            string             `json:"title" bson:"title"`
	CompanyName      string             `json:"company_name" bson:"company_name"`
	Description      string             `json:"description" bson:"description"`
	Experience       float32            `json:"experience" bson:"experience"`
	JobType          string             `json:"job_type" bson:"job_type"`               // full, part
	EmploymentType   string             `json:"employment_type" bson:"employment_type"` // remote, onsite, hybrid
	Salary           Salary             `json:"salary" bson:"salary"`
	Location         Location           `json:"location" bson:"location"`
	Contact          Contact            `json:"contact" bson:"contact"`
	Status           string             `json:"status" bson:"status"` // active, draft
	Tags             []string           `json:"tags" bson:"tags"`
	Responsibilities []string           `json:"responsibilities" bson:"responsibilities"`
	Requirements     []string           `json:"requirements" bson:"requirements"`
	Benefits         []string           `json:"benefits" bson:"benefits"`
	Deadline         *time.Time         `json:"deadline,omitempty" bson:"deadline,omitempty"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
}

type Salary struct {
	Min      uint64 `json:"min" validate:"required,gte=0"`
	Max      uint64 `json:"max" validate:"required,gte=0"`
	Currency string `json:"currency" validate:"required,oneof=UZS RUB USD"`
}

type Location struct {
	Longitude float64 `json:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Address   string  `json:"address"`
}

type Contact struct {
	Phone    string `json:"phone" validate:"required,e164"`
	Telegram string `json:"telegram" validate:"omitempty,url"`
	LinkedIn string `json:"linkedin" validate:"omitempty,url"`
}

type CreatePostRequest struct {
	Title            string     `json:"title" validate:"required,min=3"`
	CompanyName      string     `json:"company_name" validate:"required,min=3"`
	Description      string     `json:"description" validate:"omitempty,min=10"`
	JobType          string     `json:"job_type" validate:"required,oneof=full part"`
	Experience       float32    `json:"experience" validate:"required,gte=0"`
	EmploymentType   string     `json:"employment_type" validate:"required,oneof=remote onsite hybrid"`
	Salary           Salary     `json:"salary" validate:"required"`
	Location         Location   `json:"location" validate:"required"`
	Contact          Contact    `json:"contact" validate:"required"`
	Tags             []string   `json:"tags" validate:"required,min=1,dive,required"`
	Responsibilities []string   `json:"responsibilities" validate:"omitempty,dive,required"`
	Requirements     []string   `json:"requirements" validate:"omitempty,dive,required"`
	Benefits         []string   `json:"benefits" validate:"omitempty,dive,required"`
	Deadline         *time.Time `json:"deadline,omitempty"`
}

type UpdatePostRequest struct {
	Title            *string    `json:"title" validate:"omitempty,min=3"`
	CompanyName      *string    `json:"company_name" validate:"omitempty,min=3"`
	Description      *string    `json:"description" validate:"omitempty,min=10"`
	JobType          *string    `json:"job_type" validate:"omitempty,oneof=full part"`
	Experience       *float32   `json:"experience" validate:"omitempty,gte=0"`
	EmploymentType   *string    `json:"employment_type" validate:"omitempty,oneof=remote onsite hybrid"`
	Salary           *Salary    `json:"salary" validate:"omitempty"`
	Location         *Location  `json:"location" validate:"omitempty"`
	Contact          *Contact   `json:"contact" validate:"omitempty"`
	Tags             []string   `json:"tags" validate:"omitempty,min=1,dive,required"`
	Responsibilities []string   `json:"responsibilities" validate:"omitempty,dive,required"`
	Requirements     []string   `json:"requirements" validate:"omitempty,dive,required"`
	Benefits         []string   `json:"benefits" validate:"omitempty,dive,required"`
	Deadline         *time.Time `json:"deadline,omitempty"`
}
