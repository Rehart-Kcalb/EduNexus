package main

import (
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/handlers"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/middleware"
)

func AddRoutes(m *http.ServeMux, DB *db.Queries) {
	// **User Management**
	// * Handles user login requests (POST /api/login)
	m.HandleFunc("POST /api/login", handlers.HandleLogin(DB))
	// * Handles user registration requests (POST /api/register)
	m.HandleFunc("POST /api/register", handlers.HandleRegister(DB))
	// * Show User profile
	m.HandleFunc("GET /api/profile", middleware.Auth(handlers.HandleGetProfileInfo(DB)))
	m.HandleFunc("GET /api/profile/{user_name}", middleware.Auth(handlers.HandleGetProfileInfo(DB)))
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
	m.HandleFunc("GET /api/popular/", handlers.HandleGetPopularCourses(DB))
	// * Gets information about a specific course (GET /api/courses/{course_name}/)
	m.HandleFunc("GET /api/courses/{course_name}/", handlers.HandleGetCourseInfo(DB))
	// * Gets modules for a specific course (GET /api/courses/{course_name}/modules)
	m.HandleFunc("GET /api/courses/{course_name}/modules", (handlers.HandleGetCourseModules(DB)))
	// * Filters courses based on specific criteria (POST /api/filter) - Likely requires additional details in the request body
	m.HandleFunc("POST /api/filter", (handlers.HandleFilter(DB)))

	// **Course Enrollment (requires authentication)**
	// * Enrolls a user in a course (POST /api/courses/{course_name})
	m.HandleFunc("POST /api/courses/{course_name}", middleware.Auth(handlers.HandleEnrollCourse(DB)))
	m.HandleFunc("POST /api/courses/{course_name}/exit", middleware.Auth(handlers.HandleCourseLeave(DB)))

	// **Learning Management (requires authentication)**
	// * Gets an assignments for a specific course (GET /api/learning/{course_name}/assignments)
	m.HandleFunc("GET /api/learning/{course_name}/assignments", middleware.Auth(handlers.HandleGetAssignments(DB)))
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
	m.HandleFunc("GET /api/learning/{course_name}/last_grades", middleware.Auth(handlers.HandleLastGrades(DB)))

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
