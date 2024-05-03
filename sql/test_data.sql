INSERT INTO
  user_roles (title)
VALUES
  ('user'),
  ('moderator'),
  ('organization');

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
  ('Компьютерные Наука', x'fcba03'::INT),
  ('Математика', x'0ec4b8'::INT),
  ('Биология', x'55e309'::INT),
  ('Психология', x'6709e3'::INT),
  ('Девопс', x'0c49f0'::INT),
  ('AWS', x'f2550c'::INT),
  ('Разработка Игр', x'e61588'::INT);

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
  ),
  (
    'googleuniversity',
    '$2a$10$5M9QLhMVsHHPts39asoYz.E23IqCQXHnkDpLf3kEXHPEBaUhcvRua',
    3,
    'Google University',
    ''
  ),
(
    'sigmauni',
    '$2a$10$5M9QLhMVsHHPts39asoYz.E23IqCQXHnkDpLf3kEXHPEBaUhcvRua',
    3,
    'Сигма Университет',
    ''
  );


INSERT INTO
  courses (title, description, course_provider)
VALUES
  (
    'C++ programming',
    'This course will teach you, how to write a sigma blazingly fast code',
    3
  ),
  (
    'Дифферинциальное Исчисление для ПТУ',
    'Курс дифферинциального исчисления для политехнический училищ',
    4
  ),
  (
    'Дифферинциальное Исчисление для вузов',
    'Курс дифферинциального исчисления для студентов вузов по техническим специальностям',
    4
  ),
  (
    'Python with DataScience',
    'Этот курс научит вас самому востребованнуму языку программирования и самой высокооплачиваемой сфере использования',
    3
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
  course_teachers (course_id, user_id)
VALUES
  (1, 3), -- User 1 teach c++
  (2, 4) -- User 1 teach calculus
;

INSERT INTO
  enrollments (enrolled_on, course_id, user_id)
VALUES
  (NOW() - INTERVAL '7 DAY', 1, 1),
  (NOW() - INTERVAL '1 DAY', 2, 1);

INSERT INTO
  modules (title, course_id)
VALUES
  ('Переменные и базовые типы', 1),
  ('Ветвление и Циклы', 1),
  ('Функции', 1),
  ('Пределы и Непрерывность', 2),
  ('Производная: Определения и правила', 2),
  ('Переменные и базовые типы', 4),
  ('Ветвление и Циклы', 4),
  ('Функции', 4);

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
