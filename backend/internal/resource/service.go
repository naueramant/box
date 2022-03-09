package resource

import "gorm.io/gorm"

type Service struct{}

type ServiceImpl struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}
