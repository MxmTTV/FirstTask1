CREATE TABLE IF NOT EXISTS users (
         id SERIAL PRIMARY KEY,
         email VARCHAR(255) NOT NULL,
        password_hash VARCHAR(72) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL
    );

