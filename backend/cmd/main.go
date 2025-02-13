package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"infosphere-backend/internal/database"
	"infosphere-backend/internal/handlers"
	"infosphere-backend/pkg/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	databaseUrl := "postgres://nebula@localhost:5432/webagg?sslmode=disable"
	conn, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer conn.Close()

	queries := database.New(conn)

	apiMux := http.NewServeMux()

	apiMux.Handle("POST /api/login", handlers.LoginHandler(queries))
	apiMux.Handle("POST /api/signup", handlers.SignupHandler(queries))

	fmt.Println("HTTP server listening on localhost:8080")
	http.ListenAndServe(":8080", utils.CORSMiddleware(apiMux))
}
