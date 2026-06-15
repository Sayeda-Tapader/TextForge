package main

func ToUpper(s string) string {
	newStr := ""
	for _, runes := range s {
		if runes >= 'a' && runes <= 'z' {
			//do something to convert
			/* castAscii := int(runes)
			capitaliseAscii := castAscii-32 */
			newStr += string(runes - 32)
		} else {
			newStr += string(runes)
		}
	}
	return newStr
}
