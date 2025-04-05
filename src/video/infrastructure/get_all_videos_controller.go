package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/NetflixAPIgo/src/video/application"
)

type GetAllVideosController struct {
	useCase *application.GetAllVideosUseCase
}

func NewGetAllVideosController(useCase *application.GetAllVideosUseCase) *GetAllVideosController {
	return &GetAllVideosController{useCase: useCase}
}

func (ctrl *GetAllVideosController) Run(c *gin.Context) {
	videos, err := ctrl.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al obtener los videos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, videos)
}
