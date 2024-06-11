INSERT INTO public.assignments(
    module_id, course_id, title, description, content, days, assignment_type_id)
VALUES (1, 1, 'Викторина по типам данных', 'Данная викторина проверит ваши знания по типам данных в го', 
'{"questions":["1. Какое значение по умолчанию имеет тип int в Go?","2. Какой из следующих строковых литералов является допустимым в Go?","3. Каков размер float32 в Go?","4. Какой из следующих типов данных может использоваться для представления значений true или false?","5. Как объявить константу в Go?"],"variant":[["0","nil","false","undefined"],["''Hello''","\"Hello\"","`Hello`","Все вышеперечисленные"],["16 бит","32 бит","64 бит","128 бит"],["int","string","bool","byte"],["var constName = value","const constName value","const constName = value","constant constName = value"]],"answers":["0","Все вышеперечисленные","32 бит","bool","const constName = value"]}', 
null, 2);

-- INSERT INTO public.assignments(
--     module_id, course_id, title, description, content, days, assignment_type_id)
-- VALUES (3, 1, 'Сопоставление по функциям', 'Данное упражнение проверит ваши знания о функциях в Go', 
-- '{"matching_pairs":[{"left":"Объявление функции","right":"func functionName()"},{"left":"Возврат нескольких значений","right":"return value1, value2"}, {"left":"Анонимная функция","right":"func() {}"},{"left":"Передача функции в качестве аргумента","right":"funcName(func(a int))"}, {"left":"Функция с переменным числом аргументов","right":"func functionName(args ...int)"}]}', 
-- null, 2);

INSERT INTO public.assignments(
    module_id, course_id, title, description, content, days, assignment_type_id)
VALUES
  (
    1,
    1,
    'Summator',
    '',
    '{
            "language": "go",
            "code_test": "package main\n\nimport \"testing\"\n\n// sumator adds two integers and returns the result\n// TestSumator tests the sumator function\nfunc TestSumator(t *testing.T) {\n    tests := []struct {\n        i1, i2, expected int\n    }{\n        {1, 2, 3},\n        {-1, -1, -2},\n        {0, 0, 0},\n        {123, 456, 579},\n    }\n\n    for _, test := range tests {\n        result := sumator(test.i1, test.i2)\n        if result != test.expected {\n            t.Errorf(\"sumator(%d, %d) = %d; want %d\", test.i1, test.i2, result, test.expected)\n        }\n    }\n}",
            "quiz_question": ""
    }',
    null,
    3
);

INSERT INTO public.assignments(
    module_id, course_id, title, description, content, days, assignment_type_id)
VALUES
  (
    1,
    1,
    'Обратная строка',
    'Напишите функцию, которая принимает строку и возвращает её в обратном порядке.',
    '{
            "language": "go",
            "code_test": "package main\n\nimport \"testing\"\n\n// reverseString reverses the given string\n// TestReverseString tests the reverseString function\nfunc TestReverseString(t *testing.T) {\n    tests := []struct {\n        input, expected string\n    }{\n        {\"hello\", \"olleh\"},\n        {\"world\", \"dlrow\"},\n        {\"Go\", \"oG\"},\n        {\"\", \"\"},\n    }\n\n    for _, test := range tests {\n        result := reverseString(test.input)\n        if result != test.expected {\n            t.Errorf(\"reverseString(%s) = %s; want %s\", test.input, result, test.expected)\n        }\n    }\n}",
            "quiz_question": ""
    }',
    null,
    3
);

INSERT INTO public.assignments(
    module_id, course_id, title, description, content, days, assignment_type_id)
VALUES
  (
    1,
    1,
    'Проверка палиндрома',
    'Напишите функцию, которая проверяет, является ли строка палиндромом (читается одинаково вперед и назад).',
    '{
            "language": "go",
            "code_test": "package main\n\nimport \"testing\"\n\n// isPalindrome checks if a given string is a palindrome\n// TestIsPalindrome tests the isPalindrome function\nfunc TestIsPalindrome(t *testing.T) {\n    tests := []struct {\n        input string\n        expected bool\n    }{\n        {\"madam\", true},\n        {\"racecar\", true},\n        {\"hello\", false},\n        {\"\", true},\n        {\"a\", true},\n    }\n\n    for _, test := range tests {\n        result := isPalindrome(test.input)\n        if result != test.expected {\n            t.Errorf(\"isPalindrome(%s) = %v; want %v\", test.input, result, test.expected)\n        }\n    }\n}",
            "quiz_question": ""
    }',
    null,
    3
);

INSERT INTO public.assignments(
    module_id, course_id, title, description, content, days, assignment_type_id)
