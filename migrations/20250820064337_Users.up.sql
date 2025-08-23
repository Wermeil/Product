CREATE TABLE IF NOT EXISTS users(
                                    ID SERIAL PRIMARY KEY,
                                    Email VARCHAR(50) NOT NULL,
    Password VARCHAR NOT NULL,
    Deleted_at TIMESTAMP NOT NULL DEFAULT NOW(),
    Created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    Updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    )