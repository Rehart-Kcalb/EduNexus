-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
  "modules" (
    "id" BIGSERIAL PRIMARY KEY,
    "title" VARCHAR(100) NOT NULL,
    "course_id" BIGINT NOT NULL,
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id") ON DELETE CASCADE
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "modules";
-- +goose StatementEnd
