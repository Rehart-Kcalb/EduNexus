-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
  "categories" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL,
    "color" INT NOT NULL
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "categories"
-- +goose StatementEnd
