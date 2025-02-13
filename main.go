package main

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	ID       int    `json:"id"`
	TaskName string `json:"text"`
	Status   bool   `json:"done"`
}

var tasks []Task
var lastId = 0

// CRUD-операции

func CreateTask(name string, status bool) {
	lastId++
	newTask := Task{
		ID:       lastId,
		TaskName: name,
		Status:   status,
	}
	tasks = append(tasks, newTask)
	// fmt.Println("Новый задача добавлена:", newTask)
}

func ReadTasts() {
	fmt.Println("Полученные задачи:")
	for _, task := range tasks {
		jsonData, _ := json.Marshal(task)
		fmt.Println(string(jsonData))
	}
}

func UpdateTask(id int, name string, status bool) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].TaskName = name
			tasks[i].Status = status
			fmt.Println("Задача обновлена:", tasks[i])
			return
		}
	}
	fmt.Printf("Задача с %d не найдена", id)
}

func main() {
	CreateTask("Сделать функцию для создании новый задачи", false)
}
