

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
  ('Функции', 4),
  ('Основы Python', 4),
  ('Работа с данными в Python', 4),
  ('Библиотеки для DataScience в Python', 4),
  ('Введение в психологию', 5),
  ('Основные теории и методы психологии', 5),
  ('Применение психологии в повседневной жизни', 5),
  ('Основы ботаники', 6),
  ('Классификация растений', 6),
  ('Физиология растений', 6),
  ('Введение в Linux', 7),
  ('Установка и настройка Linux', 7),
  ('Основные команды Linux', 7),
  ('Основы AWS', 8),
  ('Создание облачных ресурсов', 8),
  ('Управление облачными ресурсами', 8),
  ('Введение в Docker', 9),
  ('Создание и управление контейнерами', 9),
  ('Docker Compose и оркестрация контейнеров', 9),
  ('Основы математического анализа', 10),
  ('Интегралы и их применение', 10),
  ('Численные методы в математическом анализе', 10),
  ('Основы работы с MySQL', 11),
  ('Запросы и манипуляция данными', 11),
  ('Оптимизация и индексы в MySQL', 11),
  ('Основы HTML и CSS', 12),
  ('Адаптивный дизайн', 12),
  ('Основы кибербезопасности', 13),
  ('Принципы защиты информации', 13),
  ('Киберугрозы и методы защиты', 13);

insert into
  assignments (
    module_id,
    course_id,
    title,
    "description",
    content,
    assignment_type_id
  )
