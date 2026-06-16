package repository

import (
	"ewallet-ums/internal/domain/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) Save(user *user.Entity) error {
	return repo.DB.Create(&user).Error
}

func (repo *UserRepository) FindByUsername(username string) (*user.Entity, error) {
	var (
		entity user.Entity
		err    error
	)

	err = repo.DB.Where("username = ?", username).First(&entity).Error

	return &entity, err
}
