package repository

import (
	"go-clean-app/domain"
	"go-clean-app/infrastructure/dao"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) FindAll(ctx *gin.Context, db *gorm.DB) (domain.Users, error) {
	var daoUsers dao.Users
	err := db.Model(&daoUsers).Find(&daoUsers).Error
	if err != nil {
		return nil, err
	}
	entity, err := daoUsers.ToEntity()
	if err != nil {
		return nil, err
	}
	return entity, nil
}
