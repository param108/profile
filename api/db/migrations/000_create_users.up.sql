CREATE TYPE user_role AS ENUM ('user', 'admin');

CREATE TABLE users (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       handle TEXT NOT NULL UNIQUE,
       profile TEXT,
       role user_role
);

CREATE INDEX IF NOT EXISTS users_on_handle ON users (handle);
