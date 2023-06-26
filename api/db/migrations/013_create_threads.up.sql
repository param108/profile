CREATE TABLE threads (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       user_id uuid NOT NULL,
       created_at timestamp with time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
       deleted boolean DEFAULT false NOT NULL,
       writer uuid NOT NULL
);

CREATE INDEX idx_threads_user_id ON threads(user_id, writer);

CREATE TABLE thread_tweets (
       id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
       user_id uuid NOT NULL,
       tweet_id uuid NOT NULL,
       thread_id uuid NOT NULL,
       created_at timestamp with time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
       deleted boolean DEFAULT false NOT NULL,
       writer uuid NOT NULL
);

CREATE INDEX idx_thread_tweets_user_id ON thread_tweets(user_id, thread_id, writer);
