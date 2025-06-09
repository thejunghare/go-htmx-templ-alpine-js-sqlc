package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/thejunghare/db"
	"github.com/thejunghare/models"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Tasks
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func GetAllTasksTempl() ([]models.Tasks, error) {
	var tasks []models.Tasks
	result := db.DB.Find(&tasks)
	return tasks, result.Error
}

func AddTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	json.NewDecoder(r.Body).Decode(&task)
	result := db.DB.Create(&task)
	if result.Error != nil {
		panic(result.Error)
	}
	json.NewEncoder(w).Encode(task)
}
