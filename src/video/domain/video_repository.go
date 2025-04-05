package domain

import "github.com/vicpoo/NetflixAPIgo/src/video/domain/entities"

type VideoRepository interface {
	Save(video *entities.Video) error
	GetByID(id int) (*entities.Video, error)
	GetAll() ([]entities.Video, error)
}
