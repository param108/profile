DROP INDEX IF EXISTS idx_image_compressed_failed;

ALTER TABLE tweets
DROP COLUMN image_compressed_failed;

ALTER TABLE tweets
DROP COLUMN image_compressed;

ALTER TABLE tweets
DROP COLUMN image;
