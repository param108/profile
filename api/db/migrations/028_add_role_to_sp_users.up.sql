ALTER TABLE sp_users
ADD COLUMN role VARCHAR(8) NOT NULL DEFAULT 'user';
