package testing

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// START OMIT
// name = TestXXX, prefix = Test
func isTest(name, prefix string) bool {
	if !strings.HasPrefix(name, prefix) {
		return false
	}
	if len(name) == len(prefix) { // "Test" is ok
		return true
	}
	rune, _ := utf8.DecodeRuneInString(name[len(prefix):])
	return !unicode.IsLower(rune)
}

// END OMIT
