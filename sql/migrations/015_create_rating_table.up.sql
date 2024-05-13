CREATE TABLE "rating" (
  "id" BIGSERIAL PRIMARY KEY,
  "comment_id" BIGINT NOT NULL,
  "rate" bool NOT NULL,
  FOREIGN KEY ("comment_id") REFERENCES "comments" ("id")
);