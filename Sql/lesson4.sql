-- CREATE TABLE customers(
-- 	id SERIAL PRIMARY KEY,
-- 	name TEXT NOT NULL,
-- 	phone TEXT UNIQUE,
-- 	city TEXT NOT NULL,
-- 	created_at TIMESTAMP NOT NULL DEFAULT NOW()
-- );
-- CREATE TABLE categories(
-- 	id SERIAL PRIMARY KEY,
-- 	title TEXT NOT NULL UNIQUE
-- );
-- CREATE TABLE products(
-- 	id SERIAL PRIMARY KEY,
-- 	title TEXT NOT NULL,
-- 	category_id INT REFERENCES categories(id) ON DELETE SET NULL,
-- 	price NUMERIC(10,2) NOT NULL CHECK (price > 0),
-- 	is_active BOOLEAN NOT NULL DEFAULT TRUE
-- );

-- CREATE TABLE orders(
-- 	id SERIAL PRIMARY KEY,
-- 	customer_id INT REFERENCES customers(id) ON DELETE SET NULL,
-- 	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
-- 	status TEXT NOT NULL CHECK (status IN ('new','paid','shipped','cancelled'))
-- );

-- CREATE TABLE order_items(
-- 	order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
-- 	product_id INT REFERENCES products(id) ON DELETE SET NULL, 
-- 	qty INT NOT NULL CHECK (qty > 0),
-- 	price NUMERIC(10,2) NOT NULL CHECK (price > 0),
-- 	PRIMARY KEY (order_id, product_id)
-- );
-- CREATE TABLE payments(
-- 	id SERIAL PRIMARY KEY,
-- 	order_id INT REFERENCES orders(id) ON DELETE SET NULL,
-- 	paid_at TIMESTAMP,
-- 	amount NUMERIC(10,2) NOT NULL CHECK(amount > 0),
-- 	method TEXT NOT NULL CHECK(method IN ('card','cash','transfer'))
-- );
-- CREATE TABLE shipments(
-- 	id SERIAL PRIMARY KEY,
-- 	order_id INT REFERENCES orders(id) ON DELETE SET NULL,
-- 	shipped_at TIMESTAMP,
-- 	carrier TEXT NOT NULL CHECK(carrier IN ('glovo','yandex','pickup','courier')),
-- 	tracking TEXT
-- )


-- INSERT INTO customers (id, name, phone, city, created_at) VALUES
-- (1,'Aruzhan','+77010000001','Almaty',   '2026-01-10 10:00'),
-- (2,'Dias',   '+77010000002','Almaty',   '2026-01-11 12:00'),
-- (3,'Madi',   '+77010000003','Astana',   '2026-01-12 09:30'),
-- (4,'Amina',  '+77010000004','Shymkent', '2026-01-13 18:20'),
-- (5,'Ilyas',  '+77010000005','Astana',   '2026-01-20 14:10'),
-- (6,'Saniya', '+77010000006','Almaty',   '2026-01-25 20:45');


-- INSERT INTO categories (id, title) VALUES
-- (1,'Food'),
-- (2,'Electronics'),
-- (3,'Books'),
-- (4,'Clothes'),
-- (5,'Toys'); 


-- INSERT INTO products (id, title, category_id, price, is_active) VALUES
-- (1,'Burger',       1, 2500,  TRUE),
-- (2,'Cola',         1,  600,  TRUE),
-- (3,'Pizza',        1, 4500,  TRUE),
-- (4,'Headphones',   2,19990,  TRUE),
-- (5,'USB-C Cable',  2, 2990,  TRUE),
-- (6,'Laptop',       2,399990, TRUE),
-- (7,'SQL Basics',   3, 5500,  TRUE),
-- (8,'React Guide',  3, 8000,  TRUE),
-- (9,'Hoodie',       4,15990,  TRUE),
-- (10,'Sneakers',    4,45990, TRUE);


-- INSERT INTO orders (id, customer_id, created_at, status) VALUES
-- (101, 1, '2026-02-01 11:00', 'paid'),
-- (102, 1, '2026-02-02 13:10', 'new'),
-- (103, 2, '2026-02-02 16:40', 'shipped'),
-- (104, 3, '2026-02-03 09:15', 'cancelled'),
-- (105, 4, '2026-02-04 19:00', 'paid'),
-- (106, 4, '2026-02-05 10:05', 'shipped'),
-- (107, 6, '2026-02-06 12:00', 'new'),
-- (108, NULL,'2026-02-07 12:30', 'paid'); 


-- INSERT INTO order_items (order_id, product_id, qty, price) VALUES
-- (101, 1, 2, 2500),
-- (101, 2, 2,  600),

-- (102, 3, 1, 4500),

-- (103, 4, 1, 19990),
-- (103, 5, 2,  2990),

-- (104, 7, 1, 5500),

-- (105, 9, 1, 15990),
-- (105, 2, 3,   600),

-- (106, 6, 1, 399990),
-- (106, 5, 1,   2990),

-- (108, 8, 1, 8000);


-- INSERT INTO payments (id, order_id, paid_at, amount, method) VALUES
-- (201, 101, '2026-02-01 11:05', 6200,  'card'),    
-- (202, 105, '2026-02-04 19:10', 17790, 'cash'),     
-- (203, 106, '2026-02-05 10:10', 402980,'transfer'), 
-- (204, 108, '2026-02-07 12:35', 8000,  'card'),
-- (205, NULL,'2026-02-08 10:00', 9999,  'cash');   


