-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "rating" (
    "id" BIGSERIAL PRIMARY KEY,
    "comment_id" BIGINT NOT NULL,
    "rate" bool NOT NULL,
    FOREIGN KEY ("comment_id") REFERENCES "comments" ("id") ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "rating";
-- +goose StatementEnd