-- +migrate Up
CREATE TABLE dedications (
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(100),
    message    TEXT NOT NULL,
    song_url   TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);