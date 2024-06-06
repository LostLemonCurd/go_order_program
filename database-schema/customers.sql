create table customers
(
    Id          int auto_increment
        primary key,
    FirstName   varchar(45) not null,
    LastName    varchar(45) not null,
    PhoneNumber varchar(45) null,
    Address     varchar(45) not null,
    Email       varchar(60) not null
);

INSERT INTO go_tp.customers (Id, FirstName, LastName, PhoneNumber, Address, Email) VALUES (3, 'Lounis', 'Ord', '0672772722', '23 rue Mon', 'l.ord@gmail.com');
INSERT INTO go_tp.customers (Id, FirstName, LastName, PhoneNumber, Address, Email) VALUES (4, 'Cheese', 'Cheesios', '0677282828', '45 rue Pompidou', 'lounis.ourrad@hotmail.fr');
INSERT INTO go_tp.customers (Id, FirstName, LastName, PhoneNumber, Address, Email) VALUES (5, 'Pierre', 'Barrette', '0728282234', '106 rue Greille', 'm.petit@gmail.com');
INSERT INTO go_tp.customers (Id, FirstName, LastName, PhoneNumber, Address, Email) VALUES (6, 'Maxime', 'Little', '0728281919', '67 rue Hroem', 'p.little@hotmail.fr');
INSERT INTO go_tp.customers (Id, FirstName, LastName, PhoneNumber, Address, Email) VALUES (7, 'Maxence', 'Non', '0627272828', '89bis rue Mamie', 'fred@gmail.com');
INSERT INTO go_tp.customers (Id, FirstName, LastName, PhoneNumber, Address, Email) VALUES (8, 'Doriane', 'Farandolle', '0672928282', '35 rue Bien Aimée', 'dodo@hotmail.fr');
INSERT INTO go_tp.customers (Id, FirstName, LastName, PhoneNumber, Address, Email) VALUES (9, 'Samuel', 'Auditif', '0292827272', '100 rue Bretagne', 'sam.a@gmail.com');
INSERT INTO go_tp.customers (Id, FirstName, LastName, PhoneNumber, Address, Email) VALUES (12, 'Petit', 'Maxime', '0738383838', '78 rue loin', 'petit@gmail.dev');
INSERT INTO go_tp.customers (Id, FirstName, LastName, PhoneNumber, Address, Email) VALUES (13, 'Pierre', 'Barret', '06728282828', '87 rue dellez', 'pierre@gmail.com');
