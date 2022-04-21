package content_test

import (
	"clean-hexa/business/content"
	"clean-hexa/business/content/spec"
	"errors"
	"os"
	"reflect"
	"testing"
)

var service content.Service
var content1, content2 content.Content
var insertSpec, updateSpec, failedSpec, errorSpec spec.UpsertContentSpec
var creator, updater string
var errorFindID int
var errorInsert error = errors.New("error on insert")
var errorFind error = errors.New("error on find")

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetContentByID(t *testing.T) {
	t.Run("Expect found the content", func(t *testing.T) {
		foundContent, _ := service.GetContentByID(content1.ID)
		if !reflect.DeepEqual(*foundContent, content1) {
			t.Error("Expect content has to be equal with content1", foundContent, content1)
		}
	})

	t.Run("Expect not found the content", func(t *testing.T) {
		content, err := service.GetContentByID(90)

		if err != nil {
			t.Error("Expect error is nil. Error: ", err)
		} else if content != nil {
			t.Error("Expect content must be not found (nil)")
		}
	})
}

func TestGetContentByTags(t *testing.T) {
	t.Run("Expect found the contents", func(t *testing.T) {
		contents, _ := service.GetContents()

		if len(contents) != 2 {
			t.Error("Expect content length must be two")
			t.FailNow()
		}

		if reflect.DeepEqual(contents[0], content1) {
			if !reflect.DeepEqual(contents[1], content2) {
				t.Error("Expect 2nd content is equal to content2")
			}
		} else if reflect.DeepEqual(contents[0], content2) {
			if !reflect.DeepEqual(contents[1], content1) {
				t.Error("Expect 2nd content is equal to content1")
			}
		} else {
			t.Error("Expect contents is content1 and content2")
		}
	})
}

func setup() {
	//initialize content1
	content1.ID = 1
	content1.Name = "Content one"
	content1.Description = "Description one"
	content1.Version = 1

	//initialize content 2
	content2.ID = 2
	content2.Name = "Content two"
	content2.Description = "Description two"
	content2.Version = 2

	repo := newInMemoryRepository()
	service = content.NewService(&repo)

	insertSpec.Name = "New Content"
	insertSpec.Description = "New Description"

	updateSpec.Name = "Update Content"
	updateSpec.Description = "Update Description"

	failedSpec.Name = ""
	failedSpec.Description = "Failed Description"

	errorSpec.Name = "Error Content"
	errorSpec.Description = "Error Description"

	creator = "creator"
	updater = "updater"

	errorFindID = 120
}

type inMemoryRepository struct {
	contentByID map[int]content.Content
	contents    []content.Content
}

func newInMemoryRepository() inMemoryRepository {
	var repo inMemoryRepository
	repo.contentByID = make(map[int]content.Content)
	repo.contents = []content.Content{}

	repo.contentByID[content1.ID] = content1
	repo.contentByID[content2.ID] = content2

	repo.contents = append(repo.contents, content1)
	repo.contents = append(repo.contents, content2)

	return repo
}

func (repo *inMemoryRepository) FindContentByID(ID int) (*content.Content, error) {
	if ID == errorFindID {
		return nil, errorFind
	}

	content, ok := repo.contentByID[ID]
	if !ok {
		return nil, nil
	}

	return &content, nil
}

func (repo *inMemoryRepository) FindAll() ([]content.Content, error) {
	var contents []content.Content
	contents = repo.contents

	return contents, nil
}
