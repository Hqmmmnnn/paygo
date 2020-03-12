CREATE TABLE IF NOT EXISTS users (
	id uuid REFERENCES accounts(id) ON UPDATE CASCADE ON DELETE CASCADE,
	first_name VARCHAR(40),
	last_name VARCHAR(40),
	patronymic VARCHAR(40)
);