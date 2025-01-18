package usecase_test

import (
	"go-clean-app/application/interactor"
	"go-clean-app/domain"
	"go-clean-app/test/helpers"
	"log"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUser(t *testing.T) {

	envPath := "../../../.env.test"
	ctx, db := helpers.Initialize(envPath)

	t.Run("UserInteractor", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		userPort := mock_port.NewMockUserPort(ctrl)
		awsPort := mock_port.NewMockAWSPort(ctrl)
		interfactor := interactor.NewUserInteractor(db, awsPort, userPort)

		t.Run("FindAll_正常に処理が終了すること", func(t *testing.T) {

			userPort.EXPECT().Users(ctx).Return(
				&domain.Users{
					&domain.User{
						ID:        1,
						Name:      "test",
						ImageURL:  nil,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
				nil,
			)
			users, err := interfactor.FindAll(ctx)
			if err != nil {
				log.Fatalf("Failed to FindAll: %v\n", err)
			}
			assert.Equal(t, users[0].ID, 1)
			assert.Equal(t, users[0].Name, "test")
		})
	})
}
