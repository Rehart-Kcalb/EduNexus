-- name: AllCategories :many
SELECT
  "name",
  "color"
FROM
  categories;

-- name: GetCategoryCourses :many
SELECT DISTINCT
  (courses.title)
FROM
  courses
  INNER JOIN course_categories cc ON courses.id = cc.course_id
  INNER JOIN categories ON cc.category_id = categories.id
WHERE
  categories.name = $1;

-- name: GetPasswordByLogin :one
SELECT
  PASSWORD
FROM
  users
WHERE
  login = $1;

-- name: GetMyCourses :many
SELECT
  courses.title
FROM
  enrollments
  INNER JOIN courses ON enrollments.course_id = courses.id
WHERE
  enrollments.user_id = $1;

-- name: GetClaimsByLogin :one
SELECT
  users.id,
  user_roles.title
FROM
  users
  LEFT JOIN user_roles ON user_roles.id = users.user_role_id
WHERE
  users.login = $1;

-- name: CreateUser :exec
INSERT INTO
  users ("login", "password", "user_role_id")
VALUES
  ($1, $2, 1);
