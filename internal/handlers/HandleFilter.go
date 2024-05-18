package handlers

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/utils"
)

func HandleFilter(DB *db.Queries) http.HandlerFunc {
	type PostData struct {
		Title      string   `json:"title"`
		Categories []string `json:"categories"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		PagParams := utils.GetPaginationParams(r.URL.Query())
		var post_data PostData
		err := json.NewDecoder(r.Body).Decode(&post_data)
		if err != nil {
			log.Println("Error while decoding" + err.Error())
		}
		var course_ids []int64 = make([]int64, 10)
		for _, name := range post_data.Categories {
			course_id, err := DB.GetCategoryId(context.Background(), name)
			if err != nil {
				log.Println(err)
				return
			}
			course_ids = append(course_ids, course_id)
		}
		courses, err := DB.FilterCourses(context.Background(), db.FilterCoursesParams{TitleParam: post_data.Title, Column2: course_ids, Limit: int32(PagParams.Limit), Offset: int32(PagParams.Offset)})
		if err != nil {
			log.Println(err)
			return
		}
		count, err := DB.CountCourses(context.Background(), db.CountCoursesParams{TitleParam: post_data.Title, Column2: course_ids})
		if err != nil {
			// TODO: error handler
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
