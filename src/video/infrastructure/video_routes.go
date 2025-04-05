package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type VideoRouter struct {
	engine *gin.Engine
}

func NewVideoRouter(engine *gin.Engine) *VideoRouter {
	return &VideoRouter{engine: engine}
}

func (r *VideoRouter) Run() {
	createController, getController, getAllController, cacheController, uploadController := InitVideoDependencies()

	videoGroup := r.engine.Group("/videos")
	{
		videoGroup.POST("/", createController.Run)
		videoGroup.GET("/", getAllController.Run)
		videoGroup.GET("/:id", getController.Run)
		videoGroup.POST("/upload", uploadController.UploadHandler)
		videoGroup.POST("/:id/cache", cacheController.CacheVideoHandler)
		videoGroup.GET("/:id/cache", cacheController.GetCachedVideoStreamHandler)
	}
}
