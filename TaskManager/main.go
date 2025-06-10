package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
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

	// Todo: fix this instead component templ
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		task, err := queries.GetAllTask(ctx)
		if err != nil {
			fmt.Println("tasks error: ", err)
			return
		}

		err = RenderTaskPage(task).Render(ctx, w)
		if err != nil {
			fmt.Println("failed to return the templ: ", err)

		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	tasks, err := queries.GetAllTask(ctx)
	if err != nil {
		fmt.Println("tasks error: ", err)
	}

	// fmt.Printf("Type of x: %T\n", task)
	fmt.Printf("Type of x: %T\n", tasks)
}
