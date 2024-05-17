package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/handlers"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/middleware"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
)

func main() {
	conn_string := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	pool := ConnectToDB(conn_string)

	Migrate(conn_string)

	mux := http.NewServeMux()

	Cors := cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowedHeaders: []string{"Content-Type", "Authorization"}, AllowedMethods: []string{http.MethodGet, http.MethodOptions, http.MethodHead, http.MethodPost}})
	handler := Cors.Handler(mux)
	DB := db.New(pool)

	LoadMuxWithHandlers(mux, DB)

	http.ListenAndServe(":8080", handler)
}

func Migrate(conn_string string) {
	m, err := migrate.New("file://sql/migrations", conn_string)
	if err != nil {
		log.Fatal("Creating migrations was failed: " + err.Error())
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Migration up was failed : " + err.Error())
	}
	log.Println("Migration was succesful")

}

func LoadMuxWithHandlers(m *http.ServeMux, DB *db.Queries) {

	m.HandleFunc("POST /api/login", handlers.HandleLogin(DB))
	m.HandleFunc("POST /api/register", handlers.HandleRegister(DB))

	m.HandleFunc("GET /api/categories/", handlers.HandleGetAllCategories(DB))
	m.HandleFunc("GET /api/categories/{category_name}", handlers.HandleGetCategoryCourses(DB))
	m.HandleFunc("GET /api/courses/", handlers.HandleGetCourses(DB))
	m.HandleFunc("GET /api/courses/{course_name}/", handlers.HandleGetCourseInfo(DB))
	m.HandleFunc("GET /api/courses/{course_name}/modules", (handlers.HandleGetCourseModules(DB)))

	m.HandleFunc("POST /api/courses/{course_name}", middleware.Auth(handlers.HandleEnrollCourse(DB)))

	m.HandleFunc("POST /api/courses", middleware.Auth(handlers.HandleCreateCourse(DB)))
	m.HandleFunc("GET /api/learning/", middleware.Auth(handlers.HandleGetMyCourses(DB)))
	m.HandleFunc("GET /api/learning/{course_name}/", middleware.Auth(handlers.HandleGetCourseLectures(DB)))
	m.HandleFunc("GET /api/learning/{course_name}/{lecture_id}", middleware.Auth(handlers.HandleGetLectureContent(DB)))
}

func ConnectToDB(conn_string string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), conn_string)
	if err != nil {
		log.Fatal("Problem with database connection :" + err.Error())
	}
	//fmt.Println(conn_string)
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Problem with database connection :" + err.Error())
	}
	log.Printf("Connection to database was succesful\n")
	return pool
}
