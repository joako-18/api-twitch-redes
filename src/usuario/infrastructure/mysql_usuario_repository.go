package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/NetflixAPIgo/src/core"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/domain/entities"
)

type MySQLUsuarioRepository struct {
	conn *sql.DB
}

func NewMySQLUsuarioRepository() domain.IUsuario {
	conn := core.GetDB()
	return &MySQLUsuarioRepository{conn: conn}
}

func (r *MySQLUsuarioRepository) Save(usuario *entities.Usuario) error {
	if err := usuario.HashPassword(); err != nil {
		log.Println("Error al hashear la contraseña:", err)
		return fmt.Errorf("error al procesar la contraseña")
	}

	query := `
		INSERT INTO users (name, lastname, username, password, email)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := r.conn.Exec(
		query,
		usuario.Name,
		usuario.Lastname,
		usuario.Username,
		usuario.Password,
		usuario.Email,
	)
	if err != nil {
		log.Println("Error al guardar el usuario:", err)
		return fmt.Errorf("error al crear el usuario")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID:", err)
		return fmt.Errorf("error al obtener el ID del usuario creado")
	}

	usuario.ID = int32(id)
	return nil
}

func (r *MySQLUsuarioRepository) Update(usuario *entities.Usuario) error {
	if usuario.Password != "" {
		if err := usuario.HashPassword(); err != nil {
			log.Println("Error al hashear la contraseña:", err)
			return fmt.Errorf("error al procesar la nueva contraseña")
		}
		query := `
			UPDATE users 
			SET name = ?, lastname = ?, username = ?, 
				password = COALESCE(?, password), 
				email = ? 
			WHERE id = ?
		`
		result, err := r.conn.Exec(
			query,
			usuario.Name,
			usuario.Lastname,
			usuario.Username,
			usuario.Password,
			usuario.Email,
			usuario.ID,
		)
		if err != nil {
			log.Println("Error al actualizar el usuario:", err)
			return fmt.Errorf("error al actualizar el usuario")
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return fmt.Errorf("usuario con ID %d no encontrado", usuario.ID)
		}
		return nil
	}

	query := `
		UPDATE users 
		SET name = ?, lastname = ?, username = ?, email = ? 
		WHERE id = ?
	`
	result, err := r.conn.Exec(
		query,
		usuario.Name,
		usuario.Lastname,
		usuario.Username,
		usuario.Email,
		usuario.ID,
	)
	if err != nil {
		log.Println("Error al actualizar el usuario:", err)
		return fmt.Errorf("error al actualizar el usuario")
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", usuario.ID)
	}

	return nil
}

func (r *MySQLUsuarioRepository) Delete(id int32) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := r.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el usuario:", err)
		return fmt.Errorf("error al eliminar el usuario")
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", id)
	}

	return nil
}

func (r *MySQLUsuarioRepository) GetById(id int32) (*entities.Usuario, error) {
	query := `
		SELECT id, name, lastname, username, password, email 
		FROM users 
		WHERE id = ?
	`
	row := r.conn.QueryRow(query, id)

	usuario := &entities.Usuario{}
	err := row.Scan(
		&usuario.ID,
		&usuario.Name,
		&usuario.Lastname,
		&usuario.Username,
		&usuario.Password,
		&usuario.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario con ID %d no encontrado", id)
		}
		log.Println("Error al buscar usuario por ID:", id, "error:", err)
		return nil, fmt.Errorf("error al obtener el usuario")
	}

	return usuario, nil
}

func (r *MySQLUsuarioRepository) GetAll() ([]entities.Usuario, error) {
	query := `
		SELECT id, name, lastname, username, email 
		FROM users
	`
	rows, err := r.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los usuarios:", err)
		return nil, fmt.Errorf("error al obtener los usuarios")
	}
	defer rows.Close()

	var usuarios []entities.Usuario
	for rows.Next() {
		var usuario entities.Usuario
		err := rows.Scan(
			&usuario.ID,
			&usuario.Name,
			&usuario.Lastname,
			&usuario.Username,
			&usuario.Email,
		)
		if err != nil {
			log.Println("Error al escanear usuario:", err)
			return nil, fmt.Errorf("error al procesar los datos de usuarios")
		}
		usuarios = append(usuarios, usuario)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error después de leer usuarios:", err)
		return nil, fmt.Errorf("error al finalizar la lectura de usuarios")
	}

	return usuarios, nil
}

func (r *MySQLUsuarioRepository) GetByEmail(email string) (*entities.Usuario, error) {
	query := `
		SELECT id, name, lastname, username, password, email 
		FROM users 
		WHERE email = ?
	`
	row := r.conn.QueryRow(query, email)

	usuario := &entities.Usuario{}
	err := row.Scan(
		&usuario.ID,
		&usuario.Name,
		&usuario.Lastname,
		&usuario.Username,
		&usuario.Password,
		&usuario.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario con email %s no encontrado", email)
		}
		log.Println("Error al buscar usuario por email:", email, "error:", err)
		return nil, fmt.Errorf("error al obtener el usuario por email")
	}

	return usuario, nil
}

func (r *MySQLUsuarioRepository) GetByUsername(username string) (*entities.Usuario, error) {
	query := `
		SELECT id, name, lastname, username, password, email 
		FROM users 
		WHERE username = ?
	`
	row := r.conn.QueryRow(query, username)

	usuario := &entities.Usuario{}
	err := row.Scan(
		&usuario.ID,
		&usuario.Name,
		&usuario.Lastname,
		&usuario.Username,
		&usuario.Password,
		&usuario.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario con username %s no encontrado", username)
		}
		log.Println("Error al buscar usuario por username:", username, "error:", err)
		return nil, fmt.Errorf("error al obtener el usuario por username")
	}

	return usuario, nil
}
