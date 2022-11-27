-- Write your migrate up statements here

create table users(
    id uuid primary key
);

create table url_link(
    id uuid primary key,
    long  text not null,
    code varchar(10) unique,
    user_id uuid  constraint url_link_users_foreign references users on delete CASCADE null
);

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
