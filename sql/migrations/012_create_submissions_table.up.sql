CREATE TABLE
  "submissions" (
    "id" BIGSERIAL PRIMARY KEY,
    "assignments_id" BIGINT NOT NULL,
    "deadline_id" BIGINT NOT NULL,
    "delay" INT NOT NULL,
    "content" JSONB NOT NULL,
    FOREIGN KEY ("assignments_id") REFERENCES "assignments" ("id"),
    FOREIGN KEY ("deadline_id") REFERENCES "deadlines" ("id")
  );