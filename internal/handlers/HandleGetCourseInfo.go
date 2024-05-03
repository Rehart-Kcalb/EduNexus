package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetCourseInfo(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		course_name := r.PathValue("course_name")
		course_id, err := DB.GetCourseId(context.Background(), course_name)
		if err != nil {
			log.Println(err)
			return
			// TODO: Do Actually something with errors
		}
		teachers, err := DB.GetCourseTeachers(context.Background(), course_id)
		if err != nil {
			log.Println(err)
			return
			// TODO: Do Actually something with errors
		}
		_ = teachers

		enrolled, err := DB.GetCourseEnrolledAmount(context.Background(), course_id)
		if err != nil {
			log.Println(err)
			return
			// TODO: Do Actually something with errors
		}
		modules, err := DB.GetCourseModules(context.Background(), course_name)
		if err != nil {
			log.Println(err)
			return
			// TODO: Do Actually something with errors
		}
		details, err := DB.GetCourseDetails(context.Background(), course_id)
		if err != nil {
			log.Println(err)
			return
			// TODO: Do Actually something with errors
		}
		types.NewJsonResponse(struct {
			Modules  any   `json:"modules"`
			Teacher  any   `json:"teachers"`
			Details  any   `json:"details"`
			Enrolled int64 `json:"enrolled"`
		}{modules, teachers, details, enrolled}, http.StatusOK).Respond(w)
	}
}
