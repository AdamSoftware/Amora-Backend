create table if not exists venueslist (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name varchar(255) not null,
    address varchar(255) not null,
    city varchar(100) not null,
    state varchar(100) not null,
    zip_code varchar(20) not null,
    phone_number varchar(20),
    email varchar(100),
    website varchar(255),
    image_url varchar(255),
    is_active boolean not null default true
);
