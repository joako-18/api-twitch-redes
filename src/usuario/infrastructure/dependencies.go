package infrastructure

import (
	"github.com/vicpoo/NetflixAPIgo/src/usuario/application"
)

func InitUsuarioDependencies() (
	*CreateUsuarioController,
	*GetUsuarioByIdController,
	*UpdateUsuarioController,
	*DeleteUsuarioController,
	*GetAllUsuariosController,
	*GetUsuarioByEmailController,
	*AuthController,
) {
	repo := NewMySQLUsuarioRepository()

	createUseCase := application.NewCreateUsuarioUseCase(repo)
	getByIdUseCase := application.NewGetUsuarioByIdUseCase(repo)
	updateUseCase := application.NewUpdateUsuarioUseCase(repo)
	deleteUseCase := application.NewDeleteUsuarioUseCase(repo)
	getAllUseCase := application.NewGetAllUsuariosUseCase(repo)
	getByEmailUseCase := application.NewGetUsuarioByEmailUseCase(repo)
	authUseCase := application.NewAuthUseCase(repo)

	createController := NewCreateUsuarioController(createUseCase)
	getByIdController := NewGetUsuarioByIdController(getByIdUseCase)
	updateController := NewUpdateUsuarioController(updateUseCase)
	deleteController := NewDeleteUsuarioController(deleteUseCase)
	getAllController := NewGetAllUsuariosController(getAllUseCase)
	getByEmailController := NewGetUsuarioByEmailController(getByEmailUseCase)
	authController := NewAuthController(authUseCase)

	return createController, getByIdController, updateController,
		deleteController, getAllController, getByEmailController,
		authController
}
