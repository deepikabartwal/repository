package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type task struct {
	Task string
	Time string
}
type JsonStorage struct {
	FileName string
}

func (storage *JsonStorage) SaveToFile(args []string) {
	file, err := ioutil.ReadFile(storage.FileName)
	if err != nil {
		log.Fatal(err)
	}

	data := []task{}
	json.Unmarshal(file, &data)
	for _, taskDescription := range args {
		task := task{taskDescription, time.Now().Format("Mon Jan 2 15:04:05")}
		data = append(data, task)
	}
	jsonData, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(storage.FileName, jsonData, 0644)
}

func (storage *JsonStorage) ShowToDos() {
	file, err := ioutil.ReadFile(storage.FileName)
	if err != nil {
		log.Fatal(err)
	}
	todos := []task{}
	json.Unmarshal(file, &todos)
	for _, task := range todos {
		fmt.Printf("Task: %s\n Time: %s\n", task.Task, task.Time)
	}
}

func (storage *JsonStorage) Delete(indexToBeDeleted int64) {
	file, err := ioutil.ReadFile(storage.FileName)
	if err != nil {
		log.Fatal(err)
	}
	todos := []task{}
	json.Unmarshal(file, &todos)
	todos = append(todos[:indexToBeDeleted], todos[indexToBeDeleted+1:]...)
	jsonData, err := json.MarshalIndent(todos, "", "")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(storage.FileName, jsonData, 0644)
}
