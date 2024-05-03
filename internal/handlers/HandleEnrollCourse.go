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
		_ = course_name
		err := DB.EnrollIntoCourse(context.Background(), db.EnrollIntoCourseParams{Title: course_name, UserID: user_id})
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
