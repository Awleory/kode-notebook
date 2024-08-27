CREATE TABLE users (
	id serial primary key,
	email varchar(255) not null unique,
	password_hash varchar(255) not null
);

CREATE TABLE notes(
	id serial primary key,
	owner_id integer,
	foreign key (owner_id) references users (id) on delete cascade,
	title varchar(255) not null,
	text text
);