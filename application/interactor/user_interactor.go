package interactor

import (
	"go-clean-app/application/port"
	"go-clean-app/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserInteractor struct {
	db       *gorm.DB
	awsPort  port.AWSPort
	userPort port.UserPort
}

func NewUserInteractor(
	db *gorm.DB,
	awsPort port.AWSPort,
	userPort port.UserPort,
) *UserInteractor {
	return &UserInteractor{
		db:       db,
		awsPort:  awsPort,
		userPort: userPort,
	}
}

func (interactor *UserInteractor) FindAll(ctx *gin.Context) (domain.Users, error) {
	_users, err := interactor.userPort.FindAll(ctx, interactor.db)
	if err != nil {
		return nil, err
	}
	var users domain.Users
	for _, user := range _users {
		url, err := interactor.awsPort.CreatePreSignedURL("testBucket", user.ID)
		if err != nil {
		}
		user.ImageURL = url
		users = append(users, user)
	}
	return users, nil
}
