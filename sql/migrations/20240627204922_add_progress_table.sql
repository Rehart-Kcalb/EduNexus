-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
  "progress" (
    assignment_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    done DATE NOT NULL,
    pass bool NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("assignment_id") REFERENCES "assignments" ("id") ON DELETE CASCADE
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "progress";
-- +goose StatementEnd
