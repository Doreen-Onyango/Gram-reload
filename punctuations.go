package reload

func PuncTuation(str []string) []string {
	punc := []string{".", ",", "!", "?", ":", ";"}
	// check for puctuations at the middle of a string slice str
	for index, val := range str {
		for _, value := range punc {
			// check if the val at index0 is equal to the value which is a punctuation and
			// check if that val has the same value at the val[len(val)-1] and
			// check if that val does not appear at the end of the str slice
			if string(val[0]) == value && string(val[len(val)-1]) == value && str[len(str)-1] != str[index] {
				// the current element index-1 is then added with the val and append the result to str
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
	// check for the first quote
	count := 0
	// iterate over the str slice, the index and val being the current element
	for index, val := range str {
		// check if the val is ' and its count is 0 then increase the count to the next index and append the result to the str
		if val == "'" && count == 0 {
			count++
			str[index+1] = val + str[index+1]
			str = append(str[:index], str[index+1:]...)
		}
	}
	// checks for other quotes
	for index, val := range str {
		// check if the val is a ' and if its true then move the ' to index-1 and append the result to the string
		if val == "'" {
			str[index-1] += val
			str = append(str[:index], str[index+1:]...)
		}
	}
	return str
}
