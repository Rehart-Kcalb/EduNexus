-- name: AllCategories :many
SELECT
  "name",
  "color"
FROM
  categories;

-- name: GetCategoryCourses :many
SELECT
  DISTINCTs (courses.title)
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
  courses.title,
  users.firstname AS organization_name
FROM
  courses
  LEFT JOIN users ON users.id = courses.course_provider
WHERE
  users.id = courses.course_provider
  AND courses.id IN (
    SELECT
      courses.id
    FROM
      courses
      INNER JOIN enrollments ON enrollments.course_id = courses.id
    WHERE
      enrollments.user_id = $1
  );

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

-- name: GetCourses :many
SELECT
  courses.title,
  users.firstname AS organization_name
FROM
  courses
  LEFT JOIN users ON users.id = courses.course_provider
WHERE
  users.id = courses.course_provider;

-- name: GetCourseModules :many
SELECT
  modules.title
FROM
  modules
  INNER JOIN courses ON courses.id = modules.course_id
WHERE
  courses.title = $1;

-- name: EnrollIntoCourse :exec
INSERT INTO
  enrollments (course_id, user_id, enrolled_on)
VALUES
  (
    (
      SELECT
        id
      FROM
        courses
      WHERE
        title = $1
      LIMIT
        1
    ),
    $2,
    NOW()
  );

-- name: GetCourseTeachers :many 
SELECT
  u.firstname,u.surname
FROM
  courses
  INNER JOIN course_teachers ct ON ct.course_id = courses.id 
  inner join users u on u.id = ct.user_id
  where ct.course_id = $1;

-- name: GetCourseEnrolledAmount :one
select count(id) from enrollments where enrollments.course_id = $1;

-- name: GetCourseId :one 
select id from courses where title = $1 limit 1;

-- name: GetCourseDetails :one
SELECT
  c.description,c.id
FROM
  courses c where c.id = $1;
