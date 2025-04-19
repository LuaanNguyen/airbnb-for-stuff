------------ User Queries ------------
-- Get all users
SELECT u_id, u_email, u_phone_number, u_first_name, u_last_name, u_nick_name 
FROM users;

-- Get user by ID
SELECT u_id, u_email, u_phone_number, u_first_name, u_last_name, u_nick_name 
FROM users 
WHERE u_id = $1;

-- Create new user
INSERT INTO users (u_id, u_email, u_phone_number, u_first_name, u_last_name, u_nick_name, u_password)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING u_id;

-- Update user
UPDATE users 
SET u_email = $1, u_phone_number = $2, u_first_name = $3, u_last_name = $4, u_nick_name = $5
WHERE u_id = $6
RETURNING u_id;

-- User authentication
SELECT u_id, u_password 
FROM users 
WHERE u_email = $1;

------------ Address Queries ------------
-- Get user addresses
SELECT a_id, a_street, a_city, a_state, a_zipcode, a_country 
FROM addresses 
WHERE u_id = $1;

-- Create address
INSERT INTO addresses (a_id, u_id, a_street, a_city, a_state, a_zipcode, a_country)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING a_id;

-- Update address
UPDATE addresses 
SET a_street = $1, a_city = $2, a_state = $3, a_zipcode = $4, a_country = $5
WHERE a_id = $6
RETURNING a_id;

-- Delete address
DELETE FROM addresses 
WHERE a_id = $1;

------------Item Queries ------------
-- Get all items
SELECT i_id, i_name, i_description, i_image, c_id, i_price, i_date_listed, i_quantity, i_available 
FROM items;

-- Get item by ID
SELECT i_id, i_name, i_description, i_image, c_id, i_price, i_date_listed, i_quantity, i_available 
FROM items 
WHERE i_id = $1;

-- Create item
INSERT INTO items (i_id, i_name, i_description, i_image, c_id, i_price, i_date_listed, i_quantity, i_available)
VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, $7, $8)
RETURNING i_id;

-- Update item
UPDATE items 
SET i_name = $1, i_description = $2, i_image = $3, c_id = $4, i_price = $5, i_quantity = $6, i_available = $7
WHERE i_id = $8
RETURNING i_id;

-- Delete item
DELETE FROM items 
WHERE i_id = $1;

-- Search items
SELECT i_id, i_name, i_description, i_image, c_id, i_price, i_date_listed, i_quantity, i_available 
FROM items 
WHERE 
    i_name ILIKE $1 OR 
    i_description ILIKE $1
ORDER BY i_date_listed DESC;

-- Get available items
SELECT i_id, i_name, i_description, i_image, c_id, i_price, i_date_listed, i_quantity, i_available 
FROM items 
WHERE i_available = true AND i_quantity > 0;

------------ Category Queries ------------
-- Get all categories
SELECT c_id, c_name, c_description 
FROM categories;

-- Get items by category
SELECT i_id, i_name, i_description, i_image, c_id, i_price, i_date_listed, i_quantity, i_available 
FROM items 
WHERE c_id = $1;

------------ Transaction Queries ------------
-- Create transaction
INSERT INTO transactions (t_id, u_id, t_type, i_id, t_date, t_amount)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, $5)
RETURNING t_id;

-- Get user transactions
SELECT t_id, t_type, i_id, t_date, t_amount 
FROM transactions 
WHERE u_id = $1
ORDER BY t_date DESC;

------------ Get transaction by ID ------------
SELECT t_id, u_id, t_type, i_id, t_date, t_amount 
FROM transactions 
WHERE t_id = $1;

-- Update transaction
UPDATE transactions 
SET t_type = $1, t_amount = $2
WHERE t_id = $3
RETURNING t_id;

------------ Review Queries ------------
-- Create review
INSERT INTO reviews (r_id, r_comment, r_star, u_id)
VALUES ($1, $2, $3, $4)
RETURNING r_id;

-- Get item reviews
SELECT r_id, r_comment, r_star, u_id 
FROM reviews 
WHERE i_id = $1
ORDER BY r_id DESC;

-- Get user reviews
SELECT r_id, r_comment, r_star, i_id 
FROM reviews 
WHERE u_id = $1
ORDER BY r_id DESC;

-- Update review
UPDATE reviews 
SET r_comment = $1, r_star = $2
WHERE r_id = $3
RETURNING r_id;

-- Delete review
DELETE FROM reviews 
WHERE r_id = $1; 