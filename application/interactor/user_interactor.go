package interactor

import (
	"fmt"
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
	var allUsers domain.Users
	for _, db := range interactor.db.GetShards() {
		users, err := interactor.userPort.FindAll(ctx, db)
		if err != nil {
			return nil, err
		}
		for i := range users {
			// TODO N+1問題
			// UserIDsを渡して一括で取得する作りが理想
			url, err := interactor.awsPort.CreatePreSignedURL("testBucket", users[i].ID)
			if err != nil {
				return nil, err
			}
			users[i].ImageURL = url
		}
		fmt.Println(db)
		allUsers = append(allUsers, users...)
	}
	return allUsers, nil
}
