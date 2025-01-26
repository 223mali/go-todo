package controller

import (
	"fmt"
	"net/http"

	config "github.com/223mali/go-todo/src/configs"
	"github.com/223mali/go-todo/src/model"
	"github.com/223mali/go-todo/src/repository"
	"github.com/gin-gonic/gin"
)

// @Summary Get all tasks
// @Schemes
// @Description do ping
// @Tags Tasks
// @Accept json
// @Produce json
// @Success 200 {array} model.TaskRequest
// @Router /api/v1/tasks [get]
func GetTasksHandler(c *gin.Context) {

	db := config.ConnectToDatabase()
	taskRepository := repository.NewTaskRepositoryImpl(db)
	tasks, err := taskRepository.FindAll()

	tasksResponse := []model.TaskRequest{}
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, model.TaskRequest{
			Name:        task.Name,
			Description: task.Description,
		})
	}
	fmt.Println("tasksResponse", tasksResponse)
	c.JSON(http.StatusOK, gin.H{"tasks": tasksResponse, "error": err})
}

// @Summary Get Single tasks
// @Schemes
// @Description do ping
// @Tags Tasks
// @Accept json
// @Param test query string true "test" minlength(5)  maxlength(10)
// @Produce json
// @Success 200 {object} model.TaskRequest
// @Router /api/v1/tasks/task [get]
func GetTaskHandler(c *gin.Context) {

	db := config.ConnectToDatabase()
	taskRepository := repository.NewTaskRepositoryImpl(db)
	tasks, err := taskRepository.FindAll()

	tasksResponse := []model.TaskRequest{}
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, model.TaskRequest{
			Name:        task.Name,
			Description: task.Description,
		})
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasksResponse, "error": err})
}

type CreateTaskInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// @Summary Create Single task
// @Schemes
// @Description Creates a single task
// @Tags Tasks
// @Accept json
// @Param data body model.TaskRequest false "data"
// @Produce json
// @Success 200 {object} model.TaskRequest
// @Router /api/v1/tasks/task [post]
func CreateTask(c *gin.Context) {
	var createTask model.TaskRequest
	db := config.ConnectToDatabase()
	if err := c.ShouldBindJSON(&createTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("body", createTask)

	taskRepository := repository.NewTaskRepositoryImpl(db)

	response, error := taskRepository.CreateTask(createTask)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": response})

}