-- INSERT INTO shipments (id, order_id, shipped_at, carrier, tracking) VALUES
-- (301, 103, '2026-02-03 10:00', 'glovo',    'GLOVO-103-AAA'),
-- (302, 106, '2026-02-05 18:00', 'courier','CR-106-BBB'),
-- (303, NULL,'2026-02-06 09:00', 'pickup', NULL); 

-- 1
-- SELECT c.id, c.name,
-- 	CASE
-- 		WHEN (SELECT COUNT(*)
-- 			FROM orders o
-- 			WHERE o.customer_id = c.id ) = 0 THEN 'new'
-- 		WHEN (SELECT COUNT(*)
-- 			FROM orders o
-- 			WHERE o.customer_id = c.id ) BETWEEN 1 AND 2 THEN 'returning'
-- 		ELSE 'loyal'
-- 	END AS client_type
-- FROM customers c

-- 2
-- SELECT id, created_at, status,
-- 	CASE 
-- 		WHEN status = 'cancelled' THEN 'high'
-- 		WHEN status = 'new' AND created_at < NOW() - INTERVAL '3 days' THEN 'medium'
-- 		ELSE 'low'
-- 	END AS risk_label
-- FROM orders

-- 3
-- SELECT p.id, p.title, p.price, p.category_id
-- FROM products p
-- WHERE p.price = (
-- 	SELECT MAX(price)
-- 	FROM products
-- 	WHERE category_id = p.category_id)

-- 4
-- SELECT p.id, p.title,
-- 	CASE 
-- 		WHEN EXISTS (
-- 			SELECT 1
-- 			FROM order_items oi
-- 			WHERE oi.product_id = p.id
-- 		) THEN 'sold'
-- 		ELSE 'not_sold'
-- 	END AS sold_flag
-- FROM products p

-- 5
-- SELECT o.id, 
-- 	COALESCE(
--         (SELECT SUM(amount)
--          FROM payments p
--          WHERE p.order_id = o.id), 0
--     ) AS paid_amount,
--     COALESCE(
--         (SELECT SUM(qty * price)
--          FROM order_items oi
--          WHERE oi.order_id = o.id), 0
--     ) AS items_total
-- FROM orders o
-- WHERE 
-- COALESCE(
--     (SELECT SUM(amount)
--      FROM payments p
--      WHERE p.order_id = o.id), 0
-- )
-- >
-- COALESCE(
--     (SELECT SUM(qty * price)
--      FROM order_items oi
--      WHERE oi.order_id = o.id), 0
-- )

-- 6
-- SELECT o.id, o.status
-- FROM orders o
-- WHERE EXISTS (
--     SELECT 1
--     FROM order_items oi
--     WHERE oi.order_id = o.id
--       AND oi.product_id IN (
--           SELECT id
--           FROM products
--           WHERE price > 20000
--       )
-- )

-- 7
-- SELECT o.id,
--     CASE
--         WHEN COALESCE(
--             (SELECT SUM(qty*price)
--              FROM order_items oi
--              WHERE oi.order_id = o.id), 0) = 0 THEN 'zero'
--         WHEN COALESCE(
--             (SELECT SUM(qty*price)
--              FROM order_items oi
--              WHERE oi.order_id = o.id), 0) < 10000 THEN 'small'
--         WHEN COALESCE(
--             (SELECT SUM(qty*price)
--              FROM order_items oi
--              WHERE oi.order_id = o.id),0) <= 50000 THEN 'medium'
--         ELSE 'big'
--     END AS check_class
-- FROM orders o

-- 8
-- SELECT c.id, c.name
-- FROM customers c
-- WHERE EXISTS (
--     SELECT 1
--     FROM orders o
--     WHERE o.customer_id = c.id
--       AND EXISTS (
--           SELECT 1
--           FROM payments p
--           WHERE p.order_id = o.id
--       )
-- ) AND NOT EXISTS (
--     SELECT 1
--     FROM orders o
--     WHERE o.customer_id = c.id
--       AND EXISTS (
--           SELECT 1
--           FROM shipments s
--           WHERE s.order_id = o.id
--       )
-- )

-- 9
-- SELECT o.id,
--     CASE
--         WHEN EXISTS (
--             SELECT 1
--             FROM order_items oi
--             WHERE oi.order_id = o.id
--         ) THEN 'filled'
--         ELSE 'empty'
--     END AS items_state
-- FROM orders o

-- 10
-- SELECT o.id,
--     CASE
--         WHEN NOT EXISTS (
--             SELECT 1
--             FROM order_items oi
--             WHERE oi.order_id = o.id
--         ) THEN 'D'
--         WHEN EXISTS (
--             SELECT 1
--             FROM payments p
--             WHERE p.order_id = o.id
--         ) AND COALESCE(
--             (SELECT SUM(qty*price)
--              FROM order_items oi
--              WHERE oi.order_id = o.id),0
--         ) >= 50000 THEN 'A'  
--         WHEN EXISTS (
--             SELECT 1
--             FROM payments p
--             WHERE p.order_id = o.id
--         ) THEN 'B' 
--         ELSE 'C'
--     END AS business_grade
-- FROM orders o