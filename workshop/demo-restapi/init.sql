CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    price NUMERIC(10, 2) NOT NULL CHECK (price > 0),
    description TEXT NOT NULL
);

-- Seed test data
INSERT INTO products (name, price, description) VALUES
    ('Laptop', 999.99, 'High-performance laptop'),
    ('Mouse', 29.99, 'Wireless ergonomic mouse'),
    ('Keyboard', 79.99, 'Mechanical keyboard with RGB')
ON CONFLICT (name) DO NOTHING;
