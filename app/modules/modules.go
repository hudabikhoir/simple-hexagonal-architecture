package modules

import (
	"clean-hexa/api"
	contentV1Controller "clean-hexa/api/v1/content"
	contentV2Controller "clean-hexa/api/v2/content"
	contentService "clean-hexa/business/content"
	contentRepo "clean-hexa/repository/content"
	"clean-hexa/util"
)

func RegisterModules(dbCon *util.DatabaseConnection) api.Controller {
	contentPermitRepository := contentRepo.RepositoryFactory(dbCon)
	contentPermitService := contentService.NewService(contentPermitRepository)

	contentV1PermitController := contentV1Controller.NewController(contentPermitService)

	contentV2PermitController := contentV2Controller.NewController(contentPermitService)

	controllers := api.Controller{
		ContentV1Controller: contentV1PermitController,
		ContentV2Controller: contentV2PermitController,
	}

	return controllers
}
