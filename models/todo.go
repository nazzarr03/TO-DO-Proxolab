package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Description string             `json:"description"`
	Completed   bool               `json:"completed"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
