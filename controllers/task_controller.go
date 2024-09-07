package controllers

import (
	"fmt"
	"net/http"
	"task-list/models"
	"task-list/repositories"

	"github.com/gin-gonic/gin"
)

// AddTask controller to add a task
func AddTask(c *gin.Context) {
	var display_name string = c.PostForm("display_name")

	if display_name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": true, "message": "display_name is required"})
		return
	}

	user_id, err := repositories.UserRepository.GetUserIdByDisplayName(display_name)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to get user"})
		fmt.Println("Failed to get user:", err)
		return
	}

	if user_id < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": true, "message": "User not found"})
		return
	}

	var task_header string = c.PostForm("task_header")

	if task_header == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": true, "message": "task_header is required"})
		return
	}

	var task_description string = c.PostForm("task_description")

	if task_description == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": true, "message": "task_description is required"})
		return
	}

	task, err := repositories.TaskRepository.AddTask(user_id, task_header, task_description)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to create task"})
		fmt.Println("Failed to create task:", err)
		return
	}

	user_task, err := repositories.UserTaskRepository.AddUserTask(user_id, task.Id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to create user task"})
		fmt.Println("Failed to create user task:", err)
		return
	}

	user_task_id := int(user_task.Id)

	user, err := repositories.UserRepository.GetUserByID(user_id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to get user"})
		fmt.Println("Failed to get user:", err)
		return
	}

	var user_task_response models.UserTaskResponse = models.UserTaskResponse{
		Id:       user_task_id,
		UserData: user,
		TaskData: task,
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"error": false, "data": user_task_response})

}

func GetTasksForUser(c *gin.Context) {
	var display_name string = c.PostForm("display_name")

	if display_name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": true, "message": "display_name is required"})
		return
	}

	user_id, err := repositories.UserRepository.GetUserIdByDisplayName(display_name)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to get user"})
		fmt.Println("Failed to get user:", err)
		return
	}

	if user_id < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": true, "message": "User not found"})
		return
	}

	user_task_ids, err := repositories.UserTaskRepository.GetTaskIdsForUser(user_id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": true, "message": "Failed to get tasks for user"})
		return
	}

	var user_tasks []models.Task = []models.Task{}

	for _, task_id := range user_task_ids {
		task, err := repositories.TaskRepository.GetTaskById(task_id)

		if err != nil {
			continue
		}

		if task.Id == -1 {
			continue
		}

		user_tasks = append(user_tasks, task)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"error": false, "data": user_tasks})
}
