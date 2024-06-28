-- name: CreateAssignment :exec
Insert into assignments(module_id,course_id,title,description,content,assignment_type_id) values($1,$2,$3,$4,$5,$6);

-- name: GetMyTeached :many
Select distinct(courses.id), courses.title, courses.image,courses.description from courses 
left join course_teachers on courses.id = course_teachers.course_id
where course_teachers.user_id = $1;

-- name: CreateCourse :one
Insert into courses(title,description,image,course_provider) values ($1,$2,$3,$4) returning id;

-- name: AddCategoryCourse :exec
Insert into course_categories(course_id,category_id) values ($1,$2);

-- name: CreateModule :one
  Insert into 
    modules(title,course_id) values 
  ($1,$2) 
  returning id;

-- name: AddTeacher :exec
Insert into course_teachers(user_id,course_id) values ($1,$2);

-- name: MarkAssignmentDone :exec
  insert into 
    progress(assignment_id,user_id,done,pass) 
  values ($1,$2,now(),true);

-- name: GetAllSubmissions :many
select s.* from submissions s 
left join assignments on s.assignment_id = assignments.id
left join courses on courses.id = assignments.course_id
where assignments.course_id = $1;

