package store

import (
	"log"
	"time"

	"github.com/param108/profile/api/models"
)

// Store This interface encapsulates all storage needs.
// The idea here is that it doesn't matter the implementation of the store.
// What is important is the Data flow.
type Store interface {
	CreateUser(handle string, profile string, role models.UserRole, writer string) (*models.User, error)
	FindOrCreateUser(handle string, profile string, role models.UserRole, writer string) (*models.User, error)
	GetUser(userID string, writer string) (*models.User, error)
	GetUserByHandle(username string, writer string) (*models.User, error)
	CreateInvalidToken(token string, writer string) error
	IsInvalidToken(token string, writer string) (bool, error)

	// CreateTwitterChallenge Takes a token and returns a uuid
	// which can be passed to GetTwitterChallenge to retrieve it
	CreateTwitterChallenge(token, redirectURL, writer string) (string, error)

	// GetTwitterChallenge Takes a uuid previously returned by CreateTwitterChallenge
	// and returns the saved challenge string
	GetTwitterChallenge(key, writer string) (string, string, error)

	// DeleteOldTwitterChallenges Deletes TwitterChallenges > 24 hours old
	DeleteOldTwitterChallenges(d time.Duration)

	// InsertTweet inserts a tweet and all tags in a transaction
	// It will ignore the first line of flags as those are for display only.
	// From the rest extract tags and insert them in tweet table
	InsertTweet(userID string, tweet string,
		flags string, writer string) (*models.Tweet, []*models.Tag, error)

	// UpdateTweet edits a tweet
	UpdateTweet(userID, tweetID, tweet, flags,
		writer string) (*models.Tweet, []*models.Tag, error)

	// GetTweets get tweets between a offset and limit
	GetTweets(userID string, offset, limit int, reverse bool, writer string) ([]*models.Tweet, error)

	// GetTweets get tweets between a offset and limit
	GetTweet(userID string, tweetID string, writer string) (*models.Tweet, error)

	// DeleteTweet deletes a tweet
	DeleteTweet(userID string, tweetID string, writer string) (*models.Tweet, error)

	// GetTags return all tags for a user
	// writer is optional. Empty value is all writers
	GetTags(userID string, writer string) ([]*models.Tag, error)

	// GetTweetTags return all the TweetTags for a tweet
	GetTweetTags(userID, tweetID, writer string) ([]*models.Tag, error)

	// SearchTweetsByTags return all tweets for user
	// by tag. Return in Chronologically descending order.
	// writer is optional. Empty value is all writers.
	SearchTweetsByTags(userID string, tags []string, offset, limit int, reverse bool, writer string) ([]*models.Tweet, error)

	// OneTime APIs
	SetOneTime(val, writer string) (*models.Onetime, error)

	// GetOneTime Returns a onetime record by id and writer if it
	// is not older than expiry
	GetOneTime(id string, expiry time.Duration, writer string) (*models.Onetime, error)

	// DeleteOldOneTimes Delete all one time entries older than one hour
	DeleteOldOneTimes(expiry time.Duration, writer string) error

	// DeleteGuestData If number of tweets > 40 delete older tweets
	DeleteGuestData(userID string, maxTweets int, writer string) error

	// CreateThread Create a thread for a user
	CreateThread(userID, name, writer string) (*models.Thread, error)

	// DeleteThread Create a thread for a user
	DeleteThread(userID string, threadID string, writer string) (*models.Thread, error)

	// AddTweetToThread Add tweet to a thread
	AddTweetToThread(userID, tweetID, threadID, writer string) error

	// DelTweetFromThread Del tweet from a thread
	DelTweetFromThread(userID, tweetID, threadID, writer string) error

	// GetThread Get thread details and all the tweets attached to it
	GetThread(userID, threadID, writer string) (*models.ThreadData, error)

	//All functions after this line are for internal use only.
	// UnsafeGetAllTweets Get all Tweets paginated 20 at a time for a writer,
	// regardless of user
	UnsafeGetAllTweets(writer string, offset int, count int) ([]*models.Tweet, int, error)

	// UnsafeDelete Delete all tweets from a table for a writer
	UnsafeDelete(table string, writer string) error

	// CreateOTP Create a new OTP entry
	CreateOTP(phone string, now time.Time, writer string) error

	// GetOTP Get an OTP entry without checking
	GetOTP(phone, writer string) (*models.SpOtp, error)

	// CheckOTP Get an OTP after verification.
	// Returns error if invalid or expiry etc
	CheckOTP(phone, code string, now time.Time, writer string) (*models.SpOtp, error)

	// ExpireOTPs expire old otps
	ExpireOTPs(now time.Time, writer string) error

	// DeleteAllOTPs Delete all otps of a writer.
	DeleteAllOTPs(writer string) error

	// FindOrCreateSPUser Find or create an SP User
	FindOrCreateSPUser(phone string, writer string) (*models.SpUser, error)
}

func Periodic(s Store, writer string) {
	if err := s.DeleteOldOneTimes(time.Hour, writer); err != nil {
		log.Printf("Failed delete old one times: %s", err.Error())
	}
	if err := s.DeleteGuestData(models.GuestUserID, 40, writer); err != nil {
		log.Printf("Failed delete Guest tweets: %s", err.Error())
	}
	if err := s.ExpireOTPs(time.Now(), writer); err != nil {
		log.Printf("Failed expire OTPs: %s", err.Error())
	}
}
