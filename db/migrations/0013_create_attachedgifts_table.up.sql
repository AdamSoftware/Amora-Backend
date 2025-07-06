create table if not exists attachedgifts (
    id int generated always as identity primary key,
    event_id int not null references events(id) on delete cascade,
    gift_id int not null references gifts(id) on delete cascade,
    quantity int not null default 1,
    is_active boolean not null default true
);
