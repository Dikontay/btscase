CREATE TABLE IF NOT EXISTS cards (
                                     ID SERIAL PRIMARY KEY,
                                     Name VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Surname VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL UNIQUE,
    Address VARCHAR(255),
    Phone VARCHAR(255),
    Role VARCHAR(255)
    );