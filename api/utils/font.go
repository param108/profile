package utils

import (
	"regexp"
	"strings"
)

type FontDetails struct {
	Name string
}

// ExtractThreads
// The format is #font:<name of font>
func ExtractFonts(tweet string) ([]*FontDetails, error) {
	p, err := regexp.Compile("[#]font:([a-zA-Z0-9-]{1,})")
	if err != nil {
		return nil, err
	}

	// extract the first line
	lines := strings.Split(tweet, "\n")
	matches := p.FindAllStringSubmatch(lines[0], -1)

	ret := []*FontDetails{}
	for _, match := range matches {
		v := &FontDetails{
			Name: match[1],
		}
		ret = append(ret, v)
	}
	return ret, nil
}
