package content

import (
	"clean-hexa/business/content"
)

var ContentList = make(map[int]content.Content, 0)
var LastID int = 1

type StaticRepository struct {
	data   map[int]content.Content
	lastID int
}

func NewStaticRepository() *StaticRepository {
	return &StaticRepository{
		data:   ContentList,
		lastID: LastID,
	}
}

func (repo *StaticRepository) FindContentByID(ID int) (content *content.Content, err error) {
	return
}

func (repo *StaticRepository) FindAll() (contents []content.Content, err error) {
	tempData := content.Content{
		ID:          1,
		Name:        "test",
		Description: "test",
		Version:     1,
	}

	repo.data[1] = tempData
	result := repo.data

	for _, value := range result {
		var tempContent content.Content
		tempContent.ID = value.ID
		tempContent.Description = value.Description
		tempContent.Name = value.Name
		tempContent.Version = value.Version

		contents = append(contents, tempContent)
	}
	return
}

func (repo *StaticRepository) InsertContent(content content.Content) (err error) {
	return
}

func (repo *StaticRepository) UpdateContent(content content.Content, currentVersion int) (err error) {
	return
}
