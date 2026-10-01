CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,

    account_id INT NOT NULL,

    transaction_type VARCHAR(50) NOT NULL,

    amount DECIMAL(15, 2) NOT NULL,

    currency VARCHAR(10) NOT NULL DEFAULT 'GBP',

    description VARCHAR(255),

    reference VARCHAR(100) NOT NULL UNIQUE,

    status VARCHAR(30) NOT NULL DEFAULT 'completed',

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (account_id)
        REFERENCES accounts(id)
        ON DELETE CASCADE
);