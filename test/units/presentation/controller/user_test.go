package controller_test

import (
	"testing"

	"go-clean-app/domain"
	"go-clean-app/presentation/controller"
	"go-clean-app/test/helpers"
	mock_usecase "go-clean-app/test/units/mock/application/usecase"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUser(t *testing.T) {

	envPath := "../../../../.env.test"
	ctx, _ := helpers.Initialize(envPath)
	t.Run("UserController", func(t *testing.T) {

		t.Run("正常にControllerが終了すること", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := mock_usecase.NewMockUserUseCase(ctrl)
			uc.EXPECT().FindAll(ctx).Return(domain.Users{}, nil)
			c := controller.NewUserController(uc)

			c.Users(ctx)
			assert.Equal(t, nil, nil)
		})
	})
}
