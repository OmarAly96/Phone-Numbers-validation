package dbRepo

const Schema = `
CREATE TABLE IF NOT EXISTS phone_numbers
(
    id INTEGER UNIQUE NOT NULL,
    country TEXT NOT NULL,
	state BOOLEAN NOT NULL,
	code TEXT NOT NULL,
    number TEXT NOT NULL,
	CONSTRAINT phone_numbers_PK PRIMARY KEY (id)
);`

//DROP TABLE IF EXISTS phone_numbers;

// CREATE TABLE IF NOT EXISTS customer
// (
//     id INTEGER UNIQUE NOT NULL,
//     name TEXT NOT NULL,
//     phone TEXT NOT NULL,
// 	CONSTRAINT customer_PK PRIMARY KEY (id)
// );
