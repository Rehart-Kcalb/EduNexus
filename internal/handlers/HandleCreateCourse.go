package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/utils"
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
		if len(post_data.Title) < 5 {
			types.NewJsonResponse(struct {
				ErrorMessage string `json:"error_message"`
			}{"Title so short"}, http.StatusBadRequest).Respond(w)
			return
		}
		filePath, err := utils.SaveBase64ToFile(post_data.Image, "storage")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error saving file: %v", err), http.StatusInternalServerError)
			return
		}
		log.Println(filePath)
		log.Println(pgtype.Text{String: filePath, Valid: true})
		course_id, err := DB.CreateCourse(context.Background(), db.CreateCourseParams{Title: post_data.Title, Description: post_data.Description, Image: pgtype.Text{String: filePath, Valid: true}, CourseProvider: user_id})
		if err != nil {
			log.Println(err)
			return
		}
		err = DB.AddTeacher(context.Background(), db.AddTeacherParams{UserID: user_id, CourseID: course_id})

		if err != nil {
			log.Println(err)
		}

		for _, name := range post_data.Categories {
			category_id, err := DB.GetCategoryId(context.Background(), name)
			if err != nil {
				log.Println(err)
				return
			}
			DB.AddCategoryCourse(context.Background(), db.AddCategoryCourseParams{CourseID: course_id, CategoryID: category_id})
		}
		types.NewJsonResponse(struct {
			Status string `json:"status"`
		}{"success"}, http.StatusOK).Respond(w)
	}
}
