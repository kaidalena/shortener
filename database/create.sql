-- CREATE DATABASE test_ozon;

CREATE TABLE IF NOT EXISTS addresses(
    id SERIAL,
    long_url text UNIQUE,
    short_url text UNIQUE
);
