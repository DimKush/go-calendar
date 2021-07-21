package service

import "github.com/DimKush/go-calendar/internal/repository"

type Event interface {
}

type Service struct {
	Event
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
