package main

import (
	"fmt"
	"go-todo/docs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db = make(map[string]string)

// PingHandler godoc
// @Summary      Ping endpoint
// @Description  Returns pong message to verify API is running
// @Tags         health
// @Accept       json
// @Produce      plain
// @Success      200  {string}  string  "pong"
// @Router       /ping [get]
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

	r.Run(":8080")

	dsn := fmt.Sprintf("host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("DB connection SUCCESS")
	}

	// Listen and Server in 0.0.0.0:8080
}
