-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
  "comments" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "content" TEXT NOT NULL,
    "thread_id" BIGINT NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("thread_id") REFERENCES "threads" ("id") ON DELETE CASCADE
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "comments";
-- +goose StatementEnd
