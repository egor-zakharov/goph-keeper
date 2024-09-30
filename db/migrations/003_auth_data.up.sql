CREATE TABLE IF NOT EXISTS auth
(
    id         BIGSERIAL PRIMARY KEY,
    meta       VARCHAR,
    login      VARCHAR,
    password   VARCHAR,
    user_id    BIGINT REFERENCES users (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
