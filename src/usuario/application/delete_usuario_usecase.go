package application

import (
	"errors"

	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain"
)

type DeleteUsuarioUseCase struct {
	repo domain.IUsuario
}

func NewDeleteUsuarioUseCase(repo domain.IUsuario) *DeleteUsuarioUseCase {
	return &DeleteUsuarioUseCase{repo: repo}
}

func (uc *DeleteUsuarioUseCase) Run(id int32) error {
	if id <= 0 {
		return errors.New("ID de usuario inválido")
	}

	_, err := uc.repo.GetById(id)
	if err != nil {
		return errors.New("usuario no encontrado")
	}

	return uc.repo.Delete(id)
}
