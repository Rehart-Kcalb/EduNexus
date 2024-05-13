CREATE TABLE "course_categories" (
  "course_id" BIGINT NOT NULL,
  "category_id" BIGINT NOT NULL,
  FOREIGN KEY ("course_id") REFERENCES "courses" ("id"),
  FOREIGN KEY ("category_id") REFERENCES "categories" ("id")
);