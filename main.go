package main

import (
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

func main() {
	task1 := Task{ID: 1, TaskName: "Разбить ТЗ на подзадачи", Status: false}

	task2 := Task{}
	task2.ID = 2
	task2.TaskName = "Оперелить структуру данных"
	task2.Status = true
	fmt.Println(task1)
	fmt.Println(task2)
}
