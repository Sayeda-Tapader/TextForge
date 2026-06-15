package main

func ToLower(s string) string {
	newStr := ""
	for _, runes := range s {
		if runes >= 'A' && runes <= 'Z' {
			//do something to convert
			/* castAscii := int(runes)
			capitaliseAscii := castAscii-32 */
			newStr += string(runes + 32)
		} else {
			newStr += string(runes)
		}
	}
	return newStr
}
