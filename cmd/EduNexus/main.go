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
	"github.com/rs/cors"
)

func main() {
	conn, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		fmt.Printf("Problem with database connection :%v", err)
	}

	mux := http.NewServeMux()

	Cors := cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowedHeaders: []string{"Content-Type", "Authorization"}, AllowedMethods: []string{http.MethodGet, http.MethodOptions, http.MethodHead, http.MethodPost}})
	handler := Cors.Handler(mux)
	DB := db.New(conn)

	mux.HandleFunc("GET /api/categories/", handlers.HandleGetAllCategories(DB))
	mux.HandleFunc("GET /api/categories/{category_name}", handlers.HandleGetCategoryCourses(DB))
	mux.HandleFunc("GET /api/learning/", middleware.Auth(handlers.HandleGetMyCourses(DB)))
	mux.HandleFunc("POST /api/login", handlers.HandleLogin(DB))
	mux.HandleFunc("POST /api/register", handlers.HandleRegister(DB))
	mux.HandleFunc("GET /api/courses/", handlers.HandleGetCourses(DB))
	mux.HandleFunc("GET /api/courses/{course_name}/", handlers.HandleGetCourseInfo(DB))
	mux.HandleFunc("GET /api/courses/{course_name}/modules", (handlers.HandleGetCourseModules(DB)))
	mux.HandleFunc("POST /api/courses/{course_name}", middleware.Auth(handlers.HandleEnrollCourse(DB)))
	http.ListenAndServe(":8080", handler)
}
