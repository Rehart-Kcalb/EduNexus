-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
  "assignments" (
    "id" BIGSERIAL PRIMARY KEY,
    "module_id" BIGINT NOT NULL,
    "course_id" BIGINT NOT NULL,
    "title" VARCHAR(100) NOT NULL,
    "description" TEXT NOT NULL,
    "content" jsonb,
    "days" INT,
    "assignment_type_id" BIGINT NOT NULL,
    "created_at" DATE DEFAULT NOW(),
    FOREIGN KEY ("module_id") REFERENCES "modules" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("assignment_type_id") REFERENCES "assignments_types" ("id")
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "assignments";
-- +goose StatementEnd
