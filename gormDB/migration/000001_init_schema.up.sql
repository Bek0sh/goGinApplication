create table "user" (
    id serial primary key,
    name varchar(40) not null,
    surname varchar(40) not null,
    email varchar(60) not null unique,
    password varchar(60) not null
)