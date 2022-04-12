package content

import (
	"clean-hexa/business/content"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (repo *PostgresRepository) FindContentByID(id int) (content *content.Content, err error) {
	return content, nil
}

func (repo *PostgresRepository) FindAll() (contents []content.Content, err error) {
	result := repo.db.Find(&contents)
	if result.Error != nil {
		return nil, result.Error
	}
	return contents, nil
}

func (repo *PostgresRepository) InsertContent(content content.Content) (err error) {
	return
}

func (repo *PostgresRepository) UpdateContent(content content.Content, currentVersion int) (err error) {
	return
}
