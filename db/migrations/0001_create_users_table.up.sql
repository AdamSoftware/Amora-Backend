create table if not exists users (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email varchar(255) not null unique,
    password varchar(255) not null,
    FirstName varchar(100) not null,
    LastName varchar(100) not null,
    phone_number varchar(20) not null unique,
    address varchar(255) not null,
    city varchar(100) not null,
    state varchar(100) not null,
    zip_code varchar(20) not null,
    is_active boolean not null default true
);
