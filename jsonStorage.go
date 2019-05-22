package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Task struct {
	TaskDescription string
	Time            string
}

//JSONStorage system for storing data...
type JSONStorage struct {
	FileName string
	hello    string
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Save function to save data entered in a file mentioned...
func (storage JSONStorage) Save(tasksDescription []string) {
	fileData, err := ioutil.ReadFile(storage.FileName)

	todoList := []Task{}
	json.Unmarshal(fileData, &todoList)
	if !fileExists(storage.FileName) {
		emptyJSON, _ := json.Marshal(todoList)
		ioutil.WriteFile(storage.FileName, emptyJSON, 0666)
	}
	if err != nil {
		log.Fatal(err)
	}
	for _, taskDescription := range tasksDescription {
		task := Task{taskDescription, time.Now().Format("Mon Jan 2 15:04:05")}
		todoList = append(todoList, task)
	}
	jsonData, err := json.MarshalIndent(todoList, "", "")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(storage.FileName, jsonData, 0644)
}

//ShowToDos is for showing the data in the given file...
func (storage JSONStorage) ShowToDos() {
	file, err := ioutil.ReadFile(storage.FileName)
	if err != nil {
		log.Fatal(err)
	}
	todos := []Task{}
	json.Unmarshal(file, &todos)
	for _, task := range todos {
		fmt.Printf("Task: %s\n Time: %s\n", task.TaskDescription, task.Time)
	}
}

//Delete function for deleting the specified task in the file ...
func (storage JSONStorage) Delete(indexToBeDeleted int64) {
	file, err := ioutil.ReadFile(storage.FileName)
	if err != nil {
		log.Fatal(err)
	}
	todos := []Task{}
	json.Unmarshal(file, &todos)
	todos = append(todos[:indexToBeDeleted], todos[indexToBeDeleted:]...)
	jsonData, err := json.MarshalIndent(todos, "", "")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(storage.FileName, jsonData, 0644)
}
