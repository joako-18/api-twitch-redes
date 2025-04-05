package application

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/vicpoo/NetflixAPIgo/src/video/domain/entities"
)

type VideoCacheService struct {
	CacheDir      string
	CacheDuration time.Duration
}

func NewVideoCacheService(cacheDir string, cacheDuration time.Duration) *VideoCacheService {
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		panic(fmt.Sprintf("No se pudo crear el directorio de caché: %v", err))
	}

	return &VideoCacheService{
		CacheDir:      cacheDir,
		CacheDuration: cacheDuration,
	}
}

func (s *VideoCacheService) DownloadVideo(video *entities.Video) error {
	if isYouTubeURL(video.URL) {
		return fmt.Errorf("no se puede cachear videos de YouTube directamente")
	}

	resp, err := http.Get(video.URL)
	if err != nil {
		return fmt.Errorf("error al descargar el video: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("respuesta no exitosa: %s", resp.Status)
	}

	filename := filepath.Join(s.CacheDir, fmt.Sprintf("video_%d%s", video.ID, filepath.Ext(video.URL)))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error al crear archivo local: %v", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		os.Remove(filename)
		return fmt.Errorf("error al guardar video localmente: %v", err)
	}

	video.SetLocalPath(filename)
	video.SetIsCached(true)
	video.SetCacheExpiry(time.Now().Add(s.CacheDuration))

	return nil
}

func (s *VideoCacheService) ClearCache(video *entities.Video) error {
	if !video.GetIsCached() || video.GetLocalPath() == "" {
		return nil
	}

	if err := os.Remove(video.GetLocalPath()); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error al eliminar video del caché: %v", err)
	}

	video.SetLocalPath("")
	video.SetIsCached(false)
	video.SetCacheExpiry(time.Time{})

	return nil
}

func isYouTubeURL(url string) bool {
	return strings.Contains(url, "youtube.com") || strings.Contains(url, "youtu.be")
}
