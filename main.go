package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/thejunghare/go-threads-app/db"
	"log"
	"net/http"
	"os"
)

func main() {
	checkError(godotenv.Load())

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	checkError(err)

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, ctx)

	queries := db.New(conn)

	http.HandleFunc("/get-thread/{id}", func(w http.ResponseWriter, r *http.Request){
		thread, err := queries.GetThread(ctx, 1)
		checkError(err)
		log.Println(thread)
	})

	http.HandleFunc("get-post/{id}", func(w http.ResponseWriter, r *http.Request){
		post, err := queries.GetPost(ctx, 1)
		checkError(err)
		log.Println(post)
	})

	checkError(http.ListenAndServe(":3000", nil))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
