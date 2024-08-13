package stringx

import (
	"regexp"
	"strings"
)

// ContainNonEnglish 判斷是否包含非英文字母
func ContainNonEnglish(s string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z ]+$", s)
	return !match
}

// IsFirstLetterUp 判斷首字母是否為大寫
func IsFirstLetterUp(s string) bool {
	words := strings.Fields(s)

	for _, w := range words {
		match, _ := regexp.MatchString(`^[A-Z][a-z]*$`, w)
		if !match {
			return false
		}
	}

	return true
}
