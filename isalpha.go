package main

func IsAlpha(s string) bool {
	alphaCheck := false
	str := []rune(s)
	for _, runes := range str {
		if (runes >= 'a' && runes <= 'z') || (runes >= 'A' && runes <= 'Z') || (runes >= '0' && runes <= '9') {
			alphaCheck = true
		} else {
			return false
		}
	}
	return alphaCheck
}
