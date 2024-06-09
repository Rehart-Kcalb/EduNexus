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
select a.title, a.id  as assignment_id from courses c 
left join modules m on m.course_id = c.id
left join assignments a on a.module_id = m.id
where  c.id = $1 and a.id is not null and a.assignment_type_id = 1;

-- name: GetLectureContent :one
select title, content::text from assignments
where assignments.id = $1;

-- name: GetCourseDetails :one
SELECT
  c.description,c.id,c.image
FROM
  courses c where c.id = $1;

-- name: NewLecture :exec
insert into assignments (module_id,course_id,title,description,content,assignment_type_id)
values ($1,$2,$3,$4,$5,1);

-- name: CheckEnrollment :one
select enrolled_on from enrollments where course_id = $1 and user_id = $2;

-- name: FilterCourses :many
--select filter.title, image,organization_name, organization_logo from filter($1,$2::bigint[]) limit $3 offset $4;
select * from filter($1,$2::bigint[]) limit $3 offset $4;

-- name: GetCategoryId :one
select id from categories where name=$1 limit 1;

-- name: CountCourses :one
select count(title) from filter($1,$2::bigint[]);

-- name: GetAssignments :many
select * from assignments where course_id = $1 and assignment_type_id <> 1;

-- name: GetAssignmentById :one
select id,title,content::text,description,assignment_type_id from assignments where id = $1 limit 1;

-- name: CreateSubmission :exec
Insert into submissions(content,assignment_id,info,user_id) values ($1,$2,$3,$4);

-- name: GetCourseLecturesByModuleId :many
select a.* from assignments a 
left join modules m on m.id = $1
where a.assignment_type_id = 1;

-- name: CreateAssignment :exec
Insert into assignments(module_id,course_id,title,description,content,assignment_type_id) values($1,$2,$3,$4,$5,$6);

-- name: GetMyTeached :many
Select * from courses 
left join course_teachers on courses.id = course_teachers.course_id
where course_teachers.user_id = $1;

-- name: CreateCourse :one
Insert into courses(title,description,image,course_provider) values ($1,$2,$3,$4) returning id;

-- name: AddCategoryCourse :exec
Insert into course_categories(course_id,category_id) values ($1,$2);

-- name: CreateModule :one
Insert into modules(title,course_id) values ($1,$2) returning id;

-- name: GetReadedLecturesByModule :many
SELECT 
    m.id AS module_id,
    m.title AS module_name,
    a.id AS assignment_id,
	a.assignment_type_id,
    COALESCE(pr.done IS NOT NULL, FALSE) AS read
  FROM 
    modules m
  LEFT JOIN 
    assignments a ON m.id = a.module_id
  LEFT JOIN 
    progress pr ON a.id = pr.assignment_id
where m.id = $1 and pr.user_id = $2;

-- name: GetProfileInfo :one
select firstname,description,profile from users where users.id = $1;

-- name: UpdateProfileInfo :exec
Update users set firstname = $1, description = $2, profile = $3 where users.id = $4;

-- name: MarkAssignmentDone :exec
insert into progress(assignment_id,user_id,done,pass) values ($1,$2,now(),true);

-- name: GetAllSubmissions :many
select s.* from submissions s 
left join assignments on s.assignment_id = assignments.id
left join courses on courses.id = assignments.course_id
where assignments.course_id = $1;

-- name: AddTeacher :exec
Insert into course_teachers(user_id,course_id) values ($1,$2);

-- name: GetModuleId :one
select * from modules where course_id = $1 and title = $2;
