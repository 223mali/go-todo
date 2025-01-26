package interfaces

import "github.com/223mali/go-todo/src/model"

type TaskRepository interface {
	FindAll() ([]model.Task, error)
	CreateTask(task model.TaskRequest) (model.Task, error)
}
