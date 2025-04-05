package application

import (
	"github.com/vicpoo/NetflixAPIgo/src/video/domain"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain/entities"
)

type CreateVideoUseCase struct {
	repo domain.VideoRepository
}

func NewCreateVideoUseCase(repo domain.VideoRepository) *CreateVideoUseCase {
	return &CreateVideoUseCase{repo: repo}
}

func (uc *CreateVideoUseCase) Run(video *entities.Video) (*entities.Video, error) {
	if err := uc.repo.Save(video); err != nil {
		return nil, err
	}
	return video, nil
}
