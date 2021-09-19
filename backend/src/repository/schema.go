package dbRepo

const Schema = `
CREATE TABLE IF NOT EXISTS phone_numbers
(
    id INTEGER UNIQUE NOT NULL,
    country TEXT NOT NULL,
	state INTEGER NOT NULL,
	code TEXT NOT NULL,
    number TEXT UNIQUE NOT NULL ,
	CONSTRAINT phone_numbers_PK PRIMARY KEY (id)
);`
