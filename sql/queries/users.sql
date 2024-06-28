-- name: GetPasswordByLogin :one
SELECT
  PASSWORD
FROM
  users
WHERE
  login = $1;

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
  users ("login", "password", "user_role_id", firstname)
VALUES
  ($1, $2, 1,$1);

