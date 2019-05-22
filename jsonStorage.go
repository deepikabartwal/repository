package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type task struct {
	Task string
	Time string
}

//JSONStorage system for storing data...
type JSONStorage struct {
	FileName string
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// SaveToFile function to save data entered in a file mentioned...
func (storage *JSONStorage) SaveToFile(args []string) {
	file, err := ioutil.ReadFile(storage.FileName)

	data := []task{}
	json.Unmarshal(file, &data)
	if !fileExists(storage.FileName) {
		emptyJSON, _ := json.Marshal(data)
		ioutil.WriteFile(storage.FileName, emptyJSON, 0666)
	}
	if err != nil {
		log.Fatal(err)
	}
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

//ShowToDos is for showing the data in the given file...
func (storage *JSONStorage) ShowToDos() {
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

//Delete function for deleting the specified task in the file ...
func (storage *JSONStorage) Delete(indexToBeDeleted int64) {
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
