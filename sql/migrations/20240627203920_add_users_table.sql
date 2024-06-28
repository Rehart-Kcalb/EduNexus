-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
    "users" (
        "id" BIGSERIAL PRIMARY KEY,
        "login" VARCHAR(50) UNIQUE NOT NULL,
        "password" VARCHAR(60) NOT NULL,
        "surname" VARCHAR(50),
        "firstname" VARCHAR(100),
        "profile" VARCHAR(500),
        "description" VARCHAR(200),
        "user_role_id" INT NOT NULL,
        FOREIGN KEY ("user_role_id") REFERENCES "user_roles" ("id")
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd