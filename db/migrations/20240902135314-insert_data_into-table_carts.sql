
-- +migrate Up
INSERT INTO carts (user_id, product_id, created_at, updated_at)
VALUES 
    (1, 101, '2024-09-01 10:00:00', '2024-09-01 10:00:00'),
    (2, 102, '2024-09-01 11:00:00', '2024-09-01 11:00:00'),
    (3, 103, '2024-09-01 12:00:00', '2024-09-01 12:00:00'),
    (4, 104, '2024-09-01 13:00:00', '2024-09-01 13:00:00'),
    (5, 105, '2024-09-01 14:00:00', '2024-09-01 14:00:00');


-- +migrate Down
DELETE FROM carts
WHERE (user_id, product_id) IN 
    ((1, 101), 
     (2, 102), 
     (3, 103), 
     (4, 104), 
     (5, 105));