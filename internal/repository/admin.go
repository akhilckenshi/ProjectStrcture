package repo

import (
	"context"
	"exclusive-sample/internal/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Admin struct {
	db *mongo.Collection
}

func NewAdminRepo(db *mongo.Database) *Admin {
	return &Admin{db: db.Collection("admin")}
}

func (repo *Admin) CreateAdmin(admin *models.Admin) (*models.Admin, error) {
	admin.CreatedAt = time.Now()
	result, err := repo.db.InsertOne(context.TODO(), admin)
	if err != nil {
		errStr := fmt.Sprintf("failed to create warehouse: %v", err)
		return nil, fmt.Errorf(errStr)
	}
	admin.Id = result.InsertedID.(primitive.ObjectID)
	return admin, nil
}

func (repo *Admin) GetAllAdmin() (*[]models.Admin, error) {
	var admins []models.Admin
	cursor, err := repo.db.Find(context.TODO(), bson.M{})
	if err != nil {
		errStr := fmt.Sprintf("failed to retrieve Admin: %v", err)
		return nil, fmt.Errorf(errStr)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var admin models.Admin
		if err := cursor.Decode(&admin); err != nil {
			errStr := fmt.Sprintf("failed to decode Admin: %v", err)
			return nil, fmt.Errorf(errStr)
		}
		admins = append(admins, admin)
	}
	if err := cursor.Err(); err != nil {
		errStr := fmt.Sprintf("cursor error: %v", err)
		return nil, fmt.Errorf(errStr)
	}
	return &admins, nil
}
