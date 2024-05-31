package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
)

func HandleCreateAssignment(DB *db.Queries) http.HandlerFunc {
	type AssignmentPost struct {
		ModuleId         int64  `json:"module_id"`
		Title            string `json:"title"`
		Description      string `json:"description"`
		Content          string `json:"content"`
		AssignmentTypeId int64  `json:"assignment_type_id"`
	}
	type AssignmentParam struct {
		AssignmentPost
		CourseID int64 `json:"course_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("wer")
		course_id, err := DB.GetCourseId(context.Background(), r.PathValue("course_name"))
		if err != nil {
			log.Println(err)
		}
		var AssignPost AssignmentPost
		err = json.NewDecoder(r.Body).Decode(&AssignPost)
		if err != nil {
			log.Println(err)
		}
		Assign_Param := AssignmentParam{AssignmentPost: AssignPost, CourseID: course_id}
		_ = Assign_Param
		switch Assign_Param.AssignmentTypeId {
		case 2:
			if !CheckQuiz(Assign_Param.Content) {
				// TODO: SOMETHING
			}
		case 3:
			if !CheckCode(Assign_Param.Content) {
				// TODO:SOMETHING
			}
		case 4:
			if !CheckMatching(Assign_Param.Content) {
				// TODO: SOMETHING
			}
		case 5:
			if !CheckFillIn(Assign_Param.Content) {
				// TODO: Something
			}
		case 6:
			if !CheckFree(Assign_Param.Content) {
				// TODO: SOMETHING
			}
		case 7:
			if !CheckNumber(Assign_Param.Content) {
				// TODO: SOMETHING
			}
		}
		err = DB.CreateAssignment(context.Background(), db.CreateAssignmentParams{CourseID: Assign_Param.CourseID, ModuleID: Assign_Param.ModuleId, Content: []byte(Assign_Param.Content), Description: Assign_Param.Description, Title: Assign_Param.Title, AssignmentTypeID: Assign_Param.AssignmentTypeId})
		if err != nil {
			log.Println(err)
		}
	}
}

func CheckQuiz(quiz_def string) bool {
	return true
}

func CheckCode(code_def string) bool {
	return true
}

func CheckMatching(match_def string) bool {
	return true
}

func CheckFillIn(fill_def string) bool {
	return true
}

func CheckSql(sql_def string) bool {
	return true
}

func CheckSort(sort_def string) bool {
	return true
}

func CheckFree(free_answer string) bool {
	return true
}

func CheckNumber(number_def string) bool {
	return true
}
