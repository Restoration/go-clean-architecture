package di

import (
	"go-clean-app/application/interactor"
	"go-clean-app/application/port"
	"go-clean-app/infrastructure/driver"
	"go-clean-app/infrastructure/repository"
	"go-clean-app/presentation/controller"
)

func DiUser(db *driver.ShardingManager) *controller.UserController {
	return controller.NewUserController(
		DiUserUseCase(db),
	)
}

func DiUserUseCase(db *driver.ShardingManager) *interactor.UserInteractor {
	return interactor.NewUserInteractor(
		db,
		DiAWSGateway(),
		DiUserRepository(),
	)
}

func DiUserRepository() port.UserPort {
	return repository.NewUserRepository()
}
