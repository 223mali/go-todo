package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary ping example
	// @Schemes
	// @Description do ping
	// @Tags example
	// @Accept json
	// @Produce json
	// @Success 200 {string} Helloworld
	// @Router /tasks/ [get]
func tasksHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello list of tasks")
}
func GetTaskRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	taskGroup := r.Group("/tasks")

	
	taskGroup.GET("/", tasksHandler)

	return taskGroup

}
