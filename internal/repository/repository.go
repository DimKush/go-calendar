package repository

type Event interface {
}

type Repository struct {
	Event
}

func NewService() *Repository {
	return &Repository{}
}
