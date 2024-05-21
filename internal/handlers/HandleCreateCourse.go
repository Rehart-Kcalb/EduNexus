package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
)

func HandleCreateCourse(DB *db.Queries) http.HandlerFunc {
	type Course struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Image       string   `json:"image"`
		Categories  []string `json:"categories"`
		Modules     []string `json:"modules"`
		Teachers    []string `json:"teachers"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: Make handler
		var post_data Course
		err := json.NewDecoder(r.Body).Decode(&post_data)
		if err != nil {
			if errors.Is(err, io.EOF) {
				// TODO:Make error to send data
			} else {
				log.Println("Error while decoding" + err.Error())
			}
		}

		var category_ids []int64 = make([]int64, 0)
		for _, name := range post_data.Categories {
			category_id, err := DB.GetCategoryId(context.Background(), name)
			if err != nil {
				log.Println(err)
				return
			}
			category_ids = append(category_ids, category_id)
		}
		//DB.CreateCourse
	}
}
