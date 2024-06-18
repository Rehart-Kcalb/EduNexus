package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleCourseLeave(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		course_name := r.PathValue("course_name")
		course_id, err := DB.GetCourseId(context.Background(), course_name)
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Нет такого курса"}, http.StatusBadRequest).Respond(w)
			return
		}
		err = DB.DropCourse(context.Background(), db.DropCourseParams{UserID: r.Context().Value("id").(int64), CourseID: course_id})
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Ошибка в БД"}, http.StatusInternalServerError).Respond(w)
			return
		}

		types.NewJsonResponse(struct {
			Message string `json:"message"`
		}{"success"}, http.StatusOK).Respond(w)
	}
}
