package main

import (
	"fmt"
)

// Define a struct (like a class in OOP)
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// Define a container (slice) to store tasks
type ToDoList struct {
	Tasks []Task
}

// Method: add a task
func (t *ToDoList) AddTask(title string) {
	newTask := Task{
		ID:        len(t.Tasks) + 1,
		Title:     title,
		Completed: false,
	}
	t.Tasks = append(t.Tasks, newTask)
	fmt.Println("Added task:", title)
}

// Method: mark a task as done
func (t *ToDoList) CompleteTask(id int) {
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks[i].Completed = true
			fmt.Println("Task completed:", t.Tasks[i].Title)
			return
		}
	}
	fmt.Println("Task not found with ID:", id)
}

// Method: list all tasks
func (t *ToDoList) ShowTasks() {
	if len(t.Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("Your tasks:")
	for _, task := range t.Tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("[%d] %s %s\n", task.ID, status, task.Title)
	}
}
