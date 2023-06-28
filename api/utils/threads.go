package utils

import (
	"regexp"
	"strconv"
	"strings"
)

type ThreadDetails struct {
	ID  string
	Seq int
}

// ExtractThreads
// The format is #thread:<uuid of thread>:<0 indexed sequence number>
func ExtractThreads(tweet string) ([]*ThreadDetails, error) {
	p, err := regexp.Compile("[#]thread:([a-z0-9-]{36}):([0-9]{1,})")
	if err != nil {
		return nil, err
	}

	// extract the first line
	lines := strings.Split(tweet, "\n")
	matches := p.FindAllStringSubmatch(lines[0], -1)

	ret := []*ThreadDetails{}
	for _, match := range matches {
		seq, err := strconv.Atoi(match[2])
		if err != nil {
			// ignore if the seq is not a number
			continue
		}
		v := &ThreadDetails{
			ID:  match[1],
			Seq: seq,
		}
		ret = append(ret, v)
	}
	return ret, nil
}
