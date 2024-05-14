CREATE TABLE
  "courses" (
    "id" BIGSERIAL PRIMARY KEY,
    "title" VARCHAR(100) NOT NULL,
    "description" TEXT NOT NULL,
    "image" varchar(500),
    "course_provider" BIGINT NOT NULL,
    FOREIGN KEY ("course_provider") REFERENCES "users" ("id")
  );