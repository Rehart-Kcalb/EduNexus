package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetCourseLectures(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		course_name := r.PathValue("course_name")
		course_id, err := DB.GetCourseId(context.Background(), course_name)
		if err != nil {
			// TODO
			return
		}
		lectures, err := DB.GetCourseLectures(context.Background(), course_id)
		if err != nil {
			// TODO
			return
		}
		types.NewJsonResponse(struct {
			Lectures any `json:"lectures"`
		}{lectures}, http.StatusOK).Respond(w)
	}
}
