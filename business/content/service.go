package content

// ingoing port
type Repository interface {
	FindContentByID(id int) (content *Content, err error)
	FindAll() (contents []Content, err error)
	InsertContent(content Content) (err error)
	UpdateContent(content Content, currentVersion int) (err error)
}

// outgoing port
type Service interface {
	GetContentByID(id int) (content *Content, err error)
	GetContents() (contents []Content, err error)
	CreateContent(content Content) (err error)
	UpdateContent(content Content, currentVersion int) (err error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
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

func (s *service) CreateContent(content Content) (err error) {
	return
}

func (s *service) UpdateContent(content Content, currentVersion int) (err error) {
	return
}
