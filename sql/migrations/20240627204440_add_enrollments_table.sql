-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
  "enrollments" (
    "id" BIGSERIAL PRIMARY KEY,
    "enrolled_on" DATE NOT NULL,
    "course_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE  IF EXISTS "enrollments";
-- +goose StatementEnd
