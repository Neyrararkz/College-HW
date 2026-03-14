-- CREATE TABLE categories(
-- 	id SERIAL PRIMARY KEY,
-- 	title VARCHAR(50) NOT NULL UNIQUE
-- );
-- CREATE TABLE products(
-- 	id SERIAL PRIMARY KEY,
-- 	title VARCHAR(80) NOT NULL,
-- 	price INT NOT NULL CHECK(price>0),
-- 	category_id INT REFERENCES categories(id)
-- );
-- CREATE TABLE orders(
-- 	id SERIAL PRIMARY KEY,
-- 	customer VARCHAR(60) NOT NULL,
-- 	total INT NOT NULL CHECK(total>=0),
-- 	created_at TIMESTAMP DEFAULT NOW()
-- )

-- INSERT INTO categories (title)
-- VALUES 
-- ('Electronics'),
-- ('Accessories'),
-- ('Office'),
-- ('Gaming');

-- INSERT INTO products (title, price, category_id)
-- VALUES
-- ('Keyboard', 12000, 1),
-- ('Mouse', 7000, 1),
-- ('Headphones', 35000, 1),
-- ('USB Cable', 3000, 2),
-- ('Mouse Pad', 3000, 2),
-- ('NoteBook', 1000, 3),
-- ('Pen', 600, 3),
-- ('Gaming chair', 100000, 4),
-- ('Gamepad', 18000, 4),
-- ('Monitor', 100000, 4);

-- INSERT INTO orders (customer, total)
-- VALUES
-- ('Alex', 19000),
-- ('Dana', 35000),
-- ('Diana', 90000),
-- ('Ira', 2500),
-- ('John', 68000),
-- ('Anna', 1500);

-- 1)
-- INSERT INTO categories (title)
-- VALUES ('Smart Home');

-- INSERT INTO products (title, price, category_id)
-- VALUES
-- ('Smart Bulb', 9000, 5),
-- ('Smart Plug', 7000, 5),
-- ('Webcam', 22000, 1),
-- ('Desk Lamp', 8000, 3),
-- ('HDMI Cable', 4000, 2);

-- INSERT INTO orders (customer, total)
-- VALUES
-- ('Alex', 85000),
-- ('Dana', 12000),
-- ('Alex', 45000),
-- ('Ira', 2000),
-- ('Dana', 90000),
-- ('Vadim', 5000);

-- 2)
-- 1
-- SELECT category_id, COUNT(*) AS total_products
-- FROM products
-- GROUP BY category_id
-- ORDER BY total_products DESC

-- 2
-- SELECT category_id, COUNT(*) AS total_products
-- FROM products
-- GROUP BY category_id
-- HAVING COUNT(*) >= 3

-- 3
-- SELECT category_id, MIN(price) AS min_price, MAX(price) AS max_price, ROUND(AVG(price), 2) AS avg_price
-- FROM products
-- GROUP BY category_id

-- 4
-- SELECT category_id, COUNT(*) AS expensive_count
-- FROM products
-- WHERE price > 10000
-- GROUP BY category_id
-- HAVING COUNT(*) >= 2

-- 5
-- SELECT customer, COUNT(*) AS orders_count
-- FROM orders
-- GROUP BY customer
-- ORDER BY orders_count DESC

-- 6
-- SELECT customer, COUNT(*) AS orders_count
-- FROM orders
-- GROUP BY customer
-- HAVING COUNT(*) >= 2

-- 7
-- SELECT customer, SUM(total) AS spent_total, ROUND(AVG(total), 2) AS avg_check
-- FROM orders
-- GROUP BY customer
-- ORDER BY spent_total DESC


