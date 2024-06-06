create table orders
(
    id          int auto_increment
        primary key,
    customer_id int      not null,
    product_id  int      not null,
    quantity    int      null,
    price       float    not null,
    createdAt   datetime null,
    constraint fk_customer_id
        foreign key (customer_id) references customers (Id),
    constraint fk_product_id
        foreign key (product_id) references products (id)
);

INSERT INTO go_tp.orders (id, customer_id, product_id, quantity, price, createdAt) VALUES (1, 4, 13, 10, 210, '2024-06-06 14:23:25');
INSERT INTO go_tp.orders (id, customer_id, product_id, quantity, price, createdAt) VALUES (2, 4, 13, 2, 42, '2024-06-06 14:26:17');
INSERT INTO go_tp.orders (id, customer_id, product_id, quantity, price, createdAt) VALUES (3, 9, 22, 10, 10000, '2024-06-06 16:07:35');
INSERT INTO go_tp.orders (id, customer_id, product_id, quantity, price, createdAt) VALUES (4, 4, 22, 10, 10000, '2024-06-06 16:13:20');
INSERT INTO go_tp.orders (id, customer_id, product_id, quantity, price, createdAt) VALUES (5, 4, 22, 10, 10000, '2024-06-06 16:15:17');
INSERT INTO go_tp.orders (id, customer_id, product_id, quantity, price, createdAt) VALUES (6, 4, 22, 100, 100000, '2024-06-06 16:21:02');
INSERT INTO go_tp.orders (id, customer_id, product_id, quantity, price, createdAt) VALUES (7, 12, 25, 20, 20, '2024-06-06 17:18:14');
INSERT INTO go_tp.orders (id, customer_id, product_id, quantity, price, createdAt) VALUES (8, 13, 24, 20, 40, '2024-06-06 17:20:05');
INSERT INTO go_tp.orders (id, customer_id, product_id, quantity, price, createdAt) VALUES (9, 13, 24, 30, 60, '2024-06-06 17:21:44');
