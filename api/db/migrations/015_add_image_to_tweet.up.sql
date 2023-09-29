ALTER TABLE tweets
ADD COLUMN image varchar(150) NOT NULL DEFAULT '';

ALTER TABLE tweets
ADD COLUMN image_compressed BOOLEAN DEFAULT false;

ALTER TABLE tweets
ADD COLUMN image_compressed_failed BOOLEAN DEFAULT false;

CREATE INDEX idx_image_compressed_failed ON tweets (image_compressed, image_compressed_failed);
