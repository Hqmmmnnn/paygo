create table if not exists users (
  id uuid primary key,
  email varchar(50) not null unique,
  password varchar not null,
  first_name varchar(40),
  last_name varchar(40),
  patronymic varchar(40)
);