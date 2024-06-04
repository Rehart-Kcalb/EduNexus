package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
)

func HandleCreateLecture(DB *db.Queries) http.HandlerFunc {
	type Lecture struct {
		ModuleId    int64  `json:"module_id"`
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
		err = DB.NewLecture(context.Background(), db.NewLectureParams{CourseID: course_id, ModuleID: post_data.ModuleId, Content: []byte(post_data.Content), Description: post_data.Description})
		if err != nil {
			return
		}
	}
}
