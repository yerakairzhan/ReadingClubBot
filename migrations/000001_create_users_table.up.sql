create table Users(
    id serial primary key,
    user_id varchar unique not null,
    username varchar not null,
    created_at timestamp default current_timestamp
);

create table admins (
    id SERIAL PRIMARY KEY,
    user_id varchar NOT NULL UNIQUE
);

