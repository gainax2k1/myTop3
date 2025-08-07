# myTop3
Social networking based on similar interests.

(Capstone project for boot.dev)

Designed to use REST API.

Users will be able to list their "top 3" of a subject, and be connected to other users who share similar lists. Initially, I'm developing it with just videogames in mind for scope, but the general design could be developed to grow into multiple or different subjects, such as "myTop3: Music", "myTop3: Movies", etc.

- uses postgres

sudo apt update
sudo apt install postgresql postgresql-contrib

sudo passwd postgres
(create system user's password)

sudo -u postgres psql
(opens postgres)

CREATE DATABASE mytop3;
(creates database)


\c mytop3
(connects to the database just created)

ALTER USER postgres WITH PASSWORD 'postgres';
(set the database user's password)

- uses Goose for migrations

go install github.com/pressly/goose/v3/cmd/goose@latest

determine connection string:
(change to appropriate user:pass, host, port, db name)
"postgres://postgres:postgres@localhost:5432/mytop3"

then, goose the migrations in sql/schema:
goose postgres (connectionstring from above) up

- uses SQLC to generate Go code from SQL queries to interface with our database
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest