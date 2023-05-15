DROP INDEX IF EXISTS idx_tweet_tags_writer;
ALTER TABLE tweet_tags DROP COLUMN writer;
DROP INDEX IF EXISTS idx_tags_writer;
ALTER TABLE tags DROP COLUMN writer;
DROP INDEX IF EXISTS idx_tweet_writer;
ALTER TABLE tweets DROP COLUMN writer;
