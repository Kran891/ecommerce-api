package services

import (
	"ecommerce-api/repository"

	"github.com/google/uuid"
)

type BaseService[T any] interface {
	Create(item *T) error
	FindByID(id uuid.UUID, preloads ...string) (*T, error)
	FindAll() ([]T, error)
	Update(item *T) (err error)
	Delete(id uuid.UUID) error
}

type BaseServiceImpl[T any] struct {
	repo repository.BaseRepository[T]
}

func (s *BaseServiceImpl[T]) Create(item *T) error {
	return s.repo.Create(item)
}
func (s *BaseServiceImpl[T]) FindByID(id uuid.UUID, preloads ...string) (*T, error) {
	return s.repo.FindByID(id, preloads...)
}
func (s *BaseServiceImpl[T]) FindAll() ([]T, error) {
	return s.repo.FindAll()
}
func (s *BaseServiceImpl[T]) Update(item *T) (err error) {
	return s.repo.Update(item)
}
func (s *BaseServiceImpl[T]) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func NewBaservice[T any](repo repository.BaseRepository[T]) BaseService[T] {
	return &BaseServiceImpl[T]{repo: repo}
}
