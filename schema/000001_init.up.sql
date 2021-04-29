create table if not exists users (
  id serial not null unique,
  name varchar(255) not null,
  username varchar(255) not null unique,
  password_hash varchar(255) not null
);
create table if not exists todo_lists (
  id serial not null unique,
  title varchar(255) not null,
  description varchar(255)
);
create table if not exists users_lists (
  id serial not null unique,
  user_id int references users (id) on delete cascade not null,
  list_id int references todo_lists (id) on delete cascade not null
);
create table if not exists todo_items (
  id serial not null unique,
  title varchar(255) not null,
  description varchar(255),
  done boolean not null default false
);
create table if not exists lists_items (
  id serial not null unique,
  item_id int references todo_items (id) on delete cascade not null,
  list_id int references todo_lists (id) on delete cascade not null
);