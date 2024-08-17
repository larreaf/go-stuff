CREATE TABLE users(
    id uuid,
    created_at timestamp,
    updated_at timestamp,
    email VARCHAR(255),
    password_hash VARCHAR(255),
    user_status int,
    user_role VARCHAR(25)
);

CREATE TABLE book(
    id uuid,
    created_at timestamp,
    updated_at timestamp,
    user_id uuid,
    title VARCHAR(255),
    author VARCHAR(255),
    book_status int,
    book_attrs json
);