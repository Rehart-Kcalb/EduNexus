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

// main function serves as the entry point of our application
func main() {
	conn_string := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	pool := ConnectToDB(conn_string)
	log.Println(conn_string)

	Migrate(conn_string) // Run database migrations against the connected PostgreSQL instance

	mux := http.NewServeMux()

	Cors := cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowedHeaders: []string{"Content-Type", "Authorization"}, AllowedMethods: []string{http.MethodGet, http.MethodOptions, http.MethodHead, http.MethodPost}})
	handler := Cors.Handler(mux)

	// Initialize the database connection pool and handlers
	DB := db.New(pool)
	LoadMuxWithHandlers(mux, DB)

	// Start an HTTP server serving requests using the prepared middleware-enforced Mux with loaded handlers
	log.Fatal(http.ListenAndServe(":8080", handler)) // Serve on port 8080; log errors if server fails to start
}

// The Migrate function initializes the migrations tool and executes database schema changes specified in migration files.
func Migrate(conn_string string) {
	m, err := migrate.New("file://sql/migrations", conn_string)
	if err != nil {
		log.Fatal("Creating migrations was failed: " + err.Error()) // Logs an error if the creation of migration files fails
	}

	// The Up method applies all pending database migrations specified in the migrations file(s).
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Migration up was failed: " + err.Error()) // Logs an error if a migration fails to apply
	} else {
		// If no migrations are needed or all migrations were applied successfully, log the success message.
		log.Println("Migration was succesful")
	}
}

func LoadMuxWithHandlers(m *http.ServeMux, DB *db.Queries) {
	// **User Management**
	// * Handles user login requests (POST /api/login)
	m.HandleFunc("POST /api/login", handlers.HandleLogin(DB))
	// * Handles user registration requests (POST /api/register)
	m.HandleFunc("POST /api/register", handlers.HandleRegister(DB))
	// * Show User profile
	m.HandleFunc("GET /api/profile", middleware.Auth(handlers.HandleGetProfileInfo(DB)))
	m.HandleFunc("POST /api/profile", middleware.Auth(handlers.HandleUpdateProfileInfo(DB)))

	// ** Assets
	m.HandleFunc("GET /file/{path...}", handlers.HandleGetAsset())

	// **Course Management**
	// * Gets all available categories (GET /api/categories/)
	m.HandleFunc("GET /api/categories/", handlers.HandleGetAllCategories(DB))
	// * Gets courses within a specific category (GET /api/categories/{category_name})
	m.HandleFunc("GET /api/categories/{category_name}", handlers.HandleGetCategoryCourses(DB))
	// * Gets all available courses (GET /api/courses/)
	m.HandleFunc("GET /api/courses/", handlers.HandleGetCourses(DB))
	// * Gets information about a specific course (GET /api/courses/{course_name}/)
	m.HandleFunc("GET /api/courses/{course_name}/", handlers.HandleGetCourseInfo(DB))
	// * Gets modules for a specific course (GET /api/courses/{course_name}/modules)
	m.HandleFunc("GET /api/courses/{course_name}/modules", (handlers.HandleGetCourseModules(DB)))
	// * Filters courses based on specific criteria (POST /api/filter) - Likely requires additional details in the request body
	m.HandleFunc("POST /api/filter", (handlers.HandleFilter(DB)))

	// **Course Enrollment (requires authentication)**
	// * Enrolls a user in a course (POST /api/courses/{course_name})
	m.HandleFunc("POST /api/courses/{course_name}", middleware.Auth(handlers.HandleEnrollCourse(DB)))

	// **Learning Management (requires authentication)**
	// * Gets an assignments for a specific course (GET /api/learning/{course_name}/assignments)
	m.HandleFunc("GET /api/learning/{course_name}/assignments", handlers.HandleGetAssignments(DB))
	// * Gets an assignment within a course (GET /api/learning/{course_name}/assignment/{assignment_id})
	m.HandleFunc("GET /api/learning/{course_name}/assignments/{assignment_id}", middleware.Auth(handlers.HandleGetAssignment(DB)))
	// * Gets the content of a submitted assignment (GET /api/learning/{course_name}/{assignment_id}/{submission_id})
	m.HandleFunc("GET /api/learning/{course_name}/assignments/{assignment_id}/{submission_id}", middleware.Auth(handlers.HandleGetContentOfSubmission(DB)))
	// * Checks a assignment (POST /api/learning/{course_name}/{assignment_id})
	m.HandleFunc("POST /api/learning/{course_name}/assignments/{assignment_id}", middleware.Auth(handlers.HandleCheckSubmission(DB)))
	// * Gets all courses a user is enrolled in (GET /api/learning/)
	m.HandleFunc("GET /api/learning/", middleware.Auth(handlers.HandleGetMyCourses(DB)))
	// * Gets lectures for a specific course (GET /api/learning/{course_name})
	m.HandleFunc("GET /api/learning/{course_name}/lectures", middleware.Auth(handlers.HandleGetCourseLectures(DB)))
	// * Gets the content for a specific lecture (GET /api/learning/{course_name}/lectures/{lecture_id})
	m.HandleFunc("GET /api/learning/{course_name}/lectures/{lecture_id}", middleware.Auth(handlers.HandleGetLectureContent(DB)))
	// * Get module progress
	m.HandleFunc("GET /api/learning/{course_name}/modules/{module_name}", middleware.Auth(handlers.HandleGetModuleProgress(DB)))
	// *- Mark lecture as read
	m.HandleFunc("POST /api/learning/{course_name}/read/{lecture_id}", middleware.Auth(handlers.HandleReadLecture(DB)))

	// **Teaching functionalities (requires authentication)**
	// * Gets all courses a user is teaching (GET /api/teaching/)
	m.HandleFunc("GET /api/teaching/", middleware.Auth(handlers.HandleGetMyTeached(DB)))
	// * Gets all submissions for a specific course (GET /api/teaching/{course_name})
	m.HandleFunc("GET /api/teaching/{course_name}", middleware.Auth(handlers.HandleGetCourseSubmissions(DB)))
	// * Gets a specific submission for grading (GET /api/teaching/{course_name}/{submission_id})
	m.HandleFunc("GET /api/teaching/{course_name}/{submission_id}", middleware.Auth(handlers.HandleGetSubmissionForGrading(DB)))
	// * Grades a submission (POST /api/teaching/{course_name}/{submission_id})
	m.HandleFunc("POST /api/teaching/{course_name}/{submission_id}", middleware.Auth(handlers.HandleGradeSubmission(DB)))
	// * Creates a new assignment for a course (requires instructor privileges) (POST /api/learning/{course_name})
	m.HandleFunc("POST /api/teachings/{course_name}/assignments", middleware.Auth(handlers.HandleCreateAssignment(DB)))
	// Create course
	m.HandleFunc("POST /api/teachings", middleware.Auth(handlers.HandleCreateCourse(DB)))
	// * Create Module
	m.HandleFunc("POST /api/teaching/{course_name}/modules", middleware.Auth(handlers.HandleCreateModule(DB)))
	// * Create Lecture
	m.HandleFunc("POST /api/teaching/{course_name}/lectures", middleware.Auth(handlers.HandleCreateLecture(DB)))
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
