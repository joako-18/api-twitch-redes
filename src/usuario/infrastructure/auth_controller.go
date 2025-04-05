package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/NetflixAPIgo/src/usuario/application"
)

type AuthController struct {
	authUseCase *application.AuthUseCase
}

func NewAuthController(authUseCase *application.AuthUseCase) *AuthController {
	return &AuthController{authUseCase: authUseCase}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := ctrl.authUseCase.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	userResponse := gin.H{
		"id":       user.ID,
		"name":     user.Name,
		"lastname": user.Lastname,
		"username": user.Username,
		"email":    user.Email,
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  userResponse,
	})
}
