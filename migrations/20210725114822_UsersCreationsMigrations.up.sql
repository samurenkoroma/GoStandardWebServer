CREATE TABLE users
(
    id            uuid         not null primary key,
    name          varchar(250) not null,
    email         varchar(250) not null unique,
    age           integer      not null,
    password_hash varchar      not null
);