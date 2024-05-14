CREATE TABLE
  "deadlines" (
    "id" BIGSERIAL PRIMARY KEY,
    "assignments_id" BIGINT NOT NULL,
    "deadline" TIMESTAMP,
    "user_id" BIGINT NOT NULL,
    FOREIGN KEY ("assignments_id") REFERENCES "assignments" ("id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  );