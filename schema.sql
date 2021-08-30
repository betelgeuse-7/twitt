-- Postgres database schema for twitt project
create table users (
	id serial primary key,
	username varchar(255) not null unique,
	password text not null,
	email text not null unique,
	handle varchar(255) not null unique,
	register_date date not null default current_date,
	location varchar(255),
	bio text
);

create table tweets (
	id serial primary key,
	content varchar(280) not null,
	author int references users(id),
	date timestamp not null default localtimestamp
);

create table comments (
	id serial primary key,
	content varchar(280) not null,
	tweet int references tweets(id),
	author int references users(id),
	date timestamp not null default localtimestamp
);

create table follows (
	user_id int references users(id),
	follower_id int references users(id),
	follow_date timestamp not null default localtimestamp,
	unfollow_date timestamp,
	primary key(user_id, follower_id)
);
-- // TODO is it right?
create table likes (
	tweet_id int references tweets(id),
	who_liked int references users(id)string
);