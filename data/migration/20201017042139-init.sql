
-- +migrate Up
CREATE FUNCTION gen_random_uuid() RETURNS uuid
    LANGUAGE c
    AS '$libdir/pgcrypto', 'pg_random_uuid';

create table users (
	id uuid default gen_random_uuid() not null unique primary key,
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	name varchar(200),
	email varchar(200),
	avatar text
);

create table products (
	id serial primary key,
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	name text not null,
	price real not null,
	photo text
);

-- +migrate Down
drop table products;
drop table users;
drop function gen_random_uuid;