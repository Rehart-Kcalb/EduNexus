

INSERT INTO
  courses (title, description, course_provider, image)
VALUES
  (
    'C++ programming',
    'This course will teach you, how to write a sigma blazingly fast code',
    3,
    'https://upload.wikimedia.org/wikipedia/commons/thumb/1/18/ISO_C%2B%2B_Logo.svg/1200px-ISO_C%2B%2B_Logo.svg.png'
  ),
  (
    'Дифферинциальное Исчисление для ПТУ',
    'Курс дифферинциального исчисления для политехнический училищ',
    4,
    'https://cdn.vectorstock.com/i/preview-1x/96/04/basic-math-outline-circle-shaped-banner-vector-47399604.jpg'
  ),
  (
    'Дифферинциальное Исчисление для вузов',
    'Курс дифферинциального исчисления для студентов вузов по техническим специальностям',
    4,
    'https://cdn.dribbble.com/users/1542367/screenshots/3873991/media/960c6f236b02420ed289a8456615ceb5.jpg?resize=400x300&vertical=center'
  ),
  (
    'Python with DataScience',
    'Этот курс научит вас самому востребованнуму языку программирования и самой высокооплачиваемой сфере использования',
    3,
    'https://media.geeksforgeeks.org/wp-content/cdn-uploads/20230318230239/Python-Data-Science-Tutorial.jpg'
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