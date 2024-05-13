CREATE TABLE "comments" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGINT NOT NULL,
  "content" TEXT NOT NULL,
  FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);