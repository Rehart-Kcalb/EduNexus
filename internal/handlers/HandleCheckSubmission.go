package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleCheckSubmission(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Make this work
		course_name := r.PathValue("course_name")
		_ = course_name
		assignment_id_str := r.PathValue("assignment_id")
		assignment_id_num, err := strconv.Atoi(assignment_id_str)
		if err != nil {
			// TODO: Make error handler
			types.NewJsonResponse(struct {
				Message string `json:"error message"`
			}{"Invalid assignment id"}, http.StatusBadRequest).Respond(w)
			return
		}
		_ = assignment_id_num
		assignment, err := DB.GetAssignmentById(context.Background(), int64(assignment_id_num))
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"error message"`
			}{"Failed to get assignment"}, http.StatusBadRequest).Respond(w)
			return
		}
		log.Println(assignment)
		switch assignment.AssignmentTypeID {
		case 2:
			// Quiz
			handleQuizSubmission(w, r, assignment, DB)
		case 3:
			// CODE
			handleCodeSubmission(w, r, assignment, DB)
		case 4:
			// Matching
			handleMatchingSubmission(w, r, assignment, DB)
		case 5:
			// Fill In
			handleFillInSubmission(w, r, assignment, DB)
		case 6:
			// sql code
			handleSQLSubmission(w, r, assignment, DB)
		case 7:
			// sort
			handleSortSubmission(w, r, assignment, DB)
		case 8:
			// free answer
			handleFreeAnswerSubmission(w, r, assignment, DB)
		case 9:
			// Number
			handleNumberSubmission(w, r, assignment, DB)
		}
	}
}
