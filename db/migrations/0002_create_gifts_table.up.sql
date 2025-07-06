create table if not exists gifts (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text not null,
    description text,
    image_url text,
    price float not null default 0,
    is_active boolean not null default true
);


