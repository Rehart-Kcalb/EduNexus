package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleEnrollCourse(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("id").(int64)
		course_name := r.PathValue("course_name")
		_ = user_id
		course_id, err := DB.GetCourseId(context.Background(), course_name)
		if err != nil {
			// TODO: HANDLE ERROR
		}
		// TODO: CHECK for enrollment
		if _, err := DB.CheckEnrollment(context.Background(), db.CheckEnrollmentParams{UserID: user_id, CourseID: course_id}); err == nil {
			// TODO: SOMETHING USER already enroll
			log.Println("ENROLLED")
			return
		}
		err = DB.EnrollIntoCourse(context.Background(), db.EnrollIntoCourseParams{CourseID: course_id, UserID: user_id})
		if err != nil {
			// TODO: MAKE proper error handling
			log.Println(err)
			return
		}
		types.NewJsonResponse(struct {
			Status string `json:"status"`
		}{"success"}, http.StatusOK).Respond(w)
	}
}
