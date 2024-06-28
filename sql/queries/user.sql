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



-- name: EnrollIntoCourse :exec
INSERT INTO
  enrollments (course_id, user_id, enrolled_on)
VALUES
  (
    $1,
    $2,
    NOW()
  );

-- name: CheckEnrollment :one
select enrolled_on from enrollments where course_id = $1 and user_id = $2;

-- name: DropCourse :exec
Delete from enrollments where user_id = $1 and course_id = $2;

-- name: GetProfileInfo :one
select firstname,description,profile from users where users.id = $1;

-- name: UpdateProfileInfo :exec
Update users set firstname = $1, description = $2, profile = $3 where users.id = $4;

-- name: CreateSubmission :exec
Insert into submissions(content,assignment_id,info,user_id) values ($1,$2,$3,$4);

-- name: GetReadedLecturesByModule :many
SELECT 
    distinct(m.id) AS module_id,
    m.title AS module_name,
    a.title,
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

-- name: GetLastGradesByCourse :many
SELECT distinct(assignments.title), info, submissions.submitted_at FROM public.submissions 
	left join assignments on assignments.id = submissions.assignment_id
	left join courses on courses.id = assignments.course_id
where user_id = $1 and courses.title = $2
order by submissions.submitted_at desc;
