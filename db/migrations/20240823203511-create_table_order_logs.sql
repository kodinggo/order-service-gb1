
-- +migrate Up
CREATE TABLE order_logs (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    order_id INT NOT NULL,
    status ENUM('waiting_for_payment', 'delivery', 'complete', 'cancel'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);


-- +migrate Down
DROP TABLE order_logs;