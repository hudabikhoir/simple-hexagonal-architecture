package content

import (
	"clean-hexa/business/content/spec"
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

// ingoing port
type Repository interface {
	FindContentByID(id int) (content *Content, err error)
	FindAll() (contents []Content, err error)
}

// outgoing port
type Service interface {
	GetContentByID(id int) (content *Content, err error)
	GetContents() (contents []Content, err error)
	CreateContent(spec spec.UpsertContentSpec) (err error)
	UpdateContent(content Content, currentVersion int) (err error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) GetContentByID(id int) (content *Content, err error) {
	result, err := s.repository.FindContentByID(id)
	return result, err
}

func (s *service) GetContents() (contents []Content, err error) {
	contents, err = s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (s *service) CreateContent(spec spec.UpsertContentSpec) (err error) {
	err = s.validate.Struct(spec)
	fmt.Println("err service: ", err)

	return
}

func (s *service) UpdateContent(content Content, currentVersion int) (err error) {
	return
}
