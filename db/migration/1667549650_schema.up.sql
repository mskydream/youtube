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

create table if not exists youtube_channel(
    id bigserial PRIMARY key,
    user_id bigint not null REFERENCES user_profile(id),
    channel_name TEXT NOT NULL UNIQUE,
    created_at timestamp without time zone not null
);

commit;