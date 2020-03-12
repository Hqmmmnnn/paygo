CREATE TABLE IF NOT EXISTS accounts (
	id uuid PRIMARY KEY,
	email VARCHAR(40) NOT NULL UNIQUE,
	login VARCHAR(40) NOT NULL UNIQUE,
	password VARCHAR NOT NULL,
	balance NUMERIC(20, 2) DEFAULT 0.00 CONSTRAINT not_negative_balance CHECK (balance >= 0),
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);