package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleCreateLecture(DB *db.Queries) http.HandlerFunc {
	type Lecture struct {
		Title       string `json:"title"`
		ModuleName  string `json:"module_name"`
		Description string `json:"description"`
		Content     string `json:"content"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		course_id, err := DB.GetCourseId(context.Background(), r.PathValue("course_name"))
		if err != nil {
			return
		}
		var post_data Lecture
		err = json.NewDecoder(r.Body).Decode(&post_data)
		if err != nil {
			return
		}

		module_id, err := DB.GetModuleId(context.Background(), db.GetModuleIdParams{Title: post_data.ModuleName, CourseID: course_id})

		err = DB.NewLecture(context.Background(), db.NewLectureParams{CourseID: course_id, ModuleID: module_id, Content: []byte(fmt.Sprintf("{\"content\":\"%s\"}", post_data.Content)), Description: post_data.Description, Title: post_data.Title})
		if err != nil {
			return
		}
		types.NewJsonResponse(struct {
			Status string `json:"status"`
		}{"success"}, http.StatusOK).Respond(w)
	}
}
