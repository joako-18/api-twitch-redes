package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/NetflixAPIgo/src/video/application"
)

type GetVideoController struct {
	useCase *application.GetVideoByIDUseCase
}

func NewGetVideoController(useCase *application.GetVideoByIDUseCase) *GetVideoController {
	return &GetVideoController{useCase: useCase}
}

func (ctrl *GetVideoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	video, err := ctrl.useCase.Run(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Video no encontrado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, video)
}
