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
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	defer func() {
		database.Disconnect()
	}()

	vendorRepository := repo.NewVendorRepo(db)
	vendorService := service.NewVendorService(vendorRepository)
	vendorController := controller.NewVendorController(vendorService)

	adminRepository := repo.NewAdminRepo(db)
	adminService := service.NewAdminService(adminRepository)
	adminController := controller.NewAdminController(adminService)

	serverHttp := server.NewServerHTTP(vendorController, adminController)

	return serverHttp, nil
}
