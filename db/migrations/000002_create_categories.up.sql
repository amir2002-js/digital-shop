CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL PRIMARY KEY ,
    name VARCHAR(120) NOT NULL UNIQUE ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_categories_active ON categories(deleted_at) WHERE deleted_at IS NULL ;