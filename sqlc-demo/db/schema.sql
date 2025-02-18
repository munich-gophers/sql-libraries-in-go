
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age BIGINT NOT NULL,
    birthday timestamp, 
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

