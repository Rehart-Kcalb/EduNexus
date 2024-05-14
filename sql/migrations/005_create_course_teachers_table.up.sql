CREATE TABLE
  "course_teachers" (
    "user_id" BIGINT NOT NULL,
    "course_id" BIGINT NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id")
  );