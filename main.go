package main

import (
	"log"
	"os"

	"github.com/vicpoo/NetflixAPIgo/src/core"
	usuarioInfra "github.com/vicpoo/NetflixAPIgo/src/usuario/infrastructure"
	videoInfra "github.com/vicpoo/NetflixAPIgo/src/video/infrastructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	core.InitDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	router.Static("/uploads", "./uploads")
	router.Static("/video_cache", "./video_cache")

	os.MkdirAll("./uploads", 0755)
	os.MkdirAll("./video_cache", 0755)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	usuarioRouter := usuarioInfra.NewUsuarioRouter(router)
	usuarioRouter.Run()

	videoRouter := videoInfra.NewVideoRouter(router)
	videoRouter.Run()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("\n Servidor iniciado en http://localhost:%s", port)
	log.Println(" Rutas estáticas:")
	log.Println("   - /uploads para videos subidos")
	log.Println("   - /video_cache para videos cacheados")

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
