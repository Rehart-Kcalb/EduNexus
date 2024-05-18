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

func HandleGetCourses(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//err := r.ParseForm()
		//if err != nil {
		//log.Println(err)
		//return
		//}
		PagParams := utils.GetPaginationParams(r.URL.Query())
		courses, err := DB.FilterCourses(context.Background(), db.FilterCoursesParams{LimitParam: int32(PagParams.Limit), OffsetParam: int32(PagParams.Offset)})
		if err != nil {
			types.NewJsonResponse("Problem with DB", http.StatusInternalServerError).Respond(w)
			return
		}
		count, err := DB.CountCourses(context.Background(), db.CountCoursesParams{LimitParam: PagParams.Limit, OffsetParam: PagParams.Offset})
		if err != nil {
			log.Println(err)
			return
		}
		pages := int64(math.Round(float64(count) / float64(PagParams.Limit)))
		if pages == 0 {
			pages = 1
		}
		types.NewJsonResponse(struct {
			Courses any   `json:"courses"`
			Count   int64 `json:"pages"`
		}{courses, pages}, http.StatusOK).Respond(w)
	}
}
