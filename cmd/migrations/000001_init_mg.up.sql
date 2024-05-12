CREATE TABLE accounts 
(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(120),
    last_name VARCHAR(120),
    email VARCHAR(120) NOT NULL UNIQUE,
    password_hash VARCHAR(120),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories 
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(60)
);

CREATE TABLE products 
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(120) UNIQUE,
    category_id INTEGER NOT NULL,
    description TEXT,
    price DECIMAL(10,2),
    quantity INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category_id FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE orders 
(
    id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL,
    address TEXT NOT NULL,
    total_price DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE TABLE order_item 
(
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER,
    total_price DECIMAL(10,2),
    CONSTRAINT fk_order_id FOREIGN KEY (order_id) REFERENCES orders(id),
    CONSTRAINT fk_product_id FOREIGN KEY (product_id) REFERENCES products(id)
);
