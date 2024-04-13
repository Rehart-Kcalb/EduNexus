-- name: AllCategories :many
SELECT
  "name"
FROM
  categories;

-- name: GetCategoryCourses :many
SELECT
  courses.title
FROM
  courses
  LEFT JOIN course_categories cc ON courses.id = cc.course_id
  LEFT JOIN categories ON cc.category_id = categories.id
WHERE
  categories.name = $1;

