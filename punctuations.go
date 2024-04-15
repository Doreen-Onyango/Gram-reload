package reload

func PuncTuation(str []string) []string {
	punc := []string{".", ",", "!", "?", ":", ";"}
	for index, val := range str {
		for _, value := range punc {
			// check for punctuations at the middle of a string
			if string(val[0]) == value && string(val[len(val)-1]) == value && str[len(str)-1] != str[index] {
				str[index-1] = str[index-1] + val
				str = append(str[:index], str[index+1:]...)
			}
		}
	}
	return str
}

// checks for punctuations at the end of a string
func PuncTuate(str []string) []string {
	punc := []string{".", ",", "!", "?", ":", ";"}
	for index, val := range str {
		for _, value := range punc {
			if string(val[0]) == value && str[len(str)-1] == str[index] {
				str[index-1] = str[index-1] + val
				str = str[:len(str)-1]
			} else if string(val[0]) == value && string(val[len(val)-1]) != value {
				str[index-1] = str[index-1] + value
				str[index] = val[1:]
			}
		}
	}
	return str
}

func QuOtes(str []string) []string {
	count := 0
	for index, val := range str {
		if val == "'" && count == 0 {
			count++
			str[index+1] = val + str[index+1]
			str = append(str[:index], str[index+1:]...)
		}
	}
	for index, val := range str {
		if val == "'" {
			str[index-1] += val
			str = append(str[:index], str[index+1:]...)
		}
	}
	return str
}
