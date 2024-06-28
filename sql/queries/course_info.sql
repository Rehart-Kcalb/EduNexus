-- name: GetCourseModules :many
SELECT
  modules.title
FROM
  modules
  INNER JOIN courses ON courses.id = modules.course_id
WHERE
  courses.title = $1;

-- name: GetCourseTeachers :many 
SELECT
  u.firstname,u.surname,u.profile
FROM
  courses
  INNER JOIN course_teachers ct ON ct.course_id = courses.id 
  inner join users u on u.id = ct.user_id
  where ct.course_id = $1 AND u.user_role_id = 1;

-- name: GetCourseEnrolledAmount :one
select count(id) from enrollments where enrollments.course_id = $1;

-- name: GetCourseId :one 
select id from courses where title = $1 limit 1;

-- name: GetCourseDetails :one
SELECT
  c.description,c.id,c.image
FROM
  courses c where c.id = $1;

