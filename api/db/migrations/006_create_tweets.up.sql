CREATE TABLE tweets (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       user_id uuid NOT NULL,
       tweet VARCHAR(300) NOT NULL,
       flags VARCHAR(100) NOT NULL DEFAULT '',
       created_at TIMESTAMPTZ NOT NULL DEFAULT (now() at time zone 'utc')
);
CREATE INDEX IF NOT EXISTS idx_tweets_user_timestamp ON tweets(user_id, created_at);
CREATE INDEX IF NOT EXISTS idx_tweets_timestamp ON tweets(created_at);

CREATE TABLE tags (
       id BIGSERIAL PRIMARY KEY,
       user_id uuid NOT NULL,
       tag VARCHAR(50) NOT NULL,
       created_at TIMESTAMPTZ NOT NULL DEFAULT (now() at time zone 'utc'),
       CONSTRAINT uniq_user_id_tag UNIQUE(user_id, tag)
);

CREATE INDEX IF NOT EXISTS idx_index_user_id_tag ON tags(user_id, tag);
CREATE INDEX IF NOT EXISTS idx_tags_tag ON tags (tag);

CREATE TABLE tweet_tags (
       id BIGSERIAL PRIMARY KEY,
       tag VARCHAR(50) NOT NULL,
       tweet_id uuid NOT NULL,
       user_id uuid NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_tweet_tags_user_id_tag ON tweet_tags(user_id, tag);
