package main

func IsLower(s string) bool {
	isLowerCheck := false
	str := []rune(s)
	for _, runes := range str {
		if runes >= 'a' && runes <= 'z' {
			isLowerCheck = true
		} else {
			return false
		}
	}
	return isLowerCheck
}
