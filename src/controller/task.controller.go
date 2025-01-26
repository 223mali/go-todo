package controller

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
func GetTasksHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello list of tasks")
}
