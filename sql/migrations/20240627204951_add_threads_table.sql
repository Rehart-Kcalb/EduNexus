-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
  "threads" (
    "id" BIGSERIAL PRIMARY KEY,
    "module_id" BIGINT NOT NULL,
    "title" VARCHAR(200) NOT NULL,
    "content" TEXT,
    "user_id" BIGINT NOT NULL,
    FOREIGN KEY ("module_id") REFERENCES "modules" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "threads";
-- +goose StatementEnd
