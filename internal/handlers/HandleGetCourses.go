package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetCourses(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			return
		}
		Limit_str := r.FormValue("perPage")
		Offset_str := r.FormValue("page")
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
		courses, err := DB.GetCourses(context.Background(), db.GetCoursesParams{Limit: int32(limit_num), Offset: int32(Offset_num)})
		if err != nil {
			types.NewJsonResponse("Problem with DB", http.StatusInternalServerError).Respond(w)
			return
		}
		types.NewJsonResponse(struct {
			Courses any `json:"courses"`
		}{courses}, http.StatusOK).Respond(w)
	}
}
