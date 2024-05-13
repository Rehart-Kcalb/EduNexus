CREATE TABLE "enrollments" (
  "id" BIGSERIAL PRIMARY KEY,
  "enrolled_on" DATE NOT NULL,
  "course_id" BIGINT NOT NULL,
  "user_id" BIGINT NOT NULL,
  FOREIGN KEY ("course_id") REFERENCES "courses" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);