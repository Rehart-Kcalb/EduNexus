package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewServer(logger *log.Logger, DB *db.Queries) http.Handler {
	mux := http.NewServeMux()
	AddRoutes(mux, DB)
	var handler http.Handler = mux
	handler = middleware.CORS(mux)
	return handler
}

// main function serves as the entry point of our application
func main() {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	pool := ConnectToDB(connString)
	log.Println(connString)
	DB := db.New(pool)
	server := NewServer(log.Default(), DB)

	// Start an HTTP server serving requests using the prepared middleware-enforced Mux with loaded handlers
	log.Fatal(http.ListenAndServe(":8080", server)) // Serve on port 8080; log errors if server fails to start
}

// ConnectToDB establishes a connection to an external PostgreSQL database using the provided connection string.
func ConnectToDB(conn_string string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), conn_string)
	if err != nil {
		log.Fatal("Problem with database connection :" + err.Error())
	}
	//fmt.Println(conn_string)
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Problem with database connection :" + err.Error())
	}
	log.Println("Connection to the database was successful.")
	return pool
}
