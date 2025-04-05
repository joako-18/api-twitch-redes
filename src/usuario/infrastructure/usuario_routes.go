package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type UsuarioRouter struct {
	engine *gin.Engine
}

func NewUsuarioRouter(engine *gin.Engine) *UsuarioRouter {
	return &UsuarioRouter{
		engine: engine,
	}
}

func (router *UsuarioRouter) Run() {
	createController, getByIdController, updateController,
		deleteController, getAllController, getByEmailController,
		authController := InitUsuarioDependencies()

	usuarioGroup := router.engine.Group("/usuarios")
	{
		usuarioGroup.POST("/", createController.Run)
		usuarioGroup.GET("/", getAllController.Run)
		usuarioGroup.GET("/:id", getByIdController.Run)
		usuarioGroup.PUT("/:id", updateController.Run)
		usuarioGroup.DELETE("/:id", deleteController.Run)
		usuarioGroup.GET("/email", getByEmailController.Run)
	}

	authGroup := router.engine.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
	}
}
