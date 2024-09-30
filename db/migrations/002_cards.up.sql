CREATE TABLE IF NOT EXISTS cards
(
    id              BIGSERIAL PRIMARY KEY,
    number          VARCHAR NOT NULL,
    expiration_date VARCHAR NOT NULL,
    holder_name     VARCHAR NOT NULL,
    cvv             VARCHAR NOT NULL,
    user_id         BIGINT REFERENCES users (id),
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP
);

CREATE UNIQUE INDEX unique_number_user_id ON cards (number, user_id);