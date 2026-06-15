package main

func IsUpper(s string) bool {
	isUpperCheck := false
	str := []rune(s)
	for _, runes := range str {
		if runes >= 'A' && runes <= 'Z' {
			isUpperCheck = true
		} else {
			return false
		}
	}
	return isUpperCheck
}
