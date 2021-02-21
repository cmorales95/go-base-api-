CREATE TABLE IF NOT EXISTS users (
    id          UUID PRIMARY KEY,
    name        VARCHAR NOT NULL,
    last_name   VARCHAR NOT NULL,
    email       VARCHAR NOT NULL,
    password    VARCHAR NOT NULL,
    birth       TIMESTAMP(0), 
    is_active   BOOLEAN DEFAULT TRUE,
    created_at  TIMESTAMP(0) DEFAULT NOW(),
    updated_at  TIMESTAMP(0) DEFAULT NOW(),
    deleted_at  TIMESTAMP(0),
    google      BOOLEAN DEFAULT FALSE,
    facebook    BOOLEAN DEFAULT FALSE,
    twitter     BOOLEAN DEFAULT FALSE
);
