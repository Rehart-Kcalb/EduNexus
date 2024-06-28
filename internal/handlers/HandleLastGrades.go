package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleLastGrades(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("id").(int64)
		course_name := r.PathValue("course_name")
		grades, err := DB.GetLastGradesByCourse(context.Background(), db.GetLastGradesByCourseParams{UserID: user_id, Title: course_name})
		if err != nil {
			types.NewJsonResponse(struct {
				Grades any `json:"grades"`
			}{grades}, http.StatusOK).Respond(w)
			return
		}
		types.NewJsonResponse(struct {
			Grades any `json:"grades"`
		}{grades}, http.StatusOK).Respond(w)
	}
}
