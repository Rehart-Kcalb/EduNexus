package utils

import "regexp"

// HasNonPatternChars checks if the string contains characters that don't match the given pattern.
func HasNonPatternChars(s, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}
