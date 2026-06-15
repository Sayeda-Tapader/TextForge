package main

// func Capitalize(s string) string {
// 	// so my code needs to differentiate each word
// 	//what defines the word?
// 	// we create a word from each point where there is the start of an alpabetical letter && (the start of a non-alphabetical letter)-1

// 	//order of things
// 	//read through the string

// 	sRead := s
// 	wStart := 0
// 	wEnd := 0
// 	for i, runes := range sRead{
// 		//now that i go through, must now identify each word range
// 		if (runes >= 'A' && runes <= 'Z') || (runes >= 'a' && runes <= 'z') {
// 			word = sRead[i]
// 			break
// 		}

// 		if !(runes)
// 	}
// }

func Capitalize(s string) string {
	r := []rune(s)
	newWord := true

	for i := 0; i < len(r); i++ {
		c := r[i]

		isLetter := (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
		isDigit := (c >= '0' && c <= '9')

		if isLetter || isDigit {
			if newWord {

				if c >= 'a' && c <= 'z' {
					r[i] = c - 32
				}
				newWord = false
			} else {
				if c >= 'A' && c <= 'Z' {
					r[i] = c + 32
				}
			}
		} else {
			newWord = true
		}
	}
	return string(r)
}
