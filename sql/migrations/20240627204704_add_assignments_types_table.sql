-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
   "assignments_types" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "assignments_types"
-- +goose StatementEnd
