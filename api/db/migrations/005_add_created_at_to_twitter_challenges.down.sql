DROP INDEX idx_twitter_challenges_created_at;
ALTER TABLE twitter_challenges
      DROP COLUMN created_at;
ALTER TABLE twitter_challenges
      DROP COLUMN writer;
