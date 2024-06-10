INSERT INTO public.assignments(
    module_id, course_id, title, description, content, days, assignment_type_id)
VALUES (1, 1, 'Викторина по типам данных', 'Данная викторина проверит ваши знания по типам данных в го', 
'{"questions":["1. Какое значение по умолчанию имеет тип int в Go?","2. Какой из следующих строковых литералов является допустимым в Go?","3. Каков размер float32 в Go?","4. Какой из следующих типов данных может использоваться для представления значений true или false?","5. Как объявить константу в Go?"],"variant":[["0","nil","false","undefined"],["''Hello''","\"Hello\"","`Hello`","Все вышеперечисленные"],["16 бит","32 бит","64 бит","128 бит"],["int","string","bool","byte"],["var constName = value","const constName value","const constName = value","constant constName = value"]],"answers":["0","Все вышеперечисленные","32 бит","bool","const constName = value"]}', 
null, 2);

