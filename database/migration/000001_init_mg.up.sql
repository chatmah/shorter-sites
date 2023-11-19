CREATE TABLE url_shortener (
    id BIGSERIAL PRIMARY KEY,
    long_url VARCHAR(255) NOT NULL,
    short_code VARCHAR(20) UNIQUE NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    expiration_date TIMESTAMP
);

DO $$
BEGIN
    INSERT INTO url_shortener (
        long_url, 
        short_code,
        creation_date,
        expiration_date
    ) 
    VALUES (
        'yongkhengs.com', 
        '000000001', 
        CURRENT_TIMESTAMP, 
        CURRENT_TIMESTAMP + INTERVAL '30 days'
    );
END $$;