CREATE TABLE IF NOT EXISTS bin
(
    id         BIGSERIAL PRIMARY KEY,
    meta       VARCHAR,
    file_name  VARCHAR,
    file_size  BIGINT,
    user_id    BIGINT REFERENCES users (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
