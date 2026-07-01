package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// reads file and returns a byte slice which is stored in 'inputFile'
	inputFile, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("File reading error.", err)
		return
	}

	// creates an editable copy of sample.txt's contents
	strToChange := string(inputFile)

	// holds the changed string to be added to the file result.txt
	result := ""
	// string is converted to byte slice to be able to range through a loop
	// this is to compare the keywords with words in the file
	newStr := strings.Fields(strToChange)
	value := 0 // this variable is used in (cap), (up), (low) to calculate how many words to transform
	temp := ""

	//	using os.file.WriteString method to write to result.txt
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err) 
	}

	defer file.Close()

	// Capitalise
	for i, word := range newStr {
		if word == "(cap)" {
			//	now that we have found the keyword "(cap)", we must look back to the words before it
			//	we do that by looking at i-1, so the word before kw
			//	BUT ALSO! remember we need to remove the key word so lets do a splice :))
			newStr[i-1] = Capitalize(newStr[i-1])
			newStr = append(newStr[:i], newStr[i+1:]...)
			//	so assign the newStr to result.txt - done above
			//	but first convert from type []string to type string
			result = strings.Join(newStr, " ")
		} else if word == "(cap," {
			temp = newStr[i+1]
			a := strings.TrimRight(temp, ")")
			value, _ = strconv.Atoi(a)

			for count := value; count >= 0; count-- {
				newStr[i-count] = Capitalize(newStr[i-count])
			}
			newStr = append(newStr[:i], newStr[i+2:]...)
			result = strings.Join(newStr, " ")
		}
	}

	// Uppercase functionality
	for i, word := range newStr {//TIP ---- MAKE THESE A HELPER FUNCTION MUCH LATER
		if word == "(up)" {
			newStr[i-1] = ToUpper(newStr[i-1])
			newStr = append(newStr[:i], newStr[i+1:]...)
			result = strings.Join(newStr, " ")
		} else if word == "(up," {
			temp = newStr[i+1] // hopefully this is going to be one string and i can split it in the var below
			// a := strings.Split(temp, ")")
			a := strings.TrimRight(temp, ")")
			value, _ = strconv.Atoi(a)

			for count := value; count >= 0; count-- {
				newStr[i-count] = ToUpper(newStr[i-count])
			}
			newStr = append(newStr[:i], newStr[i+2:]...)
			result = strings.Join(newStr, " ")
		}
	}

	// Lowercase functionality
	for i, word := range newStr {
		if word == "(low)" {
			newStr[i-1] = ToLower(newStr[i-1])
			newStr = append(newStr[:i], newStr[i+1:]...)
			result = strings.Join(newStr, " ")
		} else if word == "(low," {
			temp = newStr[i+1]
			a := strings.TrimRight(temp, ")")
			value, _ = strconv.Atoi(a)

			for count := value; count > 0; count-- {
				newStr[i-count] = ToLower(newStr[i-count])
			}
			newStr = append(newStr[:i], newStr[i+2:]...)
			result = strings.Join(newStr, " ")
		}
	}

	// HEX
	for i, word := range newStr {
		if word == "(hex)" {
			//	newStr[i-1] is a string hex value - store it in temp
			temp = newStr[i-1]
			//	https://schadokar.dev/notes/convert-hexadecimal-to-decimal-and-decimal-to-hexadecimal-in-golang/

			num, _ := strconv.ParseUint(temp, 16, 64)	//	conversion of hex to decimal
			// if err != nil {
			// 	fmt.Println(err)
			// }
			
			newStr[i-1] = strconv.FormatUint(num, 10)	// conversion of uint64 to string 
			newStr = append(newStr[:i], newStr[i+1:]...)
			result = strings.Join(newStr, " ")
		}
	}

	// BIN
	for i, word := range newStr {
		if word == "(bin)" {
			//	newStr[i-1] is a string binary value - store it in temp
			temp = newStr[i-1]
			
			num, _ := strconv.ParseUint(temp, 2, 64)	//	conversion of binary to decimal
			// if err != nil {
			// 	fmt.Println(err)
			// }
			
			newStr[i-1] = strconv.FormatUint(num, 10)	// conversion of uint64 to string 
			newStr = append(newStr[:i], newStr[i+1:]...)
			result = strings.Join(newStr, " ")
		}
	}

	//	PUNCTUATION
	for i, word := range newStr {// lets look through the whole list of punctuation
		//	but now we should look at every letter in every word
		//	so var 'i' is the position of which word i am in
		//	while var char is the letter in that word
		for _, char := range word {
			//	issue with double punctuation, FIX LATER
			switch char {
			case '.':
				newStr[i-1] = newStr[i-1] + "."
				newStr[i] = strings.TrimLeft(word, ".")

				if len(newStr[i]) == 0 {
					newStr = append(newStr[:i], newStr[i+1:]...)
				}

			case ',':
				newStr[i-1] = newStr[i-1] + ","
				newStr[i] = strings.TrimLeft(word, ",")

				if len(newStr[i]) == 0 {
					newStr = append(newStr[:i], newStr[i+1:]...)
				}

			case '!':
				newStr[i-1] = newStr[i-1] + "!"
				newStr[i] = strings.TrimLeft(word, "!")

				if len(newStr[i]) == 0 {
					newStr = append(newStr[:i], newStr[i+1:]...)
				}
				
			case '?':
				newStr[i-1] = newStr[i-1] + "?"
				newStr[i] = strings.TrimLeft(word, "?")

				if len(newStr[i]) == 0 {
					newStr = append(newStr[:i], newStr[i+1:]...)
				}

			case ':':
				newStr[i-1] = newStr[i-1] + ":"
				newStr[i] = strings.TrimLeft(word, ":")

				if len(newStr[i]) == 0 {
					newStr = append(newStr[:i], newStr[i+1:]...)
				}

			case ';':
				newStr[i-1] = newStr[i-1] + ";"
				newStr[i] = strings.TrimLeft(word, ";")

				if len(newStr[i]) == 0 {
					newStr = append(newStr[:i], newStr[i+1:]...)
				}
			}



			result = strings.Join(newStr, " ")
		}
	}

	// here i write to the file
	file.WriteString(result)
}
