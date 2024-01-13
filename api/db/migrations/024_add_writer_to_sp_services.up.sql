ALTER TABLE sp_services
ADD COLUMN writer uuid NOT NULL DEFAULT uuid_generate_v4();
