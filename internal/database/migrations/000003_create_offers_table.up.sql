CREATE TABLE IF NOT EXISTS offers (
    ID SERIAL PRIMARY KEY,
    Bank VARCHAR(255) NOT NULL,
    Market VARCHAR(255) NOT NULL,
    Category VARCHAR(255),
    Precent float not null,
    Due TIMESTAMP,
    Limitation varchar(255)
);