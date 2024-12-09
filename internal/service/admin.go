package service

import (
	"exclusive-sample/internal/models"
	repo "exclusive-sample/internal/repository"
)

type Admin struct {
	repo *repo.Admin
}

func NewAdminService(repo *repo.Admin) *Admin {
	return &Admin{repo: repo}
}

func (s *Admin) CreateAdmin(admin *models.Admin) (*models.Admin, error) {
	return s.repo.CreateAdmin(admin)
}

func (s *Admin) GetAllAdmin() (*[]models.Admin, error) {
	return s.repo.GetAllAdmin()
}
