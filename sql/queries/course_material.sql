-- name: GetCourseLectures :many
SELECT DISTINCT
    a.title,
    a.id AS assignment_id,
    COALESCE(pr.done IS NOT NULL, FALSE) AS read,
    m.title as module_name
FROM 
    courses c 
LEFT JOIN 
    modules m ON m.course_id = c.id
LEFT JOIN 
    assignments a ON a.module_id = m.id
LEFT JOIN 
    progress pr ON a.id = pr.assignment_id AND pr.user_id = $2
WHERE  
    c.id = $1
    AND a.id IS NOT NULL 
    AND a.assignment_type_id = 1
order by a.id asc;

-- name: GetLectureContent :one
  select 
    title, content::text 
  from assignments
  where
    assignments.id = $1;

-- name: NewLecture :exec
insert into assignments (module_id,course_id,title,description,content,assignment_type_id)
values ($1,$2,$3,$4,$5,1);

-- name: GetAssignments :many
SELECT DISTINCT
    a.*,
    m.title as module_name,
    COALESCE(pr.done IS NOT NULL, FALSE) AS read
FROM 
    assignments a
LEFT JOIN 
    progress pr ON a.id = pr.assignment_id AND pr.user_id = $2
left join
    modules m on a.module_id = m.id
WHERE  
    a.course_id = $1 
    AND a.assignment_type_id <> 1;


-- name: GetAssignmentById :one
  select 
    id,title,content::text,description,assignment_type_id 
  from assignments 
  where 
    id = $1 limit 1;

-- name: GetCourseLecturesByModuleId :many
  select 
    a.* 
  from assignments a 
  left join 
      modules m on m.id = $1
  where 
      a.assignment_type_id = 1;

