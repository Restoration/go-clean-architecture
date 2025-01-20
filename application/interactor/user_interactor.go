package interactor

import (
	"go-clean-app/application/port"
	"go-clean-app/domain"
	"go-clean-app/infrastructure/driver"

	"github.com/gin-gonic/gin"
)

type UserInteractor struct {
	db       *driver.ShardingManager
	awsPort  port.AWSPort
	userPort port.UserPort
}

func NewUserInteractor(
	db *driver.ShardingManager,
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
	shID := interactor.db.GetShardID(1001)
	db := interactor.db.GetDBForUser(shID)
	_users, err := interactor.userPort.FindAll(ctx, db)
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
