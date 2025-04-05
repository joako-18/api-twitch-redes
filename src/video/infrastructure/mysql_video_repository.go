package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/vicpoo/NetflixAPIgo/src/core"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain/entities"
)

type MySQLVideoRepository struct {
	conn *sql.DB
}

func NewMySQLVideoRepository() domain.VideoRepository {
	conn := core.GetDB()
	return &MySQLVideoRepository{conn: conn}
}

func (r *MySQLVideoRepository) Save(video *entities.Video) error {
	var query string
	var args []interface{}

	if video.GetID() == 0 {
		query = `
			INSERT INTO videos 
			(title, description, url, local_path, is_cached, cache_expiry, user_id) 
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`
		args = []interface{}{
			video.GetTitle(),
			video.GetDescription(),
			video.GetURL(),
			video.GetLocalPath(),
			video.GetIsCached(),
			nullTime(video.GetCacheExpiry()),
			video.GetUserID(),
		}
	} else {
		query = `
			UPDATE videos SET 
			title = ?, description = ?, url = ?, local_path = ?, 
			is_cached = ?, cache_expiry = ?, user_id = ?
			WHERE id = ?
		`
		args = []interface{}{
			video.GetTitle(),
			video.GetDescription(),
			video.GetURL(),
			video.GetLocalPath(),
			video.GetIsCached(),
			nullTime(video.GetCacheExpiry()),
			video.GetUserID(),
			video.GetID(),
		}
	}

	result, err := r.conn.Exec(query, args...)
	if err != nil {
		log.Printf("Error al guardar video: %v\nQuery: %s\nArgs: %+v", err, query, args)
		return fmt.Errorf("error al guardar el video")
	}

	if video.GetID() == 0 {
		id, err := result.LastInsertId()
		if err != nil {
			log.Println("Error al obtener ID:", err)
			return fmt.Errorf("error al obtener ID del video")
		}
		video.SetID(int(id))
	}

	return nil
}

func (r *MySQLVideoRepository) GetByID(id int) (*entities.Video, error) {
	query := `
		SELECT id, title, description, url, local_path, is_cached, cache_expiry, user_id
		FROM videos 
		WHERE id = ?
	`
	row := r.conn.QueryRow(query, id)

	var video entities.Video
	var cacheExpiry sql.NullTime

	err := row.Scan(
		&video.ID,
		&video.Title,
		&video.Description,
		&video.URL,
		&video.LocalPath,
		&video.IsCached,
		&cacheExpiry,
		&video.UserID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("video con ID %d no encontrado", id)
		}
		log.Printf("Error al buscar video por ID %d: %v", id, err)
		return nil, fmt.Errorf("error al obtener el video")
	}

	if cacheExpiry.Valid {
		video.CacheExpiry = cacheExpiry.Time
	}

	return &video, nil
}

func (r *MySQLVideoRepository) GetAll() ([]entities.Video, error) {
	query := `
		SELECT id, title, description, url, local_path, is_cached, cache_expiry, user_id
		FROM videos
		ORDER BY id DESC
	`
	rows, err := r.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los videos:", err)
		return nil, fmt.Errorf("error al listar videos")
	}
	defer rows.Close()

	var videos []entities.Video
	for rows.Next() {
		var video entities.Video
		var cacheExpiry sql.NullTime

		if err := rows.Scan(
			&video.ID,
			&video.Title,
			&video.Description,
			&video.URL,
			&video.LocalPath,
			&video.IsCached,
			&cacheExpiry,
			&video.UserID,
		); err != nil {
			log.Println("Error al escanear video:", err)
			continue
		}

		if cacheExpiry.Valid {
			video.CacheExpiry = cacheExpiry.Time
		}

		videos = append(videos, video)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de leer videos:", err)
		return nil, fmt.Errorf("error al procesar los videos")
	}

	return videos, nil
}

func nullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: t, Valid: true}
}
