CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY ,
    name VARCHAR(225) NOT NULL ,
    price_before_off NUMERIC(100,2) NOT NULL ,
    price_after_off NUMERIC(100,2) NOT NULL ,
    off NUMERIC(4,2) DEFAULT 0.0,
    img_url TEXT[] NOT NULL DEFAULT '{}',
    CHECK ( ARRAY_LENGTH(img_url,1) BETWEEN 1 AND 5 ),
    description TEXT NOT NULL ,
    stock INT NOT NULL ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,

    category_id INT REFERENCES categories(id) ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_products_active ON products(deleted_at) WHERE deleted_at IS NULL ;