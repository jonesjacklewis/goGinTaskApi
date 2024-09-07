package models

// Task is a struct that represents a task
// It has an Id, TaskHeader, TaskDescription and Complete
type Task struct {
	Id              int
	TaskHeader      string
	TaskDescription string
	Complete        bool
}
