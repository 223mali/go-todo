package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string
	Description string
}

type TaskRequest struct {
	Name        string `json:"name" example:"task name"`
	Description string `json:"description" example:"task description"`
}

type UpdateTaskRequest struct {
	Id int `json:"id" example:"1"`
}
