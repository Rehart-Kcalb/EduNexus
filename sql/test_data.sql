INSERT INTO
  user_roles (title)
VALUES
  ('user'),
  ('moderator');

INSERT INTO
  assignments_types (NAME)
VALUES
  ('quiz'),
  ('code'),
  ('matching'),
  ('fill_in'),
  ('sql_code'),
  ('sort'),
  ('free_answer'),
  ('number');

INSERT INTO
  categories ("name", "color")
VALUES
  ('computer science', x'fcba03'::INT),
  ('math', x'0ec4b8'::INT),
  ('biology', x'55e309'::INT),
  ('psychology', x'6709e3'::INT),
  ('devops', x'0c49f0'::INT),
  ('aws', x'f2550c'::INT),
  ('game development', x'e61588'::INT);

INSERT INTO
  users (
    "login",
    "password",
    "user_role_id",
    "firstname",
    "surname"
  )
VALUES
  (
    'useruseruser',
    '$2a$10$5M9QLhMVsHHPts39asoYz.E23IqCQXHnkDpLf3kEXHPEBaUhcvRua',
    1,
    'jonh',
    'doe'
  ),
  (
    'adminadmin',
    '$2a$10$5M9QLhMVsHHPts39asoYz.E23IqCQXHnkDpLf3kEXHPEBaUhcvRua',
    2,
    '',
    ''
  );

INSERT INTO
  courses (title, description)
VALUES
  (
    'C++ programming',
    'This course will teach you, how to write a sigma blazingly fast code'
  ),
  (
    'Calculus 1',
    'Another Calculus 1 course for idiot'
  ),
  (
    'Calculus 1 with proofs',
    'Sigma Calculus 1 course for someone with more than 2 braincells'
  ),
  (
    'Python with DataScience',
    'This course will teach the most brainrot programming language and the most high values field of usage'
  );

INSERT INTO
  course_categories (course_id, category_id)
VALUES
  (1, 1),
  (2, 2),
  (3, 2),
  (4, 1),
  (4, 2);

INSERT INTO
  enrollments (enrolled_on, course_id, user_id)
VALUES
  (NOW() - INTERVAL '7 DAY', 1, 1),
  (NOW() - INTERVAL '1 DAY', 2, 1);

INSERT INTO
  modules (title, course_id)
VALUES
  ('Variable,basic types', 1),
  ('Conditions and loops', 1),
  ('Functions', 1),
  ('Limits and continuity', 2),
  ('Derivatives: definition and basic rules', 2),
  ('Variable,basic types', 4),
  ('Conditions and loops', 4),
  ('Functions', 4);

INSERT INTO
  assignments (
    module_id,
    description,
    CONTENT,
    assignment_type_id,
    days
  )
VALUES
  (1, 'what types are exist', NULL, 7, NULL),
  (1, 'are you gay ?', NULL, 7, NULL);

INSERT INTO
  threads (module_id, title, CONTENT, user_id)
VALUES
  (1, 'Are python developer gay?', NULL, 1);

INSERT INTO
  COMMENTS (user_id, CONTENT)
VALUES
  (1, 'absolutly');
