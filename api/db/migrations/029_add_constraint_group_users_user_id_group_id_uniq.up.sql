ALTER TABLE sp_group_users
ADD CONSTRAINT sp_group_users_group_id_user_id_uniq UNIQUE(sp_group_id, sp_user_id);
