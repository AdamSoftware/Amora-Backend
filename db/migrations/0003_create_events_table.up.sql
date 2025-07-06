create table if not exists events (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text not null,
    description text,
    DateOfTheEvent date not null,
    is_active boolean not null default true
);
