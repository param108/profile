CREATE TABLE sp_users (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       phone   VARCHAR(20) NOT NULL UNIQUE,
       name    VARCHAR(100) NOT NULL,
       photo_url TEXT NOT NULL,
       created_at TIMESTAMPTZ default NOW(),
       updated_at TIMESTAMPTZ default NOW(),
       deleted BOOLEAN NOT NULL DEFAULT FALSE,
       writer uuid NOT NULL,
       deleted_at TIMESTAMPTZ,
       profile_updated BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE sp_groups (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       name    VARCHAR(100) NOT NULL,
       parent uuid,
       deleted BOOLEAN NOT NULL DEFAULT FALSE,
       description TEXT NOT NULL DEFAULT '',
       created_at TIMESTAMPTZ default NOW(),
       updated_at TIMESTAMPTZ default NOW(),
       deleted_at TIMESTAMPTZ,
       writer uuid NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_sp_groups_name
ON sp_groups(name, writer);

CREATE TABLE sp_group_users (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       sp_group_id uuid NOT NULL,
       sp_user_id uuid NOT NULL,
       deleted BOOLEAN NOT NULL DEFAULT FALSE,
       created_at TIMESTAMPTZ default NOW(),
       deleted_at TIMESTAMPTZ,
       writer uuid NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_sp_group_users_sp_user_id
ON sp_group_users(sp_user_id, writer);

CREATE INDEX IF NOT EXISTS idx_sp_group_users_sp_group_id
ON sp_group_users(sp_group_id, writer);

CREATE TABLE sp_messages (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       sp_user_id uuid NOT NULL,
       msg_type varchar(100) NOT NULL,
       msg_value int NOT NULL DEFAULT 0,
       msg_text TEXT NOT NULL DEFAULT '',
       created_at TIMESTAMPTZ default NOW(),
       writer uuid NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_sp_messages_msg_type
ON sp_messages(msg_type, created_at, writer);

CREATE INDEX IF NOT EXISTS idx_sp_messages_user_id
ON sp_messages(sp_user_id, created_at, writer);

CREATE TABLE sp_message_comments (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       sp_user_id uuid NOT NULL,
       sp_message_id uuid NOT NULL,
       msg_text TEXT NOT NULL DEFAULT '',
       created_at TIMESTAMPTZ default NOW(),
       writer uuid NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_sp_message_comments_sp_message_id
ON sp_message_comments(sp_message_id, writer);

CREATE TABLE sp_group_messages (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       sp_group_id uuid NOT NULL,
       sp_user_id uuid NOT NULL,
       sp_message_id uuid NOT NULL,
       created_at TIMESTAMPTZ default NOW(),
       writer uuid NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_sp_group_messages_sp_group_id
ON sp_group_messages(sp_group_id, created_at, writer);
