package connection

import (
	server "exclusive-sample/internal/api"
	"exclusive-sample/internal/api/controller"
	"exclusive-sample/internal/config"
	"exclusive-sample/internal/database"
	repo "exclusive-sample/internal/repository"
	"exclusive-sample/internal/service"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {
	database, err := database.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	vendorRepository := repo.NewVendorRepo(database)
	vendorService := service.NewVendorService(vendorRepository)
	vendorController := controller.NewVendorController(vendorService)

	adminRepository := repo.NewAdminRepo(database)
	adminService := service.NewAdminService(adminRepository)
	adminController := controller.NewAdminController(adminService)

	serverHttp := server.NewServerHTTP(vendorController, adminController)

	return serverHttp, nil
}
