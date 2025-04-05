package application

import (
	"errors"

	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain/entities"
)

type UpdateUsuarioUseCase struct {
	repo domain.IUsuario
}

func NewUpdateUsuarioUseCase(repo domain.IUsuario) *UpdateUsuarioUseCase {
	return &UpdateUsuarioUseCase{repo: repo}
}

func (uc *UpdateUsuarioUseCase) Run(usuario *entities.Usuario) (*entities.Usuario, error) {
	if usuario.ID <= 0 {
		return nil, errors.New("ID de usuario inválido")
	}

	existingUser, err := uc.repo.GetById(usuario.ID)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	if usuario.Email != existingUser.Email {
		_, err := uc.repo.GetByEmail(usuario.Email)
		if err == nil {
			return nil, errors.New("el email ya está registrado por otro usuario")
		}
	}

	if usuario.Username != existingUser.Username {
		_, err := uc.repo.GetByUsername(usuario.Username)
		if err == nil {
			return nil, errors.New("el nombre de usuario ya está en uso por otro usuario")
		}
	}

	err = uc.repo.Update(usuario)
	if err != nil {
		return nil, err
	}

	updatedUser, err := uc.repo.GetById(usuario.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
