package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleFilter(DB *db.Queries) http.HandlerFunc {
	type PostData struct {
		Title      string   `json:"title"`
		Categories []string `json:"categories"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		Limit_str := r.URL.Query().Get("perPage")
		Offset_str := r.URL.Query().Get("page")
		limit_num, err := strconv.Atoi(Limit_str)
		if err != nil {
			limit_num = 8
			//log.Println(err)
		}
		Offset_num, err := strconv.Atoi(Offset_str)
		if err != nil {
			Offset_num = 1
			//log.Println(err)
		}
		Offset_num = (Offset_num - 1) * (limit_num)
		var post_data PostData
		var buff []byte = make([]byte, 1000)
		n, err := r.Body.Read(buff)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(buff[:n], &post_data)
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
		courses, err := DB.FilterCourses(context.Background(), db.FilterCoursesParams{TitleParam: post_data.Title, Column2: course_ids, LimitParam: int32(limit_num), OffsetParam: int32(Offset_num)})
		if err != nil {
			log.Println(err)
			return
		}
		types.NewJsonResponse(struct {
			Courses any `json:"courses"`
		}{courses}, http.StatusOK).Respond(w)
	}
}
