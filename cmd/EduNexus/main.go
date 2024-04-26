package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/handlers"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	conn, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		fmt.Printf("Problem with database connection :%v", err)
	}

	mux := http.NewServeMux()
	DB := db.New(conn)

	mux.HandleFunc("GET /api/categories/", handlers.HandleGetAllCategories(DB))
	mux.HandleFunc("GET /api/categories/{category_name}", handlers.HandleGetCategoryCourses(DB))
	mux.HandleFunc("GET /api/courses", middleware.Auth(handlers.HandleGetMyCourses(DB)))
	mux.HandleFunc("POST /api/login", handlers.HandleLogin(DB))
	mux.HandleFunc("POST /api/register", handlers.HandleRegister(DB))
	http.ListenAndServe(":8080", mux)
}
