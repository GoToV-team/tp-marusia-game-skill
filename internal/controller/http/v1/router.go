// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	game "github.com/ThCompiler/go_game_constractor/director"
	"github.com/ThCompiler/go_game_constractor/marusia/hub"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/evrone/go-clean-template/docs"
	"github.com/evrone/go-clean-template/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Lemonade game marusia skill
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, opLemonade game.SceneDirectorConfig, opGarden game.SceneDirectorConfig, hub *hub.ScriptHub) {
	// Options
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost", "https://skill-debugger.marusia.mail.ru"}
	corsConfig.AllowMethods = []string{"POST"}

	handler.Use(cors.New(corsConfig))
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	//// K8s probe
	handler.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	//// Prometheus metrics
	//handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newLemonadeSkillRoute(h, opLemonade, hub, l)
		newBotanicalGardenSkillRoute(h, opGarden, hub, l)
	}
}
