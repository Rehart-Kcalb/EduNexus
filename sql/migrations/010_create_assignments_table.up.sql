CREATE TABLE
  "assignments" (
    "id" BIGSERIAL PRIMARY KEY,
    "module_id" BIGINT NOT NULL,
    "course_id" BIGINT NOT NULL,
    "title" VARCHAR(100) NOT NULL,
    "description" TEXT NOT NULL,
    "content" text,
    "days" int,
    "assignment_type_id" BIGINT NOT NULL,
    FOREIGN KEY ("module_id") REFERENCES "modules" ("id"),
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id"),
    FOREIGN KEY ("assignment_type_id") REFERENCES "assignments_types" ("id")
  );
