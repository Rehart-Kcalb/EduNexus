CREATE TABLE "assignments" (
  "id" BIGSERIAL PRIMARY KEY,
  "module_id" BIGINT NOT NULL,
  "description" TEXT NOT NULL,
  "content" text,
  "days" int,
  "assignment_type_id" BIGINT NOT NULL,
  FOREIGN KEY ("module_id") REFERENCES "modules" ("id"),
  FOREIGN KEY ("assignment_type_id") REFERENCES "assignments_types" ("id")
);