CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY ,
    username VARCHAR(30) NOT NULL ,
    hashed_pass VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE ,
    role VARCHAR(20) DEFAULT 'user' CHECK (role IN ('admin','user')),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_users_active ON users(deleted_at) WHERE deleted_at IS NULL;