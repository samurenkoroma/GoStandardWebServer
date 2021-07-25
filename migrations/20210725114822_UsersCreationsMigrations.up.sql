CREATE TABLE users
(
    id            bigserial    not null primary key,
    login         varchar(250) not null,
    password_hash varchar      not null
);

CREATE TABLE articles
(
    id      bigserial not null primary key,
    title   varchar   not null unique,
    author  varchar   not null,
    content text
);