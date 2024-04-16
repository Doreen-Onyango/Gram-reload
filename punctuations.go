package reload

// check for punctuations at the middle of a str and relocates it to the required position
func PuncTuation(str []string) []string {
	punc := []string{".", ",", "!", "?", ":", ";"}
	for index, val := range str {
		for _, value := range punc {
			if string(val[0]) == value && string(val[len(val)-1]) == value && str[len(str)-1] != str[index] {
				str[index-1] = str[index-1] + val
				str = append(str[:index], str[index+1:]...)
			}
		}
	}
	return str
}

func PuncTuate(str []string) []string {
	punc := []string{".", ",", "!", "?", ":", ";"}
	// check for punctuation at the end of the str
	for index, val := range str {
		for _, value := range punc {
			if string(val[0]) == value && str[len(str)-1] == str[index] {
				str[index-1] = str[index-1] + val
				str = str[:len(str)-1]
				// check for punctuation at the start of a slice of str
			} else if string(val[0]) == value && string(val[len(val)-1]) != value {
				str[index-1] = str[index-1] + value
				str[index] = val[1:]
			}
		}
	}
	return str
}

func QuOtes(str []string) []string {
	// check for the first quote in a string and combine it with the word next to it
	count := 0
	for index, val := range str {
		if val == "'" && count == 0 {
			count++
			str[index+1] = val + str[index+1]
			str = append(str[:index], str[index+1:]...)
		}
	}
	// checks for other quote in the string and combines it with the previous word
	for index, val := range str {
		if val == "'" {
			str[index-1] += val
			str = append(str[:index], str[index+1:]...)
		}
	}
	return str
}
