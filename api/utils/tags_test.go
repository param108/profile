package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTweet struct {
	tweet string
	tags  []string
}

func TestTagExtraction(t *testing.T) {
	tweets := []testTweet{
		// 0: simple tag in middle of sentence
		{
			`
There was #no way to do this`,
			[]string{"no"},
		},
		// 1: tag at beginning of sentence
		{
			`
#the end is nigh`,
			[]string{"the"},
		},
		// 2: tag at end of sentence
		{
			`
Towards the end it was #horrible
You know I almost #died`,
			[]string{"horrible", "died"},
		},
		// 3: tag with _
		{
			`
Towards the end it was #horrible_
You know I almost #totally_died`,
			[]string{"horrible_", "totally_died"},
		},
		// 4: -, or any other character terminates
		{
			`
Towards the end it was #horrible-
You know I almost #totally-died`,
			[]string{"horrible", "totally"},
		},
		// 5: multiple tags in the same line
		// ignore tags on first line
		{
			`#top #tags #for #config #only
This tweet has #two tags in the same #line.`,
			[]string{"two", "line"},
		},
		// 6: multiple consecutive # or a # on its own
		// should be ignored
		{
			`#top #tags #for #config #only
# This tweet has #two tags in the same #line.

## Subheading`,
			[]string{"two", "line"},
		},
	}

	for tc, tweet := range tweets {
		tags, err := extractTags(tweet.tweet)
		assert.Nil(t, err, "%d: failed to extract tags", tc)

		assert.Equal(t, len(tweet.tags), len(tags), "%d: incorrect number of tags found", tc)

		for idx, tag := range tags {
			assert.Equal(t, tweet.tags[idx], tag, "%d: invalid tag found", tc)
		}

	}
}
