CREATE TABLE email_users (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       user_name VARCHAR(255) NOT NULL,
       password_hash VARCHAR(255) NOT NULL,
       writer VARCHAR(100) NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS email_users_on_user_name_writer ON email_users (user_name, writer);