CREATE TABLE
  "user_roles" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(30) NOT NULL
  );

  CREATE TABLE
  "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "login" VARCHAR(50) UNIQUE NOT NULL,
    "password" VARCHAR(60) NOT NULL,
    "surname" VARCHAR(50),
    "firstname" VARCHAR(100),
    "profile" varchar(500),
    "user_role_id" INT NOT NULL,
    FOREIGN KEY ("user_role_id") REFERENCES "user_roles" ("id")
  );

  CREATE TABLE
  "categories" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL,
    "color" INT NOT NULL
  );

  CREATE TABLE
  "courses" (
    "id" BIGSERIAL PRIMARY KEY,
    "title" VARCHAR(100) NOT NULL,
    "description" TEXT NOT NULL,
    "image" varchar(500),
    "course_provider" BIGINT NOT NULL,
    FOREIGN KEY ("course_provider") REFERENCES "users" ("id")
  );

  CREATE TABLE
  "course_teachers" (
    "user_id" BIGINT NOT NULL,
    "course_id" BIGINT NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id")
  );

  CREATE TABLE
  "course_categories" (
    "course_id" BIGINT NOT NULL,
    "category_id" BIGINT NOT NULL,
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id"),
    FOREIGN KEY ("category_id") REFERENCES "categories" ("id")
  );

  CREATE TABLE
  "enrollments" (
    "id" BIGSERIAL PRIMARY KEY,
    "enrolled_on" DATE NOT NULL,
    "course_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  );

  CREATE TABLE
  "modules" (
    "id" BIGSERIAL PRIMARY KEY,
    "title" VARCHAR(100) NOT NULL,
    "course_id" BIGINT NOT NULL,
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id")
  );

  CREATE TABLE
  "assignments_types" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL
  );

  CREATE TABLE
  "assignments" (
    "id" BIGSERIAL PRIMARY KEY,
    "module_id" BIGINT NOT NULL,
    "course_id" BIGINT NOT NULL,
    "title" VARCHAR(100) NOT NULL,
    "description" TEXT NOT NULL,
    "content" jsonb,
    "days" int,
    "assignment_type_id" BIGINT NOT NULL,
    FOREIGN KEY ("module_id") REFERENCES "modules" ("id"),
    FOREIGN KEY ("course_id") REFERENCES "courses" ("id"),
    FOREIGN KEY ("assignment_type_id") REFERENCES "assignments_types" ("id")
  );

  CREATE TABLE
  "deadlines" (
    "id" BIGSERIAL PRIMARY KEY,
    "assignments_id" BIGINT NOT NULL,
    "deadline" TIMESTAMP,
    "user_id" BIGINT NOT NULL,
    FOREIGN KEY ("assignments_id") REFERENCES "assignments" ("id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  );

  CREATE TABLE
  "submissions" (
    "id" BIGSERIAL PRIMARY KEY,
    "assignment_id" BIGINT NOT NULL,
    "delay" INT  NULL,
    "content" JSONB NOT NULL,
    "info" text  NULL,
    "user_id" BIGINT NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    FOREIGN KEY ("assignments_id") REFERENCES "assignments" ("id")
  );


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

  CREATE TABLE
  "comments" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "content" TEXT NOT NULL,
    "thread_id" BIGINT NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    FOREIGN KEY ("thread_id") REFERENCES "threads" ("id")
  );

  CREATE TABLE
  "rating" (
    "id" BIGSERIAL PRIMARY KEY,
    "comment_id" BIGINT NOT NULL,
    "rate" bool NOT NULL,
    FOREIGN KEY ("comment_id") REFERENCES "comments" ("id")
  );