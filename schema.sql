CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    password  TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT Now(),
    updated_at TIMESTAMP NOT NULL DEFAULT Now()
);

CREATE TABLE threads (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT Now(),
    updated_at TIMESTAMP NOT NULL DEFAULT Now()
);

CREATE TABLE posts (
    id BIGSERIAL PRIMARY KEY,
    thread_id INTEGER REFERENCES threads(id),
    author_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT Now(),
    updated_at TIMESTAMP NOT NULL DEFAULT Now()
);