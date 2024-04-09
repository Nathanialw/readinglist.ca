CREATE TABLE categories (
    id serial primary key,
    category varchar(255),
    image varchar(255)
);

CREATE TABLE booklists (
    id serial primary key,
    uuid varchar(64) not null unique,
    name varchar(255),
    author integer references users(id),
    publish_date timestamp not null
);