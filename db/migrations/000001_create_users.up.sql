CREATE TYPE role_type AS ENUM ('admin', 'user');
CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY ,
    username VARCHAR(30) NOT NULL ,
    hashedPass VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE ,
    role role_type DEFAULT 'user',
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);