begin;
create table if not exists user_profile(
    id bigserial primary key,
    first_name TEXT not null,
    last_name TEXT not null,
    email TEXT unique not null,
    pass TEXT not null,
    gender text check(gender in ('Male','Female')) not null,
    created_at timestamp without time zone not null
);


create table if not exists youtube_account(
    id  bigint PRIMARY key REFERENCES user_profile(id),
    created_at timestamp without time zone not null
);

create table if not exists youtube_channel(
    id bigserial PRIMARY key,
    youtube_account_id bigint not null REFERENCES youtube_account(id),
    channel_name TEXT NOT NULL UNIQUE,
    created_at timestamp without time zone not null
);

create table if not exists channel_subscriber(
    youtube_account_id BIGINT REFERENCES youtube_account(id),
    youtube_channel_id BIGINT REFERENCES youtube_channel(id),
    created_at timestamp without time zone not null,
    primary key (youtube_account_id, youtube_channel_id)
);

commit; 