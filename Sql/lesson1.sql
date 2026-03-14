-- A. CREATING A TABLE
-- CREATE TABLE employees_hw (
--   id BIGSERIAL PRIMARY KEY,
--   full_name VARCHAR(120) NOT NULL,
--   email VARCHAR(120) UNIQUE NOT NULL,
--   age INT CHECK (age >= 0),
--   department VARCHAR(60) NOT NULL,
--   salary INT CHECK (salary >= 0),
--   hired_at DATE DEFAULT CURRENT_DATE,
--   fired_at DATE,
--   created_at TIMESTAMPTZ DEFAULT NOW()
-- );


-- B. INSERT
-- INSERT INTO employees_hw (full_name, email, age, department, salary, hired_at, fired_at) 
-- VALUES
-- 	('Anna Petrova', 'anna.petrova@gmail.com', 19, 'IT', 650000, '2023-05-10', NULL),
-- 	('Ivan Sidorov', 'ivan.sidorov@corp.com', 28, 'Sales', 420000, '2022-11-01', NULL),
-- 	('Diana Smirnova', 'diana.smirnova@gmail.com', 24, 'Marketing', 380000, '2023-03-15', NULL),
-- 	('Roman Kuznetsov', 'roman.kuznetsov@corp.com', 31, 'IT', 720000, '2021-09-20', NULL),
-- 	('Andrey Volkov', 'andrey.volkov@gmail.com', 22, 'HR', 310000, '2022-02-14', NULL),
-- 	('Elena Morozova', 'elena.morozova@corp.com', 27, 'Marketing', 610000, '2023-06-05', NULL),
-- 	('Nikita Ivanov', 'nikita.ivanov@gmail.com', 18, 'Sales', 250000, '2024-01-10', NULL),
-- 	('Arman Bekov', 'arman.bekov@corp.com', 19, 'IT', 600000, '2023-07-01', NULL),
-- 	('Danil Popov', 'danil.popov@gmail.com', 21, 'Sales', 450000, '2022-12-12', NULL),
-- 	('Sergey Lebedev', 'sergey.lebedev@corp.com', 35, 'IT', 800000, '2020-04-18', NULL),
-- 	('Alina Romanova', 'alina.romanova@gmail.com', 23, 'Marketing', 340000, '2023-08-22', '2024-05-10'),
-- 	('Anton Pavlov', 'anton.pavlov@corp.com', 29, 'HR', 290000, '2021-11-03', '2024-02-01'),
-- 	('Anastasia Volkova', 'anastasia.volkova@gmail.com', 26, 'Sales', 520000, '2022-06-30', NULL),
-- 	('Maxim Orlov', 'maxim.orlov@corp.com', 33, 'IT', 680000, '2021-03-09', NULL),
-- 	('Timur Akhmetov', 'timur.akhmetov@gmail.com', 25, 'Marketing', 270000, '2023-09-17', '2024-01-15');


-- C. SELECT + WHERE
-- 1
-- SELECT * FROM employees_hw 
-- WHERE salary > 500000

-- 2
-- SELECT * FROM employees_hw 
-- WHERE age >= 18

-- 3
-- SELECT * FROM employees_hw 
-- WHERE department = 'IT'

-- 4
-- SELECT * FROM employees_hw 
-- WHERE department != 'IT'

-- 5
-- SELECT * FROM employees_hw 
-- WHERE id = 10

-- 6
-- SELECT * FROM employees_hw 
-- WHERE department = 'IT' AND salary >= 600000

-- 7
-- SELECT * FROM employees_hw 
-- WHERE department = 'HR' OR department = 'Marketing'

-- 8
-- SELECT * FROM employees_hw 
-- WHERE age < 20 OR salary > 700000

-- 9
-- SELECT * FROM employees_hw 
-- WHERE department = 'Sales' AND (age < 25 OR salary > 500000)

-- 10
-- SELECT * FROM employees_hw 
-- WHERE age BETWEEN 20 AND 30

-- 11
-- SELECT * FROM employees_hw 
-- WHERE salary BETWEEN 300000 AND 600000

-- 12
-- SELECT * FROM employees_hw 
-- WHERE fired_at IS NULL

-- 13
-- SELECT * FROM employees_hw 
-- WHERE fired_at IS NOT NULL

-- D. LIKE/ILIKE
-- 14 
-- SELECT * FROM employees_hw 
-- WHERE full_name ILIKE '%an%'

-- 15
-- SELECT * FROM employees_hw 
-- WHERE full_name LIKE 'A%'

-- 16
-- SELECT * FROM employees_hw 
-- WHERE email ILIKE '%corp%' OR email ILIKE '%gmail%'

-- 17
-- SELECT * FROM employees_hw 
-- WHERE department ILIKE '%s%'


-- E. ORDER BY + LIMIT 
-- 18
-- SELECT * FROM employees_hw 
-- ORDER BY salary DESC

-- 19
-- SELECT * FROM employees_hw 
-- ORDER BY salary DESC
-- LIMIT 5

-- 20
-- SELECT * FROM employees_hw 
-- ORDER BY age ASC
-- LIMIT 3

-- 21
-- SELECT * FROM employees_hw 
-- ORDER BY created_at DESC
-- LIMIT 5

-- 22
-- SELECT * FROM employees_hw 
-- ORDER BY department ASC, full_name ASC


-- F. UPDATE
-- 23
-- UPDATE employees_hw
-- SET salary = salary + salary/10
-- WHERE department = 'IT'
-- SELECT * FROM employees_hw

-- 24
-- UPDATE employees_hw
-- SET salary = salary - salary*5/100
-- WHERE department = 'Sales'
-- SELECT * FROM employees_hw

-- 25
-- UPDATE employees_hw
-- SET department = 'HR'
-- WHERE email = 'alina.romanova@gmail.com'
-- SELECT * FROM employees_hw

-- 26 
-- UPDATE employees_hw
-- SET full_name = 'Vadim Petrov'
-- WHERE id = 4
-- SELECT * FROM employees_hw

-- 27
-- UPDATE employees_hw
-- SET fired_at = '2026-03-14'
-- WHERE id = 2
-- SELECT * FROM employees_hw

-- 28
-- UPDATE employees_hw
-- SET salary = salary + 50000
-- WHERE age <= 18
-- SELECT * FROM employees_hw


-- G. DELETE 
-- 29
-- DELETE FROM employees_hw
-- WHERE id = 15
-- SELECT * FROM employees_hw

-- 30
-- DELETE FROM employees_hw
-- WHERE salary < 200000
-- SELECT * FROM employees_hw

-- 31
-- DELETE FROM employees_hw
-- WHERE fired_at IS NOT NULL
-- SELECT * FROM employees_hw


-- FINAL
-- 32
-- SELECT COUNT (*) FROM employees_hw

-- 33 
-- SELECT MIN(salary), MAX(salary)
-- FROM employees_hw

-- 34
-- SELECT * FROM employees_hw 
-- ORDER BY id
