-- name: AllCategories :many
SELECT
  "name",
  "color"
FROM
  categories;

-- name: GetCategoryId :one
select id from categories where name=$1 limit 1;

