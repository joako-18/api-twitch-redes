package infrastructure

import (
	"time"

	"github.com/vicpoo/NetflixAPIgo/src/video/application"
)

func InitVideoDependencies() (
	*CreateVideoController,
	*GetVideoController,
	*GetAllVideosController,
	*CacheVideoController,
	*UploadController,
) {
	repo := NewMySQLVideoRepository()
	cacheService := application.NewVideoCacheService("./video_cache", 7*24*time.Hour)

	createUseCase := application.NewCreateVideoUseCase(repo)
	getByIDUseCase := application.NewGetVideoByIDUseCase(repo)
	getAllUseCase := application.NewGetAllVideosUseCase(repo)

	uploadController := NewUploadController(repo)

	createController := NewCreateVideoController(createUseCase)
	getController := NewGetVideoController(getByIDUseCase)
	getAllController := NewGetAllVideosController(getAllUseCase)
	cacheController := NewCacheVideoController(cacheService, repo)

	return createController, getController, getAllController, cacheController, uploadController
}
