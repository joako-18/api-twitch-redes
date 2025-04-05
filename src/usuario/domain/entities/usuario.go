package entities

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	ID       int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"not null"`
	Lastname string `json:"lastname" gorm:"not null"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
}

func (u *Usuario) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *Usuario) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *Usuario) SetID(id int32) {
	u.ID = id
}

func (u *Usuario) SetName(name string) {
	u.Name = name
}

func (u *Usuario) SetLastname(lastname string) {
	u.Lastname = lastname
}

func (u *Usuario) SetUsername(username string) {
	u.Username = username
}

func (u *Usuario) SetPassword(password string) {
	u.Password = password
}

func (u *Usuario) SetEmail(email string) {
	u.Email = email
}

func (u *Usuario) GetID() int32 {
	return u.ID
}

func (u *Usuario) GetName() string {
	return u.Name
}

func (u *Usuario) GetLastname() string {
	return u.Lastname
}

func (u *Usuario) GetUsername() string {
	return u.Username
}

func (u *Usuario) GetPassword() string {
	return u.Password
}

func (u *Usuario) GetEmail() string {
	return u.Email
}

func (u *Usuario) ToJSON() ([]byte, error) {
	type Alias Usuario
	return json.Marshal(&struct {
		*Alias
		Password string `json:"password,omitempty"`
	}{
		Alias:    (*Alias)(u),
		Password: "",
	})
}

func NewUsuario(name, lastname, username, password, email string) *Usuario {
	return &Usuario{
		Name:     name,
		Lastname: lastname,
		Username: username,
		Password: password,
		Email:    email,
	}
}
