package modules

import (
	"clean-hexa/api"
	contentV1Controller "clean-hexa/api/v1/content"
	contentV2Controller "clean-hexa/api/v2/content"
	contentService "clean-hexa/business/content"
	"clean-hexa/config"
	contentRepo "clean-hexa/repository/content"

	authController "clean-hexa/api/v1/auth"
	authService "clean-hexa/business/auth"
	"clean-hexa/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	contentPermitRepository := contentRepo.RepositoryFactory(dbCon)
	contentPermitService := contentService.NewService(contentPermitRepository)

	contentV1PermitController := contentV1Controller.NewController(contentPermitService)

	contentV2PermitController := contentV2Controller.NewController(contentPermitService)

	authPermitService := authService.NewService(config)
	authPermitController := authController.NewController(authPermitService)

	controllers := api.Controller{
		ContentV1Controller: contentV1PermitController,
		ContentV2Controller: contentV2PermitController,
		AuthController:      authPermitController,
	}

	return controllers
}
