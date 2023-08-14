package models

import "time"

type Tweet struct {
	ID        string    `gorm:"default:uuid_generate_v4()" json:"id"`
	UserID    string    `json:"user_id"`
	Tweet     string    `json:"tweet"`
	Flags     string    `json:"flags"`
	Writer    string    `json:"writer"`
	CreatedAt time.Time `json:"created_at"`
}

type Tag struct {
	ID        string    `gorm:"default:uuid_generate_v4()" json:"id"`
	UserID    string    `json:"user_id"`
	Tag       string    `json:"tag"`
	Writer    string    `json:"writer"`
	CreatedAt time.Time `json:"created_at"`
}

type TweetTag struct {
	ID      int    `json:"id"`
	Tag     string `json:"tag"`
	TweetID string `json:"tweet_id"`
	Writer  string `json:"writer"`
	UserID  string `json:"user_id"`
}

type PostTweetsRequest struct {
	Tweet string `json:"tweet"`
	Image string `json:"image"`
}

type PutTweetRequest struct {
	Tweet   string `json:"tweet"`
	TweetID string `json:"tweet_id"`
	Flags   string `json:"flags"`
}

type DeleteTweetRequest struct {
	TweetID string `json:"tweet_id"`
}
