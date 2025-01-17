package controller_test

import (
	"testing"

	"go-clean-app/domain"
	"go-clean-app/presentation/controller"
	"go-clean-app/tests/helpers"
	mock_usecase "go-clean-app/tests/units/mock/application/usecase"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewsPost(t *testing.T) {

	ctx, _ := helpers.UnitInitialize()
	t.Run("UserController", func(t *testing.T) {

		t.Run("正常にControllerが終了すること", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			usecase := mock_usecase.NewMockUserUseCase(ctrl)
			usecase.EXPECT().FindAll(ctx).Return(domain.User{}, nil)

			c := controller.NewUserController(usecase)

			c.Users(ctx)
			assert.Equal(t, nil, nil)
		})
	})
}
