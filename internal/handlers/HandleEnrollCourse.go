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
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Нет такого курса"}, http.StatusBadRequest).Respond(w)
			return
		}
		// This check if user already enrolled
		if _, err := DB.CheckEnrollment(context.Background(), db.CheckEnrollmentParams{UserID: user_id, CourseID: course_id}); err == nil {
			// TODO: SOMETHING USER already enroll
			types.NewJsonResponse(struct {
				Status string `json:"status"`
			}{"Уже записались на курс"}, http.StatusOK).Respond(w)
			return
		}
		err = DB.EnrollIntoCourse(context.Background(), db.EnrollIntoCourseParams{CourseID: course_id, UserID: user_id})
		if err != nil {
			// TODO: MAKE proper error handling
			log.Println(err)
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Проблемы с БД"}, http.StatusInternalServerError).Respond(w)
			return
		}
		types.NewJsonResponse(struct {
			Status string `json:"status"`
		}{"success"}, http.StatusOK).Respond(w)
	}
}
