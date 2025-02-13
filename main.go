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
<<<<<<< HEAD

func main() {
	task1 := Task{ID: 1, TaskName: "Разбить ТЗ на подзадачи", Status: false}

	task2 := Task{}
	task2.ID = 2
	task2.TaskName = "Оперелить структуру данных"
	task2.Status = true
	fmt.Println(task1)
	fmt.Println(task2)
}
func CreateTask(name string, status bool) {
	lastId++
	newTask := Task{
		ID:       lastId,
		TaskName: name,
		Status:   status,
	}
	tasks = append(tasks, newTask)
	fmt.Println("Новый задача добавлена:", newTask)
}

func main() {
	CreateTask("Сделать функцию для создании новый задачи", false)
>>>>>>> 23499ec (добавлена функция CreatTask создания новой задачи)
}
