CREATE TABLE users (
    u_id INT PRIMARY KEY,
    u_email VARCHAR(255) UNIQUE NOT NULL,
    u_phone_number VARCHAR(15),
    u_first_name VARCHAR(255) NOT NULL,
    u_last_name VARCHAR(255) NOT NULL,
    u_nick_name VARCHAR(255), -- nullable
    u_password VARCHAR(255) NOT NULL
);

-- New table for available rental items
CREATE TABLE rentals (
    rental_id SERIAL PRIMARY KEY,
    i_id INT NOT NULL,
    renter_id INT NOT NULL,
    owner_id INT NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'approved', 'rejected', 'completed', 'cancelled')),
    total_price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (i_id) REFERENCES items(i_id),
    FOREIGN KEY (renter_id) REFERENCES users(u_id),
    FOREIGN KEY (owner_id) REFERENCES users(u_id)
);

CREATE TABLE addresses (
    a_id INT PRIMARY KEY,
    u_id INT NOT NULL,
    a_street VARCHAR(255) NOT NULL,
    a_city VARCHAR(100) NOT NULL,
    a_state VARCHAR(100) NOT NULL,
    a_zipcode VARCHAR(20) NOT NULL,
    a_country VARCHAR(100) NOT NULL,
    FOREIGN KEY (u_id) REFERENCES users(u_id)
);

CREATE TABLE categories (
    c_id INT PRIMARY KEY,
    c_name VARCHAR(255) NOT NULL,
    c_description TEXT NOT NULL
);

CREATE TABLE items (
    i_id INT PRIMARY KEY,
    i_name VARCHAR(255) NOT NULL,
    i_description TEXT NOT NULL,
    i_image BYTEA, -- nullable
    c_id INT NOT NULL,
    i_price INT NOT NULL,
    i_date_listed TIMESTAMP NOT NULL,
    i_quantity INT NOT NULL,
    i_available BOOLEAN NOT NULL,
    FOREIGN KEY (c_id) REFERENCES categories(c_id)
);

CREATE TYPE transaction_type AS ENUM ('Purchase', 'Sale', 'Refund', 'Rental');

CREATE TABLE transactions (
    t_id INT PRIMARY KEY,
    u_id INT NOT NULL,
    t_type transaction_type NOT NULL,
    i_id INT NOT NULL,
    t_date TIMESTAMP NOT NULL,
    t_amount INT NOT NULL,
    FOREIGN KEY (u_id) REFERENCES users(u_id),
    FOREIGN KEY (i_id) REFERENCES items(i_id)
);

CREATE TABLE reviews (
    r_id INT PRIMARY KEY,
    r_comment TEXT NOT NULL,
    r_star INT NOT NULL CHECK (r_star BETWEEN 1 AND 5),
    u_id INT NOT NULL,
    FOREIGN KEY (u_id) REFERENCES users(u_id)
);