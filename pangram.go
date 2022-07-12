package pangram

import "strings"

func IsPangram(input string) bool {
	low := strings.ToLower(input)
	for _, v := range "abcdefghijklmnopqrstuvwxyz" {
		if !strings.ContainsRune(low, v) {
			return false
		}
	}
	return true
}
