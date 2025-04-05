package application

import (
	"errors"

	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain/entities"
)

type CreateUsuarioUseCase struct {
	repo domain.IUsuario
}

func NewCreateUsuarioUseCase(repo domain.IUsuario) *CreateUsuarioUseCase {
	return &CreateUsuarioUseCase{repo: repo}
}

func (uc *CreateUsuarioUseCase) Run(usuario *entities.Usuario) (*entities.Usuario, error) {
	_, err := uc.repo.GetByEmail(usuario.Email)
	if err == nil {
		return nil, errors.New("el email ya está registrado")
	}

	_, err = uc.repo.GetByUsername(usuario.Username)
	if err == nil {
		return nil, errors.New("el nombre de usuario ya está en uso")
	}

	err = uc.repo.Save(usuario)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}
