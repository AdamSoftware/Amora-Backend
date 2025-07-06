create table if not exists subscriptions (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id int not null references users(id) on delete cascade,
    plan_id int not null references plans(id) on delete cascade,
    status text not null check (status in ('active', 'inactive', 'cancelled')),
    start_date timestamptz not null,
    next_payment timestamptz not null,
    cancelled_at timestamptz
);
