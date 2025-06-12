package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
	"github.com/thejunghare/taskManager"
)

func main() {
	// Todo: using sqlc
	err := godotenv.Load()
	if err != nil {
		fmt.Println("godotenv error: ", err)
	}
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("db conn error: ", err)
	}

	defer conn.Close(ctx)

	queries := taskManager.New(conn)

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		task, err := queries.GetAllTask(ctx)
		if err != nil {
			fmt.Println("Error while fetching tasks: ", err)
			return
		}
		//fmt.Println(task)

		err = Tasks(task).Render(ctx, w)
		if err != nil {
			fmt.Println("Failed to return the templ: ", err)

		}
	})

	http.HandleFunc("/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		err := queries.UpdateStatus(ctx, taskManager.UpdateStatusParams{
			ID:     18,
			Status: true,
		})

		if err != nil {
			log.Printf("Error while updating  %v", err)
		}
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		createId, err := queries.CreateTaskAndReturnId(ctx, taskManager.CreateTaskAndReturnIdParams{
			Name: "testing",
			CreatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
		})

		if err != nil {
			log.Printf("Error while updating %v", err)
			http.Error(w, "Failed to create task", http.StatusInternalServerError)
			return
		}

		log.Printf("Inserted ID %v", createId)
		fmt.Fprintf(w, "Created task with ID %d", createId)
	})

	http.HandleFunc("/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		err := queries.Delete(ctx, 16)
		if err != nil {
			fmt.Println("Error while deleting", err)
		}
	})

	// Todo: get single task

	log.Fatal(http.ListenAndServe(":8080", nil))

	/* tasks, err := queries.GetAllTask(ctx)
	if err != nil {
		fmt.Println("Error while fetching tasks: ", err)
	} */

	// fmt.Printf("Type of x: %T\n", task)
	//fmt.Printf("Type of x: %T\n", tasks)
}
