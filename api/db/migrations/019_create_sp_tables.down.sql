DROP INDEX IF EXISTS idx_sp_group_messages_sp_group_id;

DROP TABLE IF EXISTS sp_group_messages;

DROP INDEX IF EXISTS idx_sp_message_comments_sp_message_id;

DROP TABLE IF EXISTS sp_message_comments;

DROP INDEX IF EXISTS idx_sp_messages_user_id;

DROP INDEX IF EXISTS idx_sp_messages_msg_type;

DROP TABLE IF EXISTS sp_messages;

DROP INDEX IF EXISTS idx_sp_group_users_sp_user_id;

DROP INDEX IF EXISTS idx_sp_group_users_sp_group_id;

DROP TABLE IF EXISTS sp_group_users;

DROP INDEX IF EXISTS idx_sp_groups_name;

DROP TABLE IF EXISTS sp_groups;

DROP TABLE IF EXISTS sp_users;
