package handlers

import (
	"context"
	"log"
	"math"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/utils"
)

func HandleGetCategoryCourses(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		PagParams := utils.GetPaginationParams(r.URL.Query())
		category_name := r.PathValue("category_name")
		if category_name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		course_id, err := DB.GetCategoryId(context.Background(), category_name)
		if err != nil {
			types.NewJsonResponse(struct {
				ErrorMessage string `json:"error_message"`
			}{"Нет такой категории"}, http.StatusOK).Respond(w)
			return
		}
		log.Println(course_id)
		courses, err := DB.FilterCourses(context.Background(), db.FilterCoursesParams{Limit: PagParams.Limit, Offset: PagParams.Offset, Column2: []int64{course_id}})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(courses)
		count, err := DB.CountCourses(context.Background(), db.CountCoursesParams{Column2: []int64{course_id}})
		if err != nil {
			types.NewJsonResponse(struct {
				Courses any   `json:"courses"`
				Count   int64 `json:"pages"`
			}{[]any{}, 1}, http.StatusOK).Respond(w)
			return
		}
		pages := int64(math.Ceil(float64(count) / float64(PagParams.Limit)))
		if pages == 0 {
			pages = 1
		}
		types.NewJsonResponse(struct {
			Courses any   `json:"courses"`
			Count   int64 `json:"pages"`
		}{courses, pages}, http.StatusOK).Respond(w)
	}
}
