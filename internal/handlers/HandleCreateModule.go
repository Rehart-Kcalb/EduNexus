package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleCreateModule(DB *db.Queries) http.HandlerFunc {
	type Title struct {
		Title string `json:"title"`
	}
	return func(w http.ResponseWriter, r *http.Request) {

		var post_data Title
		err := json.NewDecoder(r.Body).Decode(&post_data)
		if err != nil {
			log.Println(err)
		}
		course_name := r.PathValue("course_name")
		course_id, err := DB.GetCourseId(context.Background(), course_name)
		if err != nil {
			return
		}
		_ = course_id
		module_id, err := DB.CreateModule(context.Background(), db.CreateModuleParams{CourseID: course_id, Title: post_data.Title})
		if err != nil {
			return
		}

		types.NewJsonResponse(struct {
			ModuleID int64 `json:"module_id"`
		}{module_id}, http.StatusOK).Respond(w)
	}
}
