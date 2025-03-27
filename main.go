package main

import (
	"net/http"

	config "github.com/223mali/go-todo/configs"
	"github.com/223mali/go-todo/docs"
	"github.com/223mali/go-todo/model"
	"github.com/223mali/go-todo/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var db = make(map[string]string)

// PingHandler godoc
// @Summary      Ping endpoint
// @Description  Returns pong message to verify API is running
// @Tags         health
// @Accept       json
// @Produce      plain
// @Success      200  {string}  string  "pong"
// @Router       /api/v1/ping [get]
func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func setupHealthAndSwaggerRoute(r *gin.Engine) *gin.Engine {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", pingHandler)

	return r
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	godotenv.Load(".env")
	docs.SwaggerInfo.Title = "Golang Todo API"
	docs.SwaggerInfo.Description = "This is a simple api to learn go."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	setupHealthAndSwaggerRoute(r)

	v1Group := r.Group("/api/v1")

	router.InitiateRouterV1(v1Group)

	db := config.ConnectToDatabase()
	db.AutoMigrate(&model.Task{})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
