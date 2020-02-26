CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY,
  first_name VARCHAR(40),
  last_name VARCHAR(40),
  patronymic VARCHAR(40)
);