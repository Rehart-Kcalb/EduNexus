package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetAssignment(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assignment_id_str := r.PathValue("assignment_id")
		assignment_id, err := strconv.Atoi(assignment_id_str)
		if err != nil {
			return
		}
		assignment, err := DB.GetAssignmentById(context.Background(), int64(assignment_id))

		types.NewJsonResponse(struct {
			Assignment any `json:"assignment"`
		}{assignment}, http.StatusOK).Respond(w)
	}
}
