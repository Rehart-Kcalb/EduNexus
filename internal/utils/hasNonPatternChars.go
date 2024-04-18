package utils

import "regexp"

func HasNonPatternChars(str, pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(str)
}
