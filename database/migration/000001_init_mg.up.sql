CREATE TABLE url_shortener 
(
    id BIGINT PRIMARY KEY,
    long_url VARCHAR(255) NOT NULL,
    short_code VARCHAR(20) UNIQUE NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    expiration_date TIMESTAMP
);