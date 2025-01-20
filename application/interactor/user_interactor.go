package interactor

import (
	"fmt"
	"go-clean-app/application/port"
	"go-clean-app/domain"
	"go-clean-app/infrastructure/driver"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	var mu sync.Mutex
	var wg sync.WaitGroup
	var errors []error
	var errorsMu sync.Mutex
	for shardID, db := range interactor.db.GetShards() {
		wg.Add(1)
		go func(id int, db *gorm.DB) {
			defer wg.Done()
			users, err := interactor.userPort.FindAll(ctx, db)
			if err != nil {
				errorsMu.Lock()
				errors = append(errors, fmt.Errorf("failed to fetch users from shard %d: %w", id, err))
				errorsMu.Unlock()
				return
			}
			for _, user := range users {
				// TODO N+1問題
				url, err := interactor.awsPort.CreatePreSignedURL("testBucket", user.ID)
				if err != nil {
					errorsMu.Lock()
					errors = append(errors, fmt.Errorf("failed to fetch image from s3 %d: %w", user.ID, err))
					errorsMu.Unlock()
					return
				}
				user.ImageURL = url
				users = append(users, user)
			}
			mu.Lock()
			allUsers = append(allUsers, users...)
			mu.Unlock()
		}(shardID, db)
	}
	wg.Wait()
	if len(errors) > 0 {
		return nil, fmt.Errorf("errors occurred: %v", errors)
	}
	return allUsers, nil
}
