package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleReadLecture(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assignment_id_str := r.PathValue("lecture_id")
		assignment_id, err := strconv.Atoi(assignment_id_str)
		if err != nil {
			return
		}
		log.Println(assignment_id)
		err = DB.MarkAssignmentDone(context.Background(), db.MarkAssignmentDoneParams{AssignmentID: int64(assignment_id), UserID: r.Context().Value("id").(int64)})
		if err != nil {
			log.Println(err)
			return
		}
		types.NewJsonResponse(struct {
			Message string `json:"message"`
		}{"success"}, http.StatusOK)
	}
}