VALUES
  (
    1,
    1,
    'Типы Данных',
    '',
    '{"content":"<h1>Лекция: Типы данных в Go</h1><p>Язык программирования C++ поддерживает различные типы данных, которые используются для хранения различных видов значений. Вот основные типы данных в Go:</p><ul><li><strong>int:</strong> Целочисленный тип данных, который может хранить целые числа.</li><li><strong>float:</strong> Тип данных с плавающей точкой для хранения чисел с плавающей запятой одинарной точности.</li><li><strong>double:</strong> Тип данных с плавающей точкой для хранения чисел с плавающей запятой двойной точности.</li><li><strong>char:</strong> Тип данных для хранения одного символа.</li><li><strong>bool:</strong> Логический тип данных, который может хранить только значения true или false.</li></ul><p>Кроме того, в Go можно создавать пользовательские типы данных с помощью структур и классов.</p><p>Знание различных типов данных в Go важно для правильного объявления переменных и выполнения операций в программе.</p><p>Это основы типов данных в Go, которые вы будете использовать при изучении и разработке программ на этом языке.</p>"}',
    1
  ),
  (
    2,
    1,
    'Ветвление и Циклы',
    '',
    '{"content":"<h1>Лекция: Ветвления и Циклы в Go</h1><p>В программировании ветвления и циклы используются для управления потоком выполнения программы. В языке C++ существует несколько основных конструкций ветвления и циклов:</p><ul><li><strong>if-else:</strong> Конструкция, которая позволяет выполнить блок кода, если условие истинно, и другой блок кода, если условие ложно.</li><li><strong>switch:</strong> Конструкция, которая позволяет выбрать один из множества вариантов выполнения кода в зависимости от значения выражения.</li><li><strong>for:</strong> Цикл, который выполняет блок кода заданное количество раз.</li><li><strong>while:</strong> Цикл, который выполняет блок кода, пока условие истинно.</li><li><strong>do-while:</strong> Цикл, который выполняет блок кода, а затем проверяет условие, чтобы решить, нужно ли продолжить выполнение цикла.</li></ul><p>Эти конструкции позволяют программистам создавать более гибкие и мощные программы, которые могут адаптироваться к различным условиям и вводным данным.</p><p>Понимание ветвлений и циклов в C++ является ключевым навыком для эффективного программирования на этом языке.</p>"}',
    1
  ),
  (
    3,
    1,
    'Функции',
    '',
    '{"content":"<h1>Лекция: Функции в Go</h1><p>Функции в языке программирования C++ позволяют группировать блоки кода для выполнения конкретной задачи. Вот основные понятия и принципы функций в C++:</p><ul><li><strong>Объявление функции:</strong> Функция объявляется с указанием её типа возвращаемого значения, имени и параметров, если они есть. Например: <code>int add(int a, int b);</code></li><li><strong>Определение функции:</strong> Функция определяется с использованием ключевого слова <code>return</code> для возврата значения. Например: <code>int add(int a, int b) { return a + b; }</code></li><li><strong>Вызов функции:</strong> Функция вызывается по её имени с передачей аргументов, если они необходимы. Например: <code>int sum = add(5, 3);</code></li><li><strong>Рекурсия:</strong> Возможность функции вызывать саму себя. Рекурсия часто используется для решения задач, которые могут быть разбиты на подзадачи того же типа.</li><li><strong>Перегрузка функций:</strong> Возможность определить несколько функций с одним и тем же именем, но с разными списками параметров. Это позволяет создавать более универсальные и гибкие функции.</li></ul><p>Функции в C++ играют ключевую роль в создании модульных, многоразовых и понятных программ. Они позволяют разбивать сложные задачи на более простые компоненты и повторно использовать код.</p>"}',
    1
  ),
  (
    4,
    2,
    'Непрерывность',
    '',
    '{"content":"<h1>Непрерывность функций</h1><h2>Определение непрерывности</h2><p>Функция f(x) непрерывна в точке c, если для любого числа ε > 0 найдется число δ > 0 такое, что для всех x из интервала (c - δ, c + δ) выполняется неравенство |f(x) - f(c)| < ε.</p><h2>Типы непрерывности</h2><ul><li><strong>Непрерывность на отрезке:</strong> Функция непрерывна на отрезке [a, b], если она непрерывна в каждой точке этого отрезка.</li><li><strong>Непрерывность слева и справа:</strong> Функция непрерывна слева в точке c, если предел функции существует и равен f(c); непрерывна справа, если предел существует и равен f(c).</li><li><strong>Непрерывность на интервале:</strong> Функция непрерывна на интервале (a, b), если она непрерывна в каждой точке этого интервала.</li></ul><h2>Примеры непрерывных функций</h2><p>Примерами непрерывных функций являются полиномы, тригонометрические функции (синус, косинус), экспоненциальные функции и т. д.</p><h2>Теоремы о непрерывности</h2><ol><li><strong>Теорема о непрерывности композиции:</strong> Если функция f(x) непрерывна в точке c, а функция g(x) непрерывна в точке f(c), то композиция g(f(x)) также непрерывна в точке c.</li><li><strong>Теорема о непрерывности на отрезке:</strong> Если функция f(x) непрерывна на отрезке [a, b], то она ограничена на этом отрезке и принимает наибольшее и наименьшее значения.</li></ol><p>Это лишь краткое введение в тему непрерывности функций. Более подробное изучение требует рассмотрения различных типов функций и более сложных теорем.</p>"}',
    1
  ),
   (
    5,
    2,
    'Производная и её применение',
    '',
    '{"content":"<h1>Лекция: Производная и её применение</h1><p>Производная функции f(x) обозначает скорость изменения функции в данной точке. Она используется для нахождения касательных к графику функции и решения задач оптимизации.</p><p>Пример: производная функции f(x) = x^2 равна 2x.</p>"}',
    2
  ),
  (
    6,
    4,
    'Типы Данных в Python',
    '',
    '{"content":"<h1>Лекция: Типы данных в Python</h1><p>Язык программирования Python поддерживает различные типы данных, включая числовые (int, float), строки (str), логические (bool) и коллекции (списки, кортежи, множества и словари).</p><p>Знание типов данных важно для правильного использования переменных и функций в Python.</p>"}',
    1
  ),
  (
    7,
    4,
    'Условные операторы и циклы в Python',
    '',
    '{"content":"<h1>Лекция: Условные операторы и циклы в Python</h1><p>Условные операторы if, elif и else используются для выполнения кода в зависимости от условий. Циклы for и while позволяют повторять выполнение кода.</p><p>Пример условного оператора:</p><pre><code>if condition:\n  # выполняется, если condition истинно\nelif another_condition:\n  # выполняется, если another_condition истинно\nelse:\n  # выполняется, если все условия ложны</code></pre>"}',
    1
  ),
  (
    8,
    4,
    'Функции в Python',
    '',
    '{"content":"<h1>Лекция: Функции в Python</h1><p>Функции в Python определяются с помощью ключевого слова def и позволяют организовать код в блоки, которые можно вызывать по имени. Функции могут принимать параметры и возвращать значения.</p><p>Пример функции:</p><pre><code>def имя_функции(параметры):\n  # тело функции\n  return значение</code></pre>"}',
    1
  ),
  (
    7,
    4,
    'Циклы в Python',
    '',
    '{"content":"<h1>Лекция: Циклы в Python</h1><p>Циклы for и while позволяют выполнять повторяющиеся задачи. Цикл for используется для итерации по последовательностям, а цикл while выполняется, пока условие истинно.</p><p>Пример:</p><pre><code>for i in range(5):\n  print(i)\n\nx = 0\nwhile x < 5:\n  print(x)\n  x += 1</code></pre>"}',
    1
  ),
  (
    12,
    5,
    'Введение в психологию',
    '',
    '{"content":"<h1>Лекция: Введение в психологию</h1><p>Психология изучает поведение и психические процессы. Основные области психологии включают когнитивную, социальную, клиническую и организационную психологию.</p>"}',
    1
  ),
  (
    13,
    5,
    'Основные теории психологии',
    '',
    '{"content":"<h1>Лекция: Основные теории психологии</h1><p>Существуют различные теории психологии, такие как психоанализ, бихевиоризм, когнитивная психология и гуманистическая психология.</p><p>Каждая теория предлагает свой взгляд на природу поведения и психических процессов.</p>"}',
    1
  ),
  (
    14,
    5,
    'Психология в повседневной жизни',
    '',
    '{"content":"<h1>Лекция: Психология в повседневной жизни</h1><p>Психология может помочь в понимании и улучшении межличностных отношений, управлении стрессом и улучшении качества жизни.</p>"}',
    1
  ),
  (
    15,
    6,
    'Основы ботаники',
    '',
    '{"content":"<h1>Лекция: Основы ботаники</h1><p>Ботаника изучает растения, их строение, функции и классификацию. Основные группы растений включают мхи, папоротники, голосеменные и цветковые растения.</p>"}',
    1
  ),
  (
    16,
    6,
    'Классификация растений',
    '',
    '{"content":"<h1>Лекция: Классификация растений</h1><p>Классификация растений основана на их морфологических и генетических характеристиках. Основные категории включают царство, отдел, класс, порядок, семейство, род и вид.</p>"}',
    1
  ),
  (
    17,
    6,
    'Физиология растений',
    '',
    '{"content":"<h1>Лекция: Физиология растений</h1><p>Физиология растений изучает функции и процессы, происходящие в растениях, такие как фотосинтез, дыхание и транспирация.</p>"}',
    1
  ),
  (
    18,
    7,
    'Основы работы в Linux',
    '',
    '{"content":"<h1>Лекция: Основы работы в Linux</h1><p>Linux — это операционная система с открытым исходным кодом. Основные команды Linux включают управление файлами и папками, установку программ и администрирование системы.</p>"}',
    1
  ),
  (
    19,
    7,
    'Установка и настройка Linux',
    '',
    '{"content":"<h1>Лекция: Установка и настройка Linux</h1><p>Установка Linux включает выбор дистрибутива, создание загрузочного носителя и настройку системы после установки.</p>"}',
    1
  ),
  (
    20,
    7,
    'Основные команды Linux',
    '',
    '{"content":"<h1>Лекция: Основные команды Linux</h1><p>Команды Linux позволяют управлять системой через командную строку. Основные команды включают ls, cd, cp, mv, rm и другие.</p>"}',
    1
  ),
  (
    21,
    8,
    'Основы AWS',
    '',
    '{"content":"<h1>Лекция: Основы AWS</h1><p>Amazon Web Services (AWS) предоставляет облачные сервисы для создания и управления ресурсами в облаке. Основные сервисы включают EC2, S3, RDS и другие.</p>"}',
    1
  ),
  (
    22,
    8,
    'Создание облачных ресурсов в AWS',
    '',
    '{"content":"<h1>Лекция: Создание облачных ресурсов в AWS</h1><p>Создание ресурсов в AWS включает настройку виртуальных машин (EC2), хранение данных (S3) и управление базами данных (RDS).</p>"}',
    1
  ),
  (
    34,
    13,
    'Введение в кибербезопасность',
    'Основные понятия и значимость кибербезопасности',
    '{"content":"<h1>Введение в кибербезопасность</h1><p>Кибербезопасность - это практика защиты систем, сетей и программ от цифровых атак. Эти атаки обычно направлены на доступ, изменение или уничтожение конфиденциальной информации, вымогательство денег у пользователей или нарушение нормальной работы бизнес-процессов.</p><p>С ростом числа киберугроз важность кибербезопасности значительно возросла. Сегодня, когда все больше информации хранится в цифровом виде, обеспечение безопасности данных стало одной из главных задач для организаций и частных лиц.</p><p>В этой лекции мы рассмотрим основные аспекты кибербезопасности, включая:</p><ul><li>Определение кибербезопасности</li><li>Основные типы кибератак</li><li>Принципы защиты информации</li><li>Значимость кибербезопасности в современном мире</li></ul><p>Понимание этих концепций поможет вам лучше осознавать риски и применять меры для защиты своих данных и систем.</p>"}',
    1
  ),
  (
    34,
    13,
    'Основные типы кибератак',
    'Описание различных видов кибератак и их особенностей',
    '{"content":"<h1>Основные типы кибератак</h1><p>Существует множество видов кибератак, каждая из которых имеет свои особенности и методы воздействия на системы и данные. Рассмотрим наиболее распространенные из них:</p><ul><li><strong>Фишинг:</strong> Мошеннические попытки получить конфиденциальную информацию, выдавая себя за доверенные лица.</li><li><strong>Вредоносное ПО:</strong> Программы, такие как вирусы, трояны и шпионское ПО, которые наносят вред системам или получают несанкционированный доступ к информации.</li><li><strong>DDoS-атаки:</strong> Атаки, направленные на перегрузку системы или сети, чтобы сделать их недоступными для пользователей.</li><li><strong>Сниффинг:</strong> Прослушивание сетевого трафика с целью перехвата данных.</li><li><strong>SQL-инъекции:</strong> Внедрение вредоносного кода в SQL-запросы, что позволяет злоумышленнику получить доступ к базе данных.</li></ul><p>Знание этих типов атак поможет вам лучше понимать, как они работают и какие меры защиты необходимо принимать для их предотвращения.</p>"}',
    1
  ),
  (
    30,
    12,
    'Основы HTML',
    'Изучение основ HTML и его структуры',
    '{"content":"<h1>Основы HTML</h1><p>HTML (Hyper Text Markup Language) - это стандартный язык разметки для создания веб-страниц. Он используется для описания структуры веб-страницы с помощью тегов.</p><p>Основные элементы HTML включают:</p><ul><li><strong>&lt;html&gt;:</strong> Корневой элемент HTML-документа.</li><li><strong>&lt;head&gt;:</strong> Содержит метаинформацию о документе, такую как его заголовок и ссылки на стили.</li><li><strong>&lt;title&gt;:</strong> Определяет заголовок документа, который отображается в заголовке браузера.</li><li><strong>&lt;body&gt;:</strong> Содержит основное содержимое HTML-документа, которое отображается в браузере.</li><li><strong>&lt;p&gt;:</strong> Определяет абзац текста.</li><li><strong>&lt;a&gt;:</strong> Создает гиперссылку на другую веб-страницу или ресурс.</li><li><strong>&lt;img&gt;:</strong> Вставляет изображение в веб-страницу.</li></ul><p>Знание этих и других тегов HTML позволяет создавать базовую структуру веб-страниц и добавлять различные элементы контента.</p>"}',
    1
  ),
  (
    30,
    12,
    'Основы CSS',
    'Изучение основ CSS и его применения для стилизации веб-страниц',
    '{"content":"<h1>Основы CSS</h1><p>CSS (Cascading Style Sheets) - это язык описания стилей, используемый для стилизации HTML-документов. С помощью CSS можно задавать внешний вид и оформление веб-страниц.</p><p>Основные элементы CSS включают:</p><ul><li><strong>Селекторы:</strong> Определяют, к каким HTML-элементам будут применяться стили. Примеры селекторов: элементные (h1, p), классовые (.class), ID (#id).</li><li><strong>Свойства:</strong> Определяют, какие именно стили будут применяться. Примеры свойств: color, font-size, background-color.</li><li><strong>Значения:</strong> Определяют конкретные значения свойств. Примеры значений: red, 16px, #ffffff.</li></ul><p>Пример CSS-кода:</p><pre><code>body {<br>  background-color: #f0f0f0;<br>}<br>h1 {<br>  color: #333;<br>}<br>.main {<br>  font-size: 18px;<br>}</code></pre><p>С помощью CSS можно изменять цвет текста, размеры шрифтов, отступы, границы и многое другое, что позволяет делать веб-страницы более привлекательными и удобными для пользователя.</p>"}',
    1
  );

