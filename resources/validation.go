package resources

import (
	"errors"
	"regexp"
	"strings"
)

var timeTest = regexp.MustCompile(`^[\d]{2}min$`)

func validateTime(time string) (string, error) {
	if strings.ToLower(time) == "build up" {
		return "Build Up", nil
	}

	if strings.ToLower(time) == "post match" {
		return "Post Match", nil
	}

	if timeTest.MatchString(time) {
		return time, nil
	}

	return "", errors.New(`Please specify one of "Build Up", "Post Match", or the format XXmin (i.e. 86min)`)
}
