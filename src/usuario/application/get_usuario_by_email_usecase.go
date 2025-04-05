package application

import (
	"errors"

	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain/entities"
)

type GetUsuarioByEmailUseCase struct {
	repo domain.IUsuario
}

func NewGetUsuarioByEmailUseCase(repo domain.IUsuario) *GetUsuarioByEmailUseCase {
	return &GetUsuarioByEmailUseCase{repo: repo}
}

func (uc *GetUsuarioByEmailUseCase) Run(email string) (*entities.Usuario, error) {
	if email == "" {
		return nil, errors.New("el email no puede estar vacío")
	}

	usuario, err := uc.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return usuario, nil
}
