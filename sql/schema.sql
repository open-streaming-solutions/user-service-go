CREATE TABLE users (
    id       uuid PRIMARY KEY,
    username text NOT NULL UNIQUE,
    email    text NOT NULL UNIQUE
);