package application

import (
	"github.com/vicpoo/NetflixAPIgo/src/video/domain"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain/entities"
)

type GetAllVideosUseCase struct {
	repo domain.VideoRepository
}

func NewGetAllVideosUseCase(repo domain.VideoRepository) *GetAllVideosUseCase {
	return &GetAllVideosUseCase{repo: repo}
}

func (uc *GetAllVideosUseCase) Run() ([]entities.Video, error) {
	return uc.repo.GetAll()
}
