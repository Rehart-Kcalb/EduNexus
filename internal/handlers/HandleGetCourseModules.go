package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetCourseModules(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		course_name := r.PathValue("course_name")
		course_modules, err := DB.GetCourseModules(context.Background(), course_name)
		if err != nil {
			log.Println("Problem with database:", err)
			return
			// TODO: Handle error properly
		}
		types.NewJsonResponse(struct {
			Course_modules any `json:"course_modules"`
		}{course_modules}, http.StatusOK).Respond(w)
	}
}
