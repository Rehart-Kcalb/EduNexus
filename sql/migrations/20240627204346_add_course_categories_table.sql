-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
  "course_categories" (
    "course_id" BIGINT NOT NULL,
    "category_id" BIGINT NOT NULL,
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("category_id") REFERENCES "categories" ("id")
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "course_categories";
-- +goose StatementEnd
