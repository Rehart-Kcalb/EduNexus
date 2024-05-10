INSERT INTO
  user_roles (title)
VALUES
  ('user'),
  ('moderator'),
  ('organization');

INSERT INTO
  assignments_types (NAME)
VALUES
  ('lecture'),
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
    "surname",
    "profile"
  )
VALUES
  (
    'useruseruser',
    '$2a$10$5M9QLhMVsHHPts39asoYz.E23IqCQXHnkDpLf3kEXHPEBaUhcvRua',
    1,
    'jonh',
    'doe',''
  ),
  (
    'adminadmin',
    '$2a$10$5M9QLhMVsHHPts39asoYz.E23IqCQXHnkDpLf3kEXHPEBaUhcvRua',
    2,
    '',
    '',''
  ),
  (
    'googleuniversity',
    '$2a$10$5M9QLhMVsHHPts39asoYz.E23IqCQXHnkDpLf3kEXHPEBaUhcvRua',
    3,
    'Google University',
    '',
    'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c1/Google_%22G%22_logo.svg/768px-Google_%22G%22_logo.svg.png'
  ),
(
    'sigmauni',
    '$2a$10$5M9QLhMVsHHPts39asoYz.E23IqCQXHnkDpLf3kEXHPEBaUhcvRua',
    3,
    'Сигма Университет',
    '',
    'https://static.vecteezy.com/system/resources/previews/009/993/728/non_2x/sigma-simple-triangle-geometric-mosaic-symbol-vector.jpg'
  );


INSERT INTO
  courses (title, description, course_provider,image)
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
  (1, 'what types are exist', NULL, 8, NULL),
  (1, 'are you gay ?', NULL, 8, NULL);

INSERT INTO
  threads (module_id, title, CONTENT, user_id)
VALUES
  (1, 'Are python developer gay?', NULL, 1);

INSERT INTO
  COMMENTS (user_id, CONTENT)
VALUES
  (1, 'absolutly');

insert into assignments(module_id,description,content,assignment_type_id)
VALUES
  (1,'','<h1>Лекция: Типы данных в C++</h1>
    <p>Язык программирования C++ поддерживает различные типы данных, которые используются для хранения различных видов значений. Вот основные типы данных в C++:</p>
    <ul>
        <li><strong>int:</strong> Целочисленный тип данных, который может хранить целые числа.</li>
        <li><strong>float:</strong> Тип данных с плавающей точкой для хранения чисел с плавающей запятой одинарной точности.</li>
        <li><strong>double:</strong> Тип данных с плавающей точкой для хранения чисел с плавающей запятой двойной точности.</li>
        <li><strong>char:</strong> Тип данных для хранения одного символа.</li>
        <li><strong>bool:</strong> Логический тип данных, который может хранить только значения true или false.</li>
    </ul>
    <p>Кроме того, в C++ можно создавать пользовательские типы данных с помощью структур и классов.</p>
    <p>Знание различных типов данных в C++ важно для правильного объявления переменных и выполнения операций в программе.</p>
    <p>Это основы типов данных в C++, которые вы будете использовать при изучении и разработке программ на этом языке.</p>',1),
  (2,'','<h1>Лекция: Ветвления и Циклы в C++</h1>
    <p>В программировании ветвления и циклы используются для управления потоком выполнения программы. В языке C++ существует несколько основных конструкций ветвления и циклов:</p>
    <ul>
        <li><strong>if-else:</strong> Конструкция, которая позволяет выполнить блок кода, если условие истинно, и другой блок кода, если условие ложно.</li>
        <li><strong>switch:</strong> Конструкция, которая позволяет выбрать один из множества вариантов выполнения кода в зависимости от значения выражения.</li>
        <li><strong>for:</strong> Цикл, который выполняет блок кода заданное количество раз.</li>
        <li><strong>while:</strong> Цикл, который выполняет блок кода, пока условие истинно.</li>
        <li><strong>do-while:</strong> Цикл, который выполняет блок кода, а затем проверяет условие, чтобы решить, нужно ли продолжить выполнение цикла.</li>
    </ul>
    <p>Эти конструкции позволяют программистам создавать более гибкие и мощные программы, которые могут адаптироваться к различным условиям и вводным данным.</p>
    <p>Понимание ветвлений и циклов в C++ является ключевым навыком для эффективного программирования на этом языке.</p>',1),
    (3,'','<h1>Лекция: Функции в C++</h1>
    <p>Функции в языке программирования C++ позволяют группировать блоки кода для выполнения конкретной задачи. Вот основные понятия и принципы функций в C++:</p>
    <ul>
        <li><strong>Объявление функции:</strong> Функция объявляется с указанием её типа возвращаемого значения, имени и параметров, если они есть. Например: <code>int add(int a, int b);</code></li>
        <li><strong>Определение функции:</strong> Функция определяется с использованием ключевого слова <code>return</code> для возврата значения. Например: <code>int add(int a, int b) { return a + b; }</code></li>
        <li><strong>Вызов функции:</strong> Функция вызывается по её имени с передачей аргументов, если они необходимы. Например: <code>int sum = add(5, 3);</code></li>
        <li><strong>Рекурсия:</strong> Возможность функции вызывать саму себя. Рекурсия часто используется для решения задач, которые могут быть разбиты на подзадачи того же типа.</li>
        <li><strong>Перегрузка функций:</strong> Возможность определить несколько функций с одним и тем же именем, но с разными списками параметров. Это позволяет создавать более универсальные и гибкие функции.</li>
    </ul>
    <p>Функции в C++ играют ключевую роль в создании модульных, многоразовых и понятных программ. Они позволяют разбивать сложные задачи на более простые компоненты и повторно использовать код.</p>
</body>',1),
(4,'','<h1>Непрерывность функций</h1>
    
    <h2>Определение непрерывности</h2>
    <p>Функция f(x) непрерывна в точке c, если для любого числа ε > 0 найдется число δ > 0 такое, что для всех x из интервала (c - δ, c + δ) выполняется неравенство |f(x) - f(c)| < ε.</p>
    
    <h2>Типы непрерывности</h2>
    <ul>
        <li><strong>Непрерывность на отрезке:</strong> Функция непрерывна на отрезке [a, b], если она непрерывна в каждой точке этого отрезка.</li>
        <li><strong>Непрерывность слева и справа:</strong> Функция непрерывна слева в точке c, если предел функции существует и равен f(c); непрерывна справа, если предел существует и равен f(c).</li>
        <li><strong>Непрерывность на интервале:</strong> Функция непрерывна на интервале (a, b), если она непрерывна в каждой точке этого интервала.</li>
    </ul>
    
    <h2>Примеры непрерывных функций</h2>
    <p>Примерами непрерывных функций являются полиномы, тригонометрические функции (синус, косинус), экспоненциальные функции и т. д.</p>
    
    <h2>Теоремы о непрерывности</h2>
    <ol>
        <li><strong>Теорема о непрерывности композиции:</strong> Если функция f(x) непрерывна в точке c, а функция g(x) непрерывна в точке f(c), то композиция g(f(x)) также непрерывна в точке c.</li>
        <li><strong>Теорема о непрерывности на отрезке:</strong> Если функция f(x) непрерывна на отрезке [a, b], то она ограничена на этом отрезке и принимает наибольшее и наименьшее значения.</li>
    </ol>
    
    <p>Это лишь краткое введение в тему непрерывности функций. Более подробное изучение требует рассмотрения различных типов функций и более сложных теорем.</p>
</body>',1);