create table if not exists subscriptions (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id int not null references users(id) on delete cascade,
    plan_id int not null references plans(id) on delete cascade,
    status text not null check (status in ('active', 'inactive', 'cancelled', 'past_due')),
    start_date timestamptz,
    next_payment timestamptz,
    cancelled_at timestamptz,
    payed boolean not null default false
);
