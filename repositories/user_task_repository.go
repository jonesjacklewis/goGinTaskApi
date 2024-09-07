package repositories

import (
	"task-list/config"
	"task-list/models"
)

type UserTaskRepositoryInterface interface {
	AddUserTask(user_id int, task_id int) (models.UserTask, error)
}

type UserTaskRepositoryImpl struct{}

var UserTaskRepository UserTaskRepositoryInterface = &UserTaskRepositoryImpl{}

func (r *UserTaskRepositoryImpl) AddUserTask(user_id int, task_id int) (models.UserTask, error) {

	var query string = `
	INSERT INTO
	UsersTasks (
	UsersId,
	TasksID
	)
	VALUES (
	?,
	?
	)
	`

	result, err := config.Database_Connection.Exec(query, user_id, task_id)

	if err != nil {
		return models.UserTask{}, err
	}

	id, err := result.LastInsertId()

	if id < 1 {
		return models.UserTask{}, err
	}

	var user_task models.UserTask = models.UserTask{
		UserId: user_id,
		TaskId: task_id,
		Id:     int(id),
	}

	return user_task, nil
}
