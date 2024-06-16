package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetModuleProgress(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO:
		user_id := r.Context().Value("id").(int64)
		course_id, err := DB.GetCourseId(context.Background(), r.PathValue("course_name"))
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Нет такого курса"}, http.StatusBadRequest).Respond(w)
			return
		}
		module_id, err := DB.GetModuleId(context.Background(), db.GetModuleIdParams{CourseID: course_id, Title: r.PathValue("module_name")})
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Нет такого модуля"}, http.StatusBadRequest).Respond(w)
			return
		}
		lecture_progress, err := DB.GetReadedLecturesByModule(context.Background(), db.GetReadedLecturesByModuleParams{UserID: user_id, ID: module_id})
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Проблемы с БД"}, http.StatusInternalServerError).Respond(w)
			return
		}
		types.NewJsonResponse(struct {
			Progress any `json:"lecture_progres"`
		}{lecture_progress}, http.StatusOK).Respond(w)
	}
}
