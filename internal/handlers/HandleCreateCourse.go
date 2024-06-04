package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/jackc/pgx/v5/pgtype"
)

func HandleCreateCourse(DB *db.Queries) http.HandlerFunc {
	type Course struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Image       string   `json:"image"`
		Categories  []string `json:"categories"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: Make handler
		var post_data Course
		user_id := r.Context().Value("id").(int64)
		err := json.NewDecoder(r.Body).Decode(&post_data)
		if err != nil {
			if errors.Is(err, io.EOF) {
				// TODO:Make error to send data
			} else {
				log.Println("Error while decoding" + err.Error())
			}
		}
		course_id, err := DB.CreateCourse(context.Background(), db.CreateCourseParams{Title: post_data.Title, Description: post_data.Description, Image: pgtype.Text{String: post_data.Image}, CourseProvider: user_id})
		if err != nil {
			return
		}

		for _, name := range post_data.Categories {
			category_id, err := DB.GetCategoryId(context.Background(), name)
			if err != nil {
				log.Println(err)
				return
			}
			DB.AddCategoryCourse(context.Background(), db.AddCategoryCourseParams{CourseID: course_id, CategoryID: category_id})
		}
	}
}
