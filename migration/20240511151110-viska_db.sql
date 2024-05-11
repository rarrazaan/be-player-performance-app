
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id uuid DEFAULT uuid_generate_v4(),
    name varchar NOT NULL,
    email varchar(320) NOT NULL UNIQUE,
    password varchar NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT NULL,

    PRIMARY KEY (id)
);

INSERT INTO users (id,name,email,"password",created_at,updated_at) VALUES
    ('f44e7c28-0154-4c9a-81b0-7735455057e5','TEST','test@example.com','$2a$10$A0PG0QQ2q1YFFEGWT2QFkOhdmfoi7zkIu05i5BSaImmv86Ry22Kwa','2024-02-17 15:04:48.101016+07',NULL);


-- +migrate Down
DROP TABLE IF EXISTS users;