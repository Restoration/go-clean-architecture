package di

import (
	"go-clean-app/application/interactor"
	"go-clean-app/application/port"
	"go-clean-app/infrastructure/repository"
	"go-clean-app/presentation/controller"

	"gorm.io/gorm"
)

func DiUser(db *gorm.DB) *controller.UserController {
	return controller.NewUserController(
		DiUserUseCase(db),
	)
}

func DiUserUseCase(db *gorm.DB) *interactor.UserInteractor {
	return interactor.NewUserInteractor(
		db,
		DiAWSGateway(),
		DiUserRepository(),
	)
}

func DiUserRepository() port.UserPort {
	return repository.NewUserRepository()
}
