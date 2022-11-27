-- Write your migrate up statements here

create table url_link
(
    id      bigserial primary key,
    long    text not null,
    code    varchar(10) unique,
    user_id uuid null
);

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
