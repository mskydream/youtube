begin;
create table if not exists user_profile(
    id bigserial primary key,
    first_name TEXT not null,
    last_name TEXT not null,
    email TEXT unique not null,
    gender text not null,
    pass TEXT not null,
    created_at timestamp without time zone not null
);
commit;