CREATE TABLE IF NOT EXISTS conf
(
    id         BIGSERIAL PRIMARY KEY,
    meta       VARCHAR,
    text      VARCHAR,
    user_id    BIGINT REFERENCES users (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
