package infrastructure

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/NetflixAPIgo/src/video/application"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain"
)

type CacheVideoController struct {
	cacheService *application.VideoCacheService
	repo         domain.VideoRepository
}

func NewCacheVideoController(
	cacheService *application.VideoCacheService,
	repo domain.VideoRepository,
) *CacheVideoController {
	return &CacheVideoController{
		cacheService: cacheService,
		repo:         repo,
	}
}

func (ctrl *CacheVideoController) CacheVideoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de video inválido"})
		return
	}

	video, err := ctrl.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video no encontrado"})
		return
	}

	if video.IsCacheValid() {
		c.JSON(http.StatusOK, gin.H{
			"message": "El video ya está disponible offline",
			"video":   video,
		})
		return
	}

	if strings.Contains(video.GetURL(), "youtube.com") || strings.Contains(video.GetURL(), "youtu.be") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se puede cachear videos de YouTube"})
		return
	}

	videoUrl := video.GetURL()
	if !strings.HasPrefix(videoUrl, "http") {
		videoUrl = "http://localhost:8000" + videoUrl
		video.SetURL(videoUrl)
	}

	if err := ctrl.cacheService.DownloadVideo(video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al descargar el video",
			"details": err.Error(),
		})
		return
	}

	if err := ctrl.repo.Save(video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al actualizar el video",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Video almacenado para uso offline",
		"video":   video,
	})
}

func (ctrl *CacheVideoController) GetCachedVideoStreamHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de video inválido"})
		return
	}

	video, err := ctrl.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video no encontrado"})
		return
	}

	if !video.IsCacheValid() {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "El video no está disponible para visualización offline",
		})
		return
	}

	c.File(video.GetLocalPath())
}
