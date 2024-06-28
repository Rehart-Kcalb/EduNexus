-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
  "submissions" (
    "id" BIGSERIAL PRIMARY KEY,
    "assignment_id" BIGINT NOT NULL,
    "delay" INT NULL,
    "content" JSONB NOT NULL,
    "info" TEXT NULL,
    "user_id" BIGINT NOT NULL,
    "submitted_at" DATE DEFAULT NOW(),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("assignment_id") REFERENCES "assignments" ("id") ON DELETE CASCADE
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "submissions";
-- +goose StatementEnd
