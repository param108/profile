package models

import "time"

type Thread struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Deleted   bool      `json:"deleted"`
	Writer    string    `json:"writer"`
}

type ThreadTweet struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	TweetID   string    `json:"tweet_id"`
	ThreadID  string    `json:"thread_id"`
	CreatedAt time.Time `json:"created_at"`
	Deleted   bool      `json:"deleted"`
	Writer    string    `json:"writer"`
}

type ThreadData struct {
	Thread
	Tweets	[]Tweet `json:"tweets"`
}

type DeleteThreadRequest struct {
	ThreadID string `json:"thread_id"`
}
