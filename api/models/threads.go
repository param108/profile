package models

import "time"

type Thread struct {
	ID        string    `gorm:"default:uuid_generate_v4()" json:"id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Deleted   bool      `json:"deleted"`
	Writer    string    `json:"writer"`
	Name      string    `json:"name"`
}

type ThreadTweet struct {
	ID        string    `gorm:"default:uuid_generate_v4()" json:"id"`
	UserID    string    `json:"user_id"`
	TweetID   string    `json:"tweet_id"`
	ThreadID  string    `json:"thread_id"`
	CreatedAt time.Time `json:"created_at"`
	Deleted   bool      `json:"deleted"`
	Writer    string    `json:"writer"`
}

type ThreadData struct {
	Thread
	Tweets []*Tweet `json:"tweets"`
}

type ThreadRawData struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	Name           string    `json:"name"`
	Deleted        bool      `json:"deleted"`
	Writer         string    `json:"writer"`
	TweetID        string    `json:"tweet_id"`
	TweetUserID    string    `json:"tweet_user_id"`
	TweetTweet     string    `json:"tweet_tweet"`
	TweetFlags     string    `json:"tweet_flags"`
	TweetWriter    string    `json:"tweet_writer"`
	TweetCreatedAt time.Time `json:"tweet_created_at"`
}

type DeleteThreadRequest struct {
	ThreadID string `json:"thread_id"`
}

type CreateThreadRequest struct {
	Name string `json:"name"`
}
