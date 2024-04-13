package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/handlers"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	conn, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Printf("Problem with database connection :%v", err)
	}

	mux := http.NewServeMux()
	DB := db.New(conn)

	mux.HandleFunc("GET /api/categories", handlers.HandleGetAllCategories(DB))
}
