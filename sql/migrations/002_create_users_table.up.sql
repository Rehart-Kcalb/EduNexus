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