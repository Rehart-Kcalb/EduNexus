-- name: FilterCourses :many
--select filter.title, image,organization_name, organization_logo from filter($1,$2::bigint[]) limit $3 offset $4;
select * from filter($1,$2::bigint[]) limit $3 offset $4;


-- name: CountCourses :one
select count(title) from filter($1,$2::bigint[]);

-- name: GetPopularCourses :many
SELECT
    c.title,
    u.firstname AS organization_name,
    u.profile AS organization_logo,
    c.image
FROM
    courses c
LEFT JOIN
    users u ON u.id = c.course_provider
INNER JOIN
    (SELECT 
        course_id, 
        COUNT(user_id) AS enrolled 
     FROM 
        enrollments 
     GROUP BY 
        course_id) enrolled_counts ON enrolled_counts.course_id = c.id
ORDER BY
    enrolled_counts.enrolled DESC;

