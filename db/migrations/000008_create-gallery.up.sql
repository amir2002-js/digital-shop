CREATE TABLE IF NOT EXISTS gallery (
    id BIGSERIAL PRIMARY KEY ,
    url TEXT NOT NULL ,
    is_main BOOLEAN NOT NULL DEFAULT FALSE,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_gallery_active ON gallery(deleted_at) WHERE deleted_at IS NULL ;
