-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
  "courses" (
    "id" BIGSERIAL PRIMARY KEY,
    "title" VARCHAR(100) NOT NULL,
    "description" TEXT NOT NULL,
    "image" VARCHAR(500),
    "course_provider" BIGINT NOT NULL,
    FOREIGN KEY ("course_provider") REFERENCES "users" ("id") ON DELETE CASCADE
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "courses";
-- +goose StatementEnd