VALUES
  (
    1,
    1,
    'Найти наибольшее число в массиве',
    'Напишите функцию, которая принимает массив целых чисел и возвращает наибольшее число в этом массиве.',
    '{
            "language": "go",
            "code_test": "package main\n\nimport \"testing\"\n\n// maxInArray finds the largest number in an array\n// TestMaxInArray tests the maxInArray function\nfunc TestMaxInArray(t *testing.T) {\n    tests := []struct {\n        input []int\n        expected int\n    }{\n        {[]int{1, 2, 3}, 3},\n        {[]int{-1, -3, -2}, -1},\n        {[]int{0, 10, -10}, 10},\n        {[]int{123, 456, 789}, 789},\n    }\n\n    for _, test := range tests {\n        result := maxInArray(test.input)\n        if result != test.expected {\n            t.Errorf(\"maxInArray(%v) = %d; want %d\", test.input, result, test.expected)\n        }\n    }\n}",
            "quiz_question": ""
    }',
    null,
    3
);


INSERT INTO public.assignments(
    module_id, course_id, title, description, content, days, assignment_type_id)
VALUES
  (
    4,
    2,
    'Викторина по непрерывным функциям',
    'Эта викторина проверит ваши знания о непрерывных функциях и их свойствах.',
    '{"questions":["1. Что означает, что функция непрерывна в точке?","2. Как записывается условие непрерывности функции f(x) в точке x0?","3. Какой из следующих примеров является непрерывной функцией?","4. Если функция f(x) непрерывна на интервале (a, b), то:","5. Является ли функция f(x) = |x| непрерывной на всём множестве действительных чисел?"],"variant":[["Функция принимает только целые значения в этой точке","Функция не имеет разрывов в этой точке","Функция имеет разрыв в этой точке","Функция не определена в этой точке"],["f(x) = x0","lim(x→x0) f(x) = f(x0)","f(x0) = f(x0) + 1","lim(x→x0) f(x) = 0"],["f(x) = 1/x при x ≠ 0","f(x) = 2x + 1","f(x) = |x| при x ≠ 0","f(x) = 1/(x-1)"],["Она принимает только положительные значения","Она имеет разрывы только на концах интервала","Она не имеет разрывов на всём интервале","Она определена только на интервале (a, b)"],["Да","Нет","Только в положительной области","Только в отрицательной области"]],"answers":["Функция не имеет разрывов в этой точке","lim(x→x0) f(x) = f(x0)","f(x) = 2x + 1","Она не имеет разрывов на всём интервале","Да"]}',
    null,
    2
  );

  INSERT INTO public.assignments(
    module_id, course_id, title, description, content, days, assignment_type_id)
VALUES
  (
    18,
    7,
    'Викторина по основам работы в Linux',
    'Эта викторина проверит ваши знания по основным командам и концепциям работы в Linux.',
    '{"questions":["1. Какая команда используется для отображения содержимого каталога?","2. Какую команду необходимо использовать для копирования файлов?","3. Какая команда показывает текущий рабочий каталог?","4. Какую команду следует использовать для перемещения файлов?","5. Какую команду следует использовать для удаления файлов?","6. Какая команда используется для отображения содержимого текстового файла?","7. Какая команда позволяет изменять права доступа к файлам?","8. Какую команду используют для получения справочной информации о других командах?","9. Какая команда показывает информацию о системе и версии ядра?","10. Какую команду следует использовать для создания нового каталога?"],"variant":[["ls","cd","mv","cp"],["mv","cp","rm","ls"],["pwd","whoami","cd","ls"],["mv","cp","rm","pwd"],["rm","mv","cp","ls"],["cat","ls","rm","cp"],["chmod","chown","ls","rm"],["man","ls","pwd","cd"],["uname -a","ls -l","top","df -h"],["mkdir","rmdir","ls","cp"]],"answers":["ls","cp","pwd","mv","rm","cat","chmod","man","uname -a","mkdir"]}',
    null,
    2
  );

INSERT INTO public.assignments(
    module_id, course_id, title, description, content, days, assignment_type_id)
VALUES
  (
    21,
    8,
    'Викторина по основам AWS',
    'Эта викторина проверит ваши знания о базовых концепциях и сервисах AWS.',
    '{"questions":["1. Что такое AWS?","2. Какой сервис AWS используется для хранения объектов?","3. Как называется сервис AWS для управления виртуальными машинами?","4. Какой сервис AWS используется для управления базами данных?","5. Какой сервис AWS используется для управления доменными именами?","6. Какой сервис AWS используется для развертывания контейнеров?","7. Как называется сервис AWS для управления сетью?","8. Какой сервис AWS предоставляет возможности машинного обучения?","9. Как называется сервис AWS для управления IoT устройствами?","10. Какой сервис AWS используется для мониторинга и логирования?","11. Какой сервис AWS предоставляет возможности резервного копирования и восстановления данных?"],"variant":[["Облачная платформа","Операционная система","Язык программирования","База данных"],["EC2","S3","RDS","VPC"],["EC2","S3","Lambda","RDS"],["S3","EC2","RDS","IAM"],["Route 53","S3","EC2","RDS"],["ECS","EC2","Lambda","S3"],["VPC","S3","EC2","Route 53"],["SageMaker","EC2","S3","RDS"],["IoT Core","S3","EC2","RDS"],["CloudWatch","S3","EC2","Route 53"],["Backup","EC2","RDS","S3"]],"answers":["Облачная платформа","S3","EC2","RDS","Route 53","ECS","VPC","SageMaker","IoT Core","CloudWatch","Backup"]}',
    null,
    2
  );

  INSERT INTO public.assignments(
    module_id, course_id, title, description, content, days, assignment_type_id)
