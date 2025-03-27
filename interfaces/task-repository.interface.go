package interfaces

import "github.com/223mali/go-todo/model"

type TaskRepository interface {
	FindAll() ([]model.Task, error)
	CreateTask(task model.TaskRequest) (model.Task, error)
}
