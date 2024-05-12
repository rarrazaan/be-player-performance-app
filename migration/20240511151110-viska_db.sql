
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id uuid DEFAULT uuid_generate_v4(),
    username varchar NOT NULL,
    email varchar(320) NOT NULL UNIQUE,
    password varchar,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_details(
    id uuid DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    full_name varchar NOT NULL,
    age int NOT NULL,
    gender varchar NOT NULL,
    address varchar NOT NULL,
    phone_number varchar NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

INSERT INTO users (id,username,email,"password",created_at,updated_at) VALUES
    ('b83d8d92-26b5-4821-a6e7-807118dad155','rafiif','karrazaan@gmail.com','$2a$10$A0PG0QQ2q1YFFEGWT2QFkOhdmfoi7zkIu05i5BSaImmv86Ry22Kwa','2024-02-17 15:04:48.101016+07',NULL),
    ('f44e7c28-0154-4c9a-81b0-7735455057e5','TEST','test@example.com','$2a$10$A0PG0QQ2q1YFFEGWT2QFkOhdmfoi7zkIu05i5BSaImmv86Ry22Kwa','2024-02-17 15:04:48.101016+07',NULL);

INSERT INTO user_details (user_id, full_name, age, gender, address, phone_number) VALUES
    ('b83d8d92-26b5-4821-a6e7-807118dad155', 'Imam Rafiif Arrazaan', 23, 'MALE', 'Jln. Kayutangan No.49 RT.03/RW.06 Pengkol, Jepara', '082245155712'),
    ('f44e7c28-0154-4c9a-81b0-7735455057e5', 'TEST NAME', 20, 'MALE', 'TEST ADDRESS', '081180081180');

-- +migrate Down
-- DROP TABLE IF EXISTS users;