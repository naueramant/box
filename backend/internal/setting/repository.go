package user

import "gorm.io/gorm"

type Repository struct{}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}
