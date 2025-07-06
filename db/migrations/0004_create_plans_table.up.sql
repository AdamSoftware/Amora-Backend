create table if not exists plans (
    id int generated always as identity primary key,
    plantype text not null,
    price float not null default 0
);
