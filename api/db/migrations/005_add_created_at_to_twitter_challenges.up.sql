ALTER TABLE twitter_challenges
      ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
ALTER TABLE twitter_challenges
      ADD COLUMN writer uuid NOT NULL;

CREATE INDEX idx_twitter_challenges_created_at on twitter_challenges(created_at);
