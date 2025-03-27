package router

import (
	"github.com/223mali/go-todo/src/controller"
	"github.com/gin-gonic/gin"
)

func GetTaskRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	taskGroup := r.Group("/tasks")

	taskGroup.GET("", controller.GetTasksHandler)
	taskGroup.GET("task", controller.GetTaskHandler)
	taskGroup.POST("task", controller.CreateTask)
	taskGroup.PUT("task/:id", controller.UpdateTask)

	return taskGroup

}
