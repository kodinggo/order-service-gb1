
-- +migrate Up
CREATE TABLE order_items (
    id INT NOT NULL  PRIMARY KEY AUTO_INCREMENT,
    order_id INT NOT NULL, 
    product_id INT NOT NULL,
    product_name VARCHAR(255),
    product_price FLOAT DEFAULT 0.0,
    qty INT DEFAULT 0 , 
    sub_total FLOAT DEFAULT 0.0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);


-- +migrate Down
DROP TABLE order_items;