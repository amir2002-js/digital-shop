CREATE TABLE IF NOT EXISTS products_tags(
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE ,
    tag_id INT NOT NULL REFERENCES tags(id) ON DELETE CASCADE ,
    PRIMARY KEY (product_id , tag_id)
);
