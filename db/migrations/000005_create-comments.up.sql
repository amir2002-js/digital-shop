CREATE TABLE IF NOT EXISTS comments(
    id BIGSERIAL PRIMARY KEY ,
    description TEXT NOT NULL ,
    title VARCHAR(120) NOT NULL ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE ,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX idx_comments_user ON comments(user_id);
CREATE INDEX idx_comments_product ON comments(product_id);
CREATE INDEX idx_comments_active ON comments(deleted_at) WHERE deleted_at IS NULL;