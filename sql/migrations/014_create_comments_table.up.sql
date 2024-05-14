CREATE TABLE
  "comments" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "content" TEXT NOT NULL,
    "thread_id" BIGINT NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    FOREIGN KEY ("thread_id") REFERENCES "threads" ("id")
  );