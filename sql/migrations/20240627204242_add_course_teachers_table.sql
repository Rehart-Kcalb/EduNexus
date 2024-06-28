-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
  "course_teachers" (
    "user_id" BIGINT NOT NULL,
    "course_id" BIGINT NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id") ON DELETE CASCADE
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "course_teachers"
-- +goose StatementEnd
