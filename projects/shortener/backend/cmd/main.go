package main

import (
	api "backend/delivery/http"
	"backend/delivery/http/middleware"
	"backend/internal/app"
	"backend/internal/store/sqlstore"
	"backend/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("error loading .env file: %s", err)
	}

	db, err := sql.Open(model.DBTypeMySQL, os.Getenv("MYSQL_CONNECTION_STRING"))
	if err != nil {
		log.Printf("failed to connect to mysql %s", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("failed to connect to mysql %s", err)
	}

	sqlStore, err := sqlstore.New(db)
	if err != nil {
		log.Printf("failed to connect to MySQL %s:", err)
	}

	mainApp := app.New(sqlStore)
	server := api.New(mainApp)

	r := mux.NewRouter()
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	r.Use(middleware.LoggingMiddleware)
	server.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server starting on port %s", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), corsHandler(r))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
