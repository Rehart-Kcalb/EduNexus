CREATE TABLE
  "threads" (
    "id" BIGSERIAL PRIMARY KEY,
    "module_id" BIGINT NOT NULL,
    "title" VARCHAR(200) NOT NULL,
    "content" TEXT,
    "user_id" BIGINT NOT NULL,
    FOREIGN KEY ("module_id") REFERENCES "modules" ("id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  );