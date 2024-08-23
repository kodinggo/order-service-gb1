
-- +migrate Up
CREATE TABLE orders (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id INT  NOT NULL,
    invoice_no VARCHAR(255),
    grand_total FLOAT DEFAULT 0.0 NOT NULL,
    status ENUM('waiting_for_payment', 'delivery', 'complete', 'cancel'),
    shipping_address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);


-- +migrate Down
DROP TABLE orders;