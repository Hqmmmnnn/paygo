create table users (
  id uuid primary key,
  email varchar(40) not null unique,
  password varchar(40) not null,
  first_name varchar(40),
  last_name varchar(40),
  patronymic varchar(40)
);