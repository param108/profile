ALTER TABLE tweets ADD COLUMN writer uuid NOT NULL;
CREATE INDEX IF NOT EXISTS idx_tweets_writer ON tweets(writer);
ALTER TABLE tags ADD COLUMN writer uuid NOT NULL;
CREATE INDEX IF NOT EXISTS idx_tags_writer ON tags(writer);
ALTER TABLE tweet_tags ADD COLUMN writer uuid NOT NULL;
CREATE INDEX IF NOT EXISTS idx_tweet_tags_writer ON tweet_tags(writer);
