-- name: GetModuleId :one
select id from modules where course_id = $1 and title = $2;

