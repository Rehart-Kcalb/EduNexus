-- name: AllCategories :many
SELECT
  "name",
  "color"
FROM
  categories;

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
  courses.image,
  users.firstname AS organization_name,
  users.profile as organization_logo
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
    $1,
    $2,
    NOW()
  );

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

-- name: GetCourseLectures :many
select a.* from courses c 
left join modules m on m.course_id = c.id
left join assignments a on a.module_id = m.id
where  c.id = $1 and a.id is not null;

-- name: GetLectureContent :one
select * from assignments
where assignments.id = $1;

-- name: GetCourseDetails :one
SELECT
  c.description,c.id,c.image
FROM
  courses c where c.id = $1;

-- name: NewLecture :exec
insert into assignments (module_id,title,description,content,assignment_type_id)
values ($1,$2,$3,$4,1);

-- name: CheckEnrollment :one
select enrolled_on from enrollments where course_id = $1 and user_id = $2;

-- name: FilterCourses :many
select * from filter($1,$2::bigint[]) limit $3 offset $4;

-- name: GetCategoryId :one
select id from categories where name=$1 limit 1;

-- name: CountCourses :one
select count(title) from filter($1,$2::bigint[]) ;
