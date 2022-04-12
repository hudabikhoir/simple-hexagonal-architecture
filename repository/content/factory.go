package content

import (
	"clean-hexa/business/content"
	"clean-hexa/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) content.Repository {
	var contentRepo content.Repository

	if dbCon.Driver == util.Postgres {
		// existing tetep jalan
		contentRepo = NewPostgresRepository(dbCon.Postgres)
	} else {
		// develop teknologi ini
		contentRepo = NewStaticRepository()
	}

	return contentRepo
}
