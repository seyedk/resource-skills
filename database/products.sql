set time zone 'UTC';
create extension pgcrypto;

CREATE TABLE products (
    id serial PRIMARY KEY,
    name VARCHAR (255) NOT NULL UNIQUE,
    teaser VARCHAR(255) NULL,
    description TEXT NULL,
    price INT NOT NULL,
    image TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE TABLE parts (
    id serial PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE TABLE products_parts (
    id serial PRIMARY KEY,
    product_id int references products(id),
    part_id int references parts(id),
    quantity int NOT NULL,
    unit VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT unique_products_part UNIQUE (product_id,part_id)
);
CREATE TABLE users (
    id serial PRIMARY KEY,
    username VARCHAR (255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE TABLE tokens (
    id serial PRIMARY KEY,
    user_id int references users(id),
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE TABLE orders (
    id serial PRIMARY KEY,
    user_id int references users(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE TABLE order_items (
    id serial PRIMARY KEY,
    order_id int references orders(id),
    product_id int references products(id),
    quantity int NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);


insert into parts (id, name, created_at, updated_at, deleted_at) values (5, 'Tow hitch', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into parts (id, name, created_at, updated_at, deleted_at) values (2, 'Autopilot', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into parts (id, name, created_at, updated_at, deleted_at) values (4, 'Alloyed wheels', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into parts (id, name, created_at, updated_at, deleted_at) values (1, 'Self-Driving', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into parts (id, name, created_at, updated_at, deleted_at) values (3, 'Apple/Andriod Car', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);

insert into products (id, name, teaser, description, price, image, created_at, updated_at, deleted_at) values (1, 'Tesla Model X Performance', 'Performance X', '', 350, '/modelx.png', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products (id, name, teaser, description, price, image, created_at, updated_at, deleted_at) values (5, 'Tesla Model Y Long Range', 'Peformance Y', '', 200, '/modely.png', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products (id, name, teaser, description, price, image, created_at, updated_at, deleted_at) values (4, 'Tesla Model S performance', 'Performance S', '', 150, '/models.png', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products (id, name, teaser, description, price, image, created_at, updated_at, deleted_at) values (2, 'Tesla Model Y Performance', 'LR Y', '', 200, '/modely.png', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products (id, name, teaser, description, price, image, created_at, updated_at, deleted_at) values (6, 'Tesla Model 3 Long Range', 'LR 3', '', 250, '/model3.png', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products (id, name, teaser, description, price, image, created_at, updated_at, deleted_at) values (3, 'Telsa Model 3 Performance', 'Performance 3', '', 150, '/model3.png', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);


insert into products_parts (id, product_id, part_id, quantity, unit, created_at, updated_at, deleted_at) values (7, 3, 3, 3, 'cnt', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products_parts (id, product_id, part_id, quantity, unit, created_at, updated_at, deleted_at) values (6, 3, 1, 1, 'cnt', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products_parts (id, product_id, part_id, quantity, unit, created_at, updated_at, deleted_at) values (9, 5, 1, 2, 'cnt', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products_parts (id, product_id, part_id, quantity, unit, created_at, updated_at, deleted_at) values (8, 4, 1, 1, 'cnt', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products_parts (id, product_id, part_id, quantity, unit, created_at, updated_at, deleted_at) values (10, 6, 1, 3, 'cnt', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products_parts (id, product_id, part_id, quantity, unit, created_at, updated_at, deleted_at) values (5, 2, 2, 3, 'cnt', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products_parts (id, product_id, part_id, quantity, unit, created_at, updated_at, deleted_at) values (11, 6, 5, 4, 'cnt', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products_parts (id, product_id, part_id, quantity, unit, created_at, updated_at, deleted_at) values (1, 1, 1, 2, 'cnt', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products_parts (id, product_id, part_id, quantity, unit, created_at, updated_at, deleted_at) values (2, 1, 2, 1, 'cnt', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products_parts (id, product_id, part_id, quantity, unit, created_at, updated_at, deleted_at) values (3, 1, 4, 2, 'cnt', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);
insert into products_parts (id, product_id, part_id, quantity, unit, created_at, updated_at, deleted_at) values (4, 2, 1, 1, 'cnt', '2021-09-23 00:00:00.000000', '2021-09-23 00:00:00.000000', null);

insert into users (id, username, password, created_at, updated_at, deleted_at) values (2, 'seyedk', 'seyed1234', '2021-09-22 18:06:45.000000', '2021-09-22 18:07:37.000000', null);
insert into users (id, username, password, created_at, updated_at, deleted_at) values (3, 'jimmyn', 'jimmyn1234', '2021-09-22 18:07:01.000000', '2021-09-22 18:07:40.000000', null);
insert into users (id, username, password, created_at, updated_at, deleted_at) values (4, 'edwing', 'edwing1234', '2021-09-22 18:07:25.000000', '2021-09-22 18:07:41.000000', null);
