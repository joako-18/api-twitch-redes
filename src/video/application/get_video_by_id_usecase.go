package application

import (
	"github.com/vicpoo/NetflixAPIgo/src/video/domain"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain/entities"
)

type GetVideoByIDUseCase struct {
	repo domain.VideoRepository
}

func NewGetVideoByIDUseCase(repo domain.VideoRepository) *GetVideoByIDUseCase {
	return &GetVideoByIDUseCase{repo: repo}
}

func (uc *GetVideoByIDUseCase) Run(id int) (*entities.Video, error) {
	return uc.repo.GetByID(id)
}
