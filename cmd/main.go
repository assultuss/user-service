package main

import (
	"log"
	"net/http"
	"user-service/api"
	"user-service/db"
	"user-service/service"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.AutoMigrate()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	h := api.NewHandler(&service.UserService{})

	router := initializeRoutes(h)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening ...")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func initializeRoutes(h *api.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/users", h.GetUsersHandler)
	mux.HandleFunc("GET /api/users/filter", h.FilterUsersHandler)
	mux.HandleFunc("POST /api/users", h.CreateUserHandler)
	return mux
}
