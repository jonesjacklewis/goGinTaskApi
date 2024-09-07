package repositories

import (
	"task-list/config"
	"task-list/models"
)

// TaskRepositoryInterface is an interface for the TaskRepository
// It has the method AddTask
type TaskRepositoryInterface interface {
	AddTask(user_id int, task_header string, task_description string) (models.Task, error)
}

// TaskRepositoryImpl is a struct that implements the TaskRepositoryInterface
type TaskRepositoryImpl struct{}

// TaskRepository is an instance of the TaskRepositoryInterface
var TaskRepository TaskRepositoryInterface = &TaskRepositoryImpl{}

// CreateUser creates a user
// It takes a display_name string as a parameter
// It returns a user and an error
func (r *TaskRepositoryImpl) AddTask(user_id int, task_header string, task_description string) (models.Task, error) {

	var query string = `
	INSERT INTO
	Tasks (
	TaskHeader,
	TaskDescription
	)
	VALUES (
	?,
	?
	)
	`

	result, err := config.Database_Connection.Exec(query, task_header, task_description)

	if err != nil {
		return models.Task{}, err
	}

	id, err := result.LastInsertId()

	if id < 1 {
		return models.Task{}, err
	}

	if err != nil {
		return models.Task{}, err
	}

	var task models.Task = models.Task{
		Id:              int(id),
		TaskHeader:      task_header,
		TaskDescription: task_description,
		Complete:        false,
	}

	return task, nil
}
