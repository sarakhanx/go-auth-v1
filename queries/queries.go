package queries

const CreateUsersTable = `
CREATE TABLE IF NOT EXISTS Users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	roles VARCHAR(255) NOT NULL
)
`

const SignupNewUser = `INSERT INTO Users(name, username, password, email, roles) VALUES($1, $2, $3, $4, $5)`
