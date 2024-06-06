create table products
(
    id          int auto_increment
        primary key,
    title       varchar(45)          not null,
    description varchar(255)         null,
    quantity    int                  null,
    price       float                not null,
    isActive    tinyint(1) default 1 not null,
    constraint id_UNIQUE
        unique (id)
);

INSERT INTO go_tp.products (id, title, description, quantity, price, isActive) VALUES (13, 'Test1', 'Dpzmdmdz', 19, 21, true);
INSERT INTO go_tp.products (id, title, description, quantity, price, isActive) VALUES (14, 'Test2', 'Dkzdkzkdzd', 23, 10, true);
INSERT INTO go_tp.products (id, title, description, quantity, price, isActive) VALUES (15, 'Test3', 'Dzkdkzdk', 0, 10, true);
INSERT INTO go_tp.products (id, title, description, quantity, price, isActive) VALUES (16, 'Test4', 'dzdzd', 0, 21, false);
INSERT INTO go_tp.products (id, title, description, quantity, price, isActive) VALUES (17, 'Test5', 'Description', 0, 7, true);
INSERT INTO go_tp.products (id, title, description, quantity, price, isActive) VALUES (18, 'Test6', 'Dhzhz kdzkdzkd dzd', 20, 10, true);
INSERT INTO go_tp.products (id, title, description, quantity, price, isActive) VALUES (22, 'Maxime', 'Sympa', 99880, 1000, true);
INSERT INTO go_tp.products (id, title, description, quantity, price, isActive) VALUES (23, 'Dzer', 'Dzdzd idzd z', 900, 23, true);
INSERT INTO go_tp.products (id, title, description, quantity, price, isActive) VALUES (24, 'Nouveau', 'roduit', 250, 2, true);
INSERT INTO go_tp.products (id, title, description, quantity, price, isActive) VALUES (25, 'Nou Pro', 'Toujous description', 280, 1, false);
