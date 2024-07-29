CREATE TABLE users(
    id       uuid PRIMARY KEY,
    nickname text NOT NULL UNIQUE,
    email    text NOT NULL UNIQUE
);