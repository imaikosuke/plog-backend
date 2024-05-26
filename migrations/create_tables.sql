CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE photologs (
    id SERIAL NOT NULL,
    user_id INTEGER NOT NULL,
    generated_text TEXT NOT NULL,
    images TEXT[] NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, id)
) PARTITION BY LIST (user_id);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    photolog_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id, photolog_id) REFERENCES photologs (user_id, id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
