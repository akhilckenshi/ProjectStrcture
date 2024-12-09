package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name" validate:"gte=3"`
	Email        string             `json:"email" bson:"email" validate:"email"`
	Password     string             `json:"password" bson:"password" validate:"min=6,max=20"`
	Picture      string             `json:"picture" bson:"picture"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	IsSuperAdmin bool               `json:"isSuperAdmin" bson:"isSuperAdmin"`
}
