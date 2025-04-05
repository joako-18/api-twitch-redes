package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/NetflixAPIgo/src/video/application"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain/entities"
)

type CreateVideoController struct {
	useCase *application.CreateVideoUseCase
}

func NewCreateVideoController(useCase *application.CreateVideoUseCase) *CreateVideoController {
	return &CreateVideoController{useCase: useCase}
}

func (ctrl *CreateVideoController) Run(c *gin.Context) {
	var request struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		URL         string `json:"url" binding:"required,url"`
		UserID      int    `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	video := &entities.Video{
		Title:       request.Title,
		Description: request.Description,
		URL:         request.URL,
		UserID:      request.UserID,
	}

	if _, err := ctrl.useCase.Run(video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el video",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, video)
}