VALUES
  (
    34,
    13,
    'Викторина по основам кибербезопасности',
    'Эта викторина проверит ваши знания по основам кибербезопасности.',
    '{"questions":["1. Что такое кибербезопасность?","2. Как называется программа, которая скрытно выполняет вредоносные действия?","3. Какая мера безопасности предотвращает несанкционированный доступ к данным?","4. Как называется метод атаки, при котором злоумышленник выдаёт себя за доверенное лицо?","5. Что такое шифрование?","6. Какая технология используется для защиты сети от несанкционированного доступа?","7. Что такое фишинг?","8. Какой из следующих методов используется для аутентификации?","9. Что такое брандмауэр?","10. Какой из следующих методов является примером двухфакторной аутентификации?","11. Какой термин используется для обозначения злонамеренной атаки на компьютерные системы?"],"variant":[["Защита информации","Антивирусная программа","Программа-брандмауэр","Резервное копирование"],["Вирус","Шпионское ПО","Троян","Руткит"],["Аутентификация","Шифрование","Антивирус","Резервное копирование"],["Фишинг","Шифрование","Вирус","Аутентификация"],["Процесс защиты данных путём преобразования их в неразборчивую форму","Процесс проверки подлинности пользователя","Процесс восстановления данных","Процесс создания резервной копии"],["Брандмауэр","Антивирус","Шифрование","Резервное копирование"],["Мошенническое письмо","Антивирусная программа","Сетевой брандмауэр","Процесс восстановления данных"],["Логин и пароль","Резервное копирование","Фишинг","Вирус"],["Программа для защиты от вирусов","Сетевой фильтр, блокирующий несанкционированный доступ","Программа для резервного копирования данных","Процесс восстановления данных"],["Пароль и одноразовый код","Логин и пароль","Шифрование и аутентификация","Резервное копирование и антивирус"],["Кибератака","Брандмауэр","Фишинг","Аутентификация"]],"answers":["Защита информации","Троян","Шифрование","Фишинг","Процесс защиты данных путём преобразования их в неразборчивую форму","Брандмауэр","Мошенническое письмо","Логин и пароль","Сетевой фильтр, блокирующий несанкционированный доступ","Пароль и одноразовый код","Кибератака"]}',
    null,
    2
  );

INSERT INTO public.assignments(
    module_id, course_id, title, description, content, days, assignment_type_id)
VALUES
  (
    30,
    12,
    'Викторина по основам HTML и CSS',
    'Эта викторина проверит ваши знания по основам HTML и CSS.',
    '{"questions":["1. Что означает HTML?","2. Какой тег используется для создания ссылки в HTML?","3. Какой тег используется для вставки изображения в HTML?","4. Какой атрибут используется для указания URL изображения в теге img?","5. Какой тег используется для создания абзаца в HTML?","6. Что означает CSS?","7. Какой тег используется для добавления CSS в HTML-документ?","8. Какой атрибут используется для добавления CSS стилей к элементу?","9. Какой символ используется для указания идентификатора (ID) в CSS?","10. Какой символ используется для указания класса в CSS?","11. Какой свойство CSS используется для изменения цвета текста?","12. Какой свойство CSS используется для изменения фона элемента?","13. Какой свойство CSS используется для изменения размера шрифта?","14. Какой тег используется для создания нумерованного списка в HTML?","15. Какой тег используется для создания ненумерованного списка в HTML?"],"variant":[["Hyper Text Markup Language","Home Tool Markup Language","Hyperlinks and Text Markup Language","Home Text Markup Language"],["<a>","<link>","<href>","<url>"],["<img>","<src>","<image>","<picture>"],["src","href","alt","url"],["<p>","<div>","<span>","<br>"],["Cascading Style Sheets","Color and Style Sheets","Creative Style Sheets","Computer Style Sheets"],["<style>","<css>","<link>","<head>"],["style","css","class","id"],["#","*","@","."],[".","#","*","@"],["color","font-color","text-color","background-color"],["background","bgcolor","bg-color","background-color"],["font-size","text-size","font","size"],["<ol>","<ul>","<li>","<list>"],["<ul>","<ol>","<li>","<list>"]],"answers":["Hyper Text Markup Language","<a>","<img>","src","<p>","Cascading Style Sheets","<style>","style","#",".","color","background","font-size","<ol>","<ul>"]}',
    null,
    2
  );

