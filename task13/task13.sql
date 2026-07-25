create table users(
id serial primary key,
name text not null,
email text unique not null,
age integer,
is_active boolean,
created_at timestamp default now()
);

create table orders(
id int primary key,
user_id serial references users(id),
product text not null,
price  NUMERIC(10, 2),
created_at timestamp default now());

INSERT INTO users (name, email, age, is_active)
VALUES
('Иван Иванов', 'ivan@example.com', 25, TRUE),
('Анна Петрова', 'anna@example.com', 30, TRUE),
('Сергей Сидоров', 'sergey@example.com', 19, FALSE),
('Мария Смирнова', 'maria@example.com', 27, TRUE),
('Алексей Кузнецов', 'alexey@example.com', 35, TRUE),
('Елена Орлова', 'elena@example.com', 28, TRUE),
('Дмитрий Волков', 'dmitry@example.com', 33, FALSE),
('Ольга Морозова', 'olga@example.com', 24, TRUE),
('Максим Лебедев', 'maxim@example.com', 41, TRUE),
('Наталья Федорова', 'natalia@example.com', 29, TRUE);

INSERT INTO orders (id, user_id, product, price)
VALUES
(1, 1, 'Ноутбук', 89999.99),
(2, 1, 'Мышь', 1999.00),
(3, 2, 'Клавиатура', 5499.00),
(4, 2, 'Монитор', 24999.99),
(5, 3, 'Наушники', 7999.90),
(6, 4, 'Смартфон', 65999.00),
(7, 4, 'Чехол', 1499.00),
(8, 5, 'SSD 1TB', 8999.99),
(9, 6, 'Процессор', 31999.00),
(10, 7, 'Материнская плата', 18999.00),
(11, 8, 'Оперативная память', 9999.99),
(12, 9, 'Блок питания', 7499.00),
(13, 10, 'Видеокарта', 72999.00),
(14, 10, 'Монитор', 29999.99),
(15, 5, 'Игровая мышь', 3499.00);


select users.name, users.age, orders.product, orders.price
from users join orders on users.id = orders.user_id where orders.price >= 10000;

