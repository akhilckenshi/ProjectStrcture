package service

import (
	"exclusive-sample/internal/models"
	repo "exclusive-sample/internal/repository"
)

type Vendor struct {
	repo *repo.Vendor
}

func NewVendorService(repo *repo.Vendor) *Vendor {
	return &Vendor{repo: repo}
}

func (s *Vendor) CreateVendor(vendor *models.Vendor) (*models.Vendor, error) {
	return s.repo.CreateVendor(vendor)
}

func (s *Vendor) GetAllVendors() (*[]models.Vendor, error) {
	return s.repo.GetAllVendors()
}
