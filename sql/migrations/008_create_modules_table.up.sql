CREATE TABLE "modules" (
  "id" BIGSERIAL PRIMARY KEY,
  "title" VARCHAR(100) NOT NULL,
  "course_id" BIGINT NOT NULL,
  FOREIGN KEY ("course_id") REFERENCES "courses" ("id")
);