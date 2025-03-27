package repository

import (
	"errors"
	"fmt"

	"github.com/223mali/go-todo/interfaces"
	"github.com/223mali/go-todo/model"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	Db *gorm.DB
}

func NewTaskRepositoryImpl(db *gorm.DB) interfaces.TaskRepository {
	fmt.Println("Heyo")
	return &TaskRepositoryImpl{Db: db}
}

func (t *TaskRepositoryImpl) FindAll() ([]model.Task, error) {
	var tasks []model.Task

	result := t.Db.Find(&tasks)

	if result.Error != nil {
		fmt.Println("FFUUCKK", result.Error)
		panic("Error fetching tasks")
	}

	if result.RowsAffected == 0 {
		return []model.Task{}, errors.New("No tasks found")
	}

	return tasks, nil
}

func (t *TaskRepositoryImpl) CreateTask(task model.TaskRequest) (model.Task, error) {
	newTask := model.Task{
		Name:        task.Name,
		Description: task.Description,
	}

	result := t.Db.Create(&newTask)

	if result.Error != nil {
		return model.Task{}, result.Error
	}

	return newTask, nil
}

func (t *TaskRepositoryImpl) UpdateTask(id int, updateObject model.TaskRequest) (model.Task, error) {
	var task model.Task

	result := t.Db.First(&task, id)

	if result.Error != nil {
		return model.Task{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Task{}, errors.New("No task found")
	}

	task.Name = updateObject.Name
	task.Description = updateObject.Description

	updateResult := t.Db.Save(&task)

	if updateResult.Error != nil {
		return model.Task{}, updateResult.Error
	}

	return task, result.Error
}
