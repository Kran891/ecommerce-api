package repository

import (
	"ecommerce-api/config"

	"github.com/google/uuid"
)

type BaseRepository[T any] interface {
	Create(item *T) error
	FindByID(id uuid.UUID, predloads ...string) (*T, error)
	FindAll() ([]T, error)
	Update(item *T) (err error)
	Delete(id uuid.UUID) error
}

type GormRepository[T any] struct{}

func (r *GormRepository[T]) Create(item *T) error {
	return config.DB.Create(item).Error
}

func (r *GormRepository[T]) FindByID(id uuid.UUID, preloads ...string) (*T, error) {
	var item T
	query := config.DB // start from base DB instance

	for _, preload := range preloads {
		query = query.Preload(preload) // chain and assign
	}

	err := query.First(&item, "id = ?", id).Error
	return &item, err
}

func (r *GormRepository[T]) FindAll() ([]T, error) {
	var items []T
	err := config.DB.Find(&items).Error
	return items, err
}

func (r *GormRepository[T]) Update(item *T) (err error) {
	err = config.DB.Save(&item).Error
	return
}

func (r *GormRepository[T]) Delete(id uuid.UUID) error {
	var item T
	return config.DB.Delete(&item, "id =?", id).Error
}
