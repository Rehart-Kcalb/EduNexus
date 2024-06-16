package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetLectureContent(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		course_name := r.PathValue("course_name")
		_ = course_name
		lecture_id_str := r.PathValue("lecture_id")
		lecture_id, err := strconv.Atoi(lecture_id_str)
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Айди лекции должен быть натуральным числом"}, http.StatusBadRequest).Respond(w)
			// TODO
			return
		}
		lecture_content, err := DB.GetLectureContent(context.Background(), int64(lecture_id))
		if err != nil {
			// TODO
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Проблемы с БД"}, http.StatusInternalServerError).Respond(w)
			return
		}
		types.NewJsonResponse(struct {
			LectureContent any `json:"content"`
		}{lecture_content}, http.StatusOK).Respond(w)
	}
}
