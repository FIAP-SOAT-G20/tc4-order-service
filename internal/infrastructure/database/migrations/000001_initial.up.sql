CREATE TABLE IF NOT EXISTS customers
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR   NOT NULL,
    email      VARCHAR   NOT NULL UNIQUE,
    cpf        VARCHAR   NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS categories
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS staffs
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    role VARCHAR CHECK (role IN ('COOK', 'ATTENDANT', 'MANAGER')),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR        NOT NULL,
    description VARCHAR,
    price       DECIMAL(19, 2) NOT NULL,
    category_id INT            NOT NULL REFERENCES categories (id),
    image_url   VARCHAR,
    staff_id    INT REFERENCES staffs (id),
    active      BOOLEAN                 DEFAULT true,
    created_at  TIMESTAMP      NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP      NOT NULL DEFAULT now()
);

DROP TYPE IF EXISTS order_status;
CREATE TYPE order_status AS ENUM ('OPEN','CANCELLED','PENDING','RECEIVED', 'PREPARING', 'READY', 'COMPLETED');

CREATE TABLE IF NOT EXISTS orders
(
    id          SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers (id),
    status     order_status DEFAULT 'OPEN',
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS order_products
(
    order_id   INT REFERENCES orders (id),
    product_id INT REFERENCES products (id),
    quantity   INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (order_id, product_id)
);

CREATE TABLE IF NOT EXISTS order_histories
(
    id         SERIAL PRIMARY KEY,
    order_id   INT REFERENCES orders (id) NOT NULL,
    staff_id   INT REFERENCES staffs (id) NULL,
    status     order_status DEFAULT 'OPEN',
    created_at TIMESTAMP NOT NULL                                                        DEFAULT now(),
    updated_at TIMESTAMP NOT NULL                                                        DEFAULT now()
);

CREATE TABLE IF NOT EXISTS payments
(
    id                  SERIAL PRIMARY KEY,
    status              VARCHAR CHECK (status IN ('PROCESSING', 'CONFIRMED', 'ABORTED', 'FAILED')) DEFAULT 'PROCESSING',
    external_payment_id VARCHAR,
    order_id            INT REFERENCES orders (id),
    qr_data             VARCHAR,
    created_at          TIMESTAMP NOT NULL                                             DEFAULT now(),
    updated_at          TIMESTAMP NOT NULL                                             DEFAULT now()
);

INSERT INTO categories (name)
VALUES ('Lanches'),
       ('Bebidas'),
       ('Sobremesas'),
       ('Acompanhamentos'),
       ('Combos');

INSERT INTO staffs (name, role)
VALUES ('João Silva', 'COOK'),
       ('Maria Oliveira', 'ATTENDANT'),
       ('Pedro Santos', 'MANAGER'),
       ('Ana Costa', 'COOK'),
       ('Carlos Pereira', 'ATTENDANT');

INSERT INTO products (name, description, price, category_id, image_url, staff_id, active)
VALUES ('X-Burger', 'Hambúrguer com queijo, alface e tomate', 25.90, 1, 'https://example.com/xburger.jpg', 1, true),
       ('Coca-Cola 350ml', 'Refrigerante Coca-Cola lata', 6.90, 2, 'https://example.com/coca.jpg', 2, true),
       ('Sundae', 'Sorvete com calda de chocolate', 12.90, 3, 'https://example.com/sundae.jpg', 1, true),
       ('Batata Frita', 'Porção de batata frita crocante', 15.90, 4, 'https://example.com/batata.jpg', 4, true),
       ('Combo Big', 'X-Burger + Batata + Refrigerante', 42.90, 5, 'https://example.com/combo.jpg', 1, true);

INSERT INTO customers (name, email, cpf)
VALUES ('Lucas Mendes', 'lucas@email.com', '123.456.789-00'),
       ('Julia Santos', 'julia@email.com', '987.654.321-00'),
       ('Rafael Costa', 'rafael@email.com', '456.789.123-00'),
       ('Mariana Lima', 'mariana@email.com', '789.123.456-00'),
       ('Bruno Oliveira', 'bruno@email.com', '321.654.987-00');

INSERT INTO orders (id, customer_id, status, created_at)
VALUES (1, 1, 'OPEN', now()),
       (2, 2, 'PENDING', now()),
       (3, 3, 'CANCELLED', '2021-10-01 10:00:00.467'),
       (4, 4, 'RECEIVED', '2021-10-01 10:00:00.467'),
       (5, 5, 'PREPARING', '2021-10-01 10:00:00.467'),
       (6, 1, 'READY', '2021-10-01 10:00:00.467'),
       (7, 2, 'COMPLETED', '2021-01-01 10:00:00.467'),
       (8, 4, 'RECEIVED', '2021-10-01 11:00:00.467'),
       (9, 4, 'RECEIVED', '2021-10-01 12:00:00.467'),
       (10, 5, 'PREPARING', '2021-10-01 08:00:00.467'),
       (11, 5, 'PREPARING', '2021-10-01 09:00:00.467'),
       (12, 1, 'READY', '2021-10-01 10:00:01.467'),
       (13, 1, 'READY', '2021-10-01 09:59:59.467'),
       (14, 1, 'CANCELLED', '2021-10-01 09:59:59.467');


INSERT INTO order_products (order_id, product_id, quantity)
VALUES (1, 1, 1),
       (1, 2, 1),
       (2, 5, 1),
       (3, 1, 1),
       (4, 1, 1),
       (4, 2, 1),
       (4, 3, 1),
       (4, 4, 1),
       (5, 2, 2),
       (5, 3, 1),
       (6, 1, 1),
       (7, 2, 1),
       (8, 1, 1),
       (9, 2, 1),
       (10, 1, 1),
       (11, 2, 1),
       (12, 1, 1),
       (13, 2, 1);



INSERT INTO order_histories (order_id, staff_id, status)
VALUES (1, null, 'OPEN'),
       (2, null, 'OPEN'),
       (2, null, 'PENDING'),
       (3, null, 'OPEN'),
       (3, null, 'PENDING'),
       (3, null, 'CANCELLED'),
       (4, null, 'OPEN'),
       (4, null, 'PENDING'),
       (4, null, 'RECEIVED'),
       (5, null, 'OPEN'),
       (5, null, 'PENDING'),
       (5, null, 'RECEIVED'),
       (5, 1, 'PREPARING'),
       (6, null, 'OPEN'),
       (6, null, 'PENDING'),
       (6, null, 'RECEIVED'),
       (6, 1, 'PREPARING'),
       (6, 2, 'READY'),
       (7, null, 'OPEN'),
       (7, null, 'PENDING'),
       (7, null, 'RECEIVED'),
       (7, 2, 'PREPARING'),
       (7, 2, 'READY'),
       (7, 2, 'COMPLETED'),
       (8, null, 'OPEN'),
       (8, null, 'PENDING'),
       (8, null, 'RECEIVED'),
       (9, null, 'OPEN'),
       (9, null, 'PENDING'),
       (9, null, 'RECEIVED'),
       (10, null, 'OPEN'),
       (10, null, 'PENDING'),
       (10, null, 'RECEIVED'),
       (10, 1, 'PREPARING'),
       (11, null, 'OPEN'),
       (11, null, 'PENDING'),
       (11, null, 'RECEIVED'),
       (11, 1, 'PREPARING'),
       (12, null, 'OPEN'),
       (12, null, 'PENDING'),
       (12, null, 'RECEIVED'),
       (12, 1, 'PREPARING'),
       (12, 2, 'READY'),
       (13, null, 'OPEN'),
       (13, null, 'PENDING'),
       (13, null, 'RECEIVED'),
       (13, 2, 'PREPARING'),
       (13, 2, 'READY'),
       (14, null, 'OPEN'),
       (14, null, 'PENDING'),
       (14, null, 'CANCELLED');
       
       

INSERT INTO payments (id, status, external_payment_id, order_id, qr_data)
VALUES (1, 'PROCESSING', '09d92b11-cd55-4a72-b2ee-7377ceefe265', 2, 'QR_DATA_345'),
       (2, 'FAILED', 'b7fa4bee-fc25-4bb4-b948-5139af948a39', 3, 'QR_DATA_789'),
       (3, 'CONFIRMED', '5c272292-4ba4-41e9-83d8-dea99afe5194', 4, 'QR_DATA_123'),
       (4, 'CONFIRMED', 'ac174c5e-c9ef-4407-a3b3-bceeb4163af3', 5, 'QR_DATA_456'),
       (5, 'CONFIRMED', '09d92b11-cd55-4a72-b2ee-7377ceefe265', 6, 'QR_DATA_345'),
       (6, 'CONFIRMED', '26e24f2a-5b00-4687-800f-a7be71104b2b', 7, 'QR_DATA_789'),
       (7, 'CONFIRMED', '26e24f2a-5b00-4687-800f-a7be71104b2c', 8, 'QR_DATA_790'),
       (8, 'CONFIRMED', '26e24f2a-5b00-4687-800f-a7be71104b2d', 9, 'QR_DATA_791'),
       (9, 'CONFIRMED', '26e24f2a-5b00-4687-800f-a7be71104b2e', 10, 'QR_DATA_792'),
       (10, 'CONFIRMED', '26e24f2a-5b00-4687-800f-a7be71104b2f', 11, 'QR_DATA_793'),
       (11, 'CONFIRMED', '26e24f2a-5b00-4687-800f-a7be71104b2g', 12, 'QR_DATA_794'),
       (12, 'CONFIRMED', '26e24f2a-5b00-4687-800f-a7be71104b2h', 13, 'QR_DATA_795'),
       (13, 'ABORTED', '26e24f2a-5b00-4687-800f-a7be71104b2i', 14, 'QR_DATA_796');
       

SELECT setval('categories_id_seq', (SELECT MAX(id) FROM categories));
SELECT setval('staffs_id_seq', (SELECT MAX(id) FROM staffs));
SELECT setval('products_id_seq', (SELECT MAX(id) FROM products));
SELECT setval('customers_id_seq', (SELECT MAX(id) FROM customers));
SELECT setval('orders_id_seq', (SELECT MAX(id) FROM orders));
SELECT setval('order_histories_id_seq', (SELECT MAX(id) FROM order_histories));
SELECT setval('payments_id_seq', (SELECT MAX(id) FROM payments));
