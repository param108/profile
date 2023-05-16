package utils

import (
	"regexp"
	"strings"
)

func extractTags(tweet string) ([]string, error) {
	p, err := regexp.Compile("[#]([a-zA-Z0-9_]{1,})")
	if err != nil {
		return nil, err
	}

	ret := []string{}

	lines := strings.Split(tweet, "\n")
	for _, line := range lines[1:] {
		val := p.FindAllStringSubmatch(line, -1)
		for _, v := range val {
			ret = append(ret, v[1])
		}
	}

	return ret, nil
}
