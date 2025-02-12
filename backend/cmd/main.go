package main

import (
	"fmt"
	"net/http"

	"infosphere-backend/internal/handlers"
	"infosphere-backend/pkg/utils"
)

func main() {

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Reached Test"))
	})

	apiMux.HandleFunc("POST /api/login", handlers.LoginHandler)
	apiMux.HandleFunc("POST /api/signup", handlers.SignupHandler)
	fmt.Println("HTTP server listening on localhost:8080")
	http.ListenAndServe(":8080", utils.CORSMiddleware(apiMux))
}
