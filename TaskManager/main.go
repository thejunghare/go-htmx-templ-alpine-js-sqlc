package main

import (
	"log"
	"net/http"

	"github.com/thejunghare/controllers"
	"github.com/thejunghare/db"
)

func main() {
	// connecting to db
	db.InitDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tasks, err := controllers.GetAllTasksTempl()
		if err != nil {
			log.Printf("Failed to get tasks: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		Hello("Prasad", tasks).Render(r.Context(), w)
	})

	// get tasks
	http.HandleFunc("/tasks", controllers.GetAllTasks)

	// add task
	http.HandleFunc("/add", controllers.AddTasks)

	http.ListenAndServe(":8080", nil)
}
