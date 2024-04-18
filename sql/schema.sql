-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2024-04-18T16:23:37.280Z

CREATE TABLE "user_roles" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(30) NOT NULL
);

CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "login" VARCHAR(50) UNIQUE NOT NULL,
  "password" VARCHAR(60) NOT NULL,
  "surname" VARCHAR(50),
  "firstname" VARCHAR(50),
  "user_role_id" INT NOT NULL
);

CREATE TABLE "categories" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(50) NOT NULL,
  "color" INT NOT NULL
);

CREATE TABLE "courses" (
  "id" BIGSERIAL PRIMARY KEY,
  "title" VARCHAR(100) NOT NULL,
  "description" TEXT NOT NULL
);

CREATE TABLE "course_categories" (
  "course_id" BIGINT,
  "category_id" BIGINT
);

CREATE TABLE "enrollments" (
  "id" BIGSERIAL PRIMARY KEY,
  "enrolled_on" DATE NOT NULL,
  "course_id" BIGINT NOT NULL,
  "user_id" BIGINT NOT NULL
);

CREATE TABLE "modules" (
  "id" BIGSERIAL PRIMARY KEY,
  "title" VARCHAR(100) NOT NULL,
  "course_id" BIGINT NOT NULL
);

CREATE TABLE "assignments_types" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(100) NOT NULL
);

CREATE TABLE "assignments" (
  "id" BIGSERIAL PRIMARY KEY,
  "module_id" BIGINT NOT NULL,
  "description" TEXT NOT NULL,
  "content" JSONB,
  "days" int,
  "assignment_type_id" BIGINT NOT NULL
);

CREATE TABLE "deadlines" (
  "id" BIGSERIAL PRIMARY KEY,
  "assignments_id" BIGINT NOT NULL,
  "deadline" TIMESTAMP,
  "user_id" BIGINT NOT NULL
);

CREATE TABLE "submissions" (
  "id" BIGSERIAL PRIMARY KEY,
  "assignments_id" BIGINT NOT NULL,
  "deadline_id" BIGINT NOT NULL,
  "delay" INT NOT NULL,
  "content" JSONB NOT NULL
);

CREATE TABLE "threads" (
  "id" BIGSERIAL PRIMARY KEY,
  "module_id" BIGINT NOT NULL,
  "title" VARCHAR(200) NOT NULL,
  "content" TEXT NOT NULL,
  "user_id" BIGINT NOT NULL
);

CREATE TABLE "comments" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGINT NOT NULL,
  "content" TEXT NOT NULL
);

CREATE TABLE "rating" (
  "id" BIGSERIAL PRIMARY KEY,
  "comment_id" BIGINT NOT NULL,
  "rate" bool NOT NULL
);

ALTER TABLE "users" ADD FOREIGN KEY ("user_role_id") REFERENCES "user_roles" ("id");

ALTER TABLE "course_categories" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "course_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "enrollments" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "enrollments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "modules" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "assignments" ADD FOREIGN KEY ("module_id") REFERENCES "modules" ("id");

ALTER TABLE "assignments" ADD FOREIGN KEY ("assignment_type_id") REFERENCES "assignments_types" ("id");

ALTER TABLE "deadlines" ADD FOREIGN KEY ("assignments_id") REFERENCES "assignments" ("id");

ALTER TABLE "deadlines" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "submissions" ADD FOREIGN KEY ("assignments_id") REFERENCES "assignments" ("id");

ALTER TABLE "submissions" ADD FOREIGN KEY ("deadline_id") REFERENCES "deadlines" ("id");

ALTER TABLE "threads" ADD FOREIGN KEY ("module_id") REFERENCES "modules" ("id");

ALTER TABLE "threads" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "rating" ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("id");
