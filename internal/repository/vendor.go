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

type Vendor struct {
	db *mongo.Collection
}

func NewVendorRepo(db *mongo.Database) *Vendor {
	return &Vendor{db: db.Collection("vendors")}
}

func (repo *Vendor) CreateVendor(vendor *models.Vendor) (*models.Vendor, error) {
	vendor.CreatedAt = time.Now()
	result, err := repo.db.InsertOne(context.TODO(), vendor)
	if err != nil {
		errStr := fmt.Sprintf("failed to create warehouse: %v", err)
		return nil, fmt.Errorf(errStr)
	}
	vendor.Id = result.InsertedID.(primitive.ObjectID)
	return vendor, nil
}

func (repo *Vendor) GetAllVendors() (*[]models.Vendor, error) {
	var vendors []models.Vendor
	cursor, err := repo.db.Find(context.TODO(), bson.M{})
	if err != nil {
		errStr := fmt.Sprintf("failed to retrieve vendor: %v", err)
		return nil, fmt.Errorf(errStr)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var vendor models.Vendor
		if err := cursor.Decode(&vendor); err != nil {
			errStr := fmt.Sprintf("failed to decode vendor: %v", err)
			return nil, fmt.Errorf(errStr)
		}
		vendors = append(vendors, vendor)
	}
	if err := cursor.Err(); err != nil {
		errStr := fmt.Sprintf("cursor error: %v", err)
		return nil, fmt.Errorf(errStr)
	}
	return &vendors, nil
}
