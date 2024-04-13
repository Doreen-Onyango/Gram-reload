package reload

import (
	"strconv"
	"strings"
)

func CapiTal(str []string) []string {
	// iterate over the length of  the whole slice of str using the i index
	for i := 0; i < len(str); i++ {
		// check if the string str at the i index contains the string "cap" and also the "cap,"
		if strings.Contains(str[i], "cap") {
			if strings.Contains(str[i], "cap,") {
				// initialize a new variable num and an empty error and store the converted integer to it
				num, _ := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				// iterate over the elements of the string from i-num to i-1 and convert each first element to uppercase using strings.Title
				for j := i - num; j < i; j++ {
					str[j] = strings.Title(str[j])
				}
				// remove elements at index i and index i+1 from the slice by appending slice before index i and after index i+2
				str = append(str[:i], str[i+2:]...)
			} else {
				// if the str[i] has "cap" then the slice at index i-1 converts to title and removes the element at index i
				str[i-1] = strings.Title(str[i-1])
				str = append(str[:i], str[i+1:]...)
			}
		}
	}
	return str
}

func UpperCase(str []string) []string {
	// iterate over the string slice using the index i
	for i := 0; i < len(str); i++ {
		// check if the string slice at index i contains "up"
		if strings.Contains(str[i], "up") {
			// check if the string slice at index i contains "up," then convert the string at index i+1 to integer and omit the next index of that slice
			if strings.Contains(str[i], "up,") {
				num, _ := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				// iterate over the elements of the string from i-num to i-1 and convert each first element to uppercase using strings.ToUpper
				for k := i - num; k < i; k++ {
					str[k] = strings.ToUpper(str[k])
				}
				// remove elements at index i and index i+1 from the slice by appending slice before index i and after index i+2
				str = append(str[:i], str[i+2:]...)
			} else {
				str[i-1] = strings.ToUpper(str[i-1])
				str = append(str[:i], str[i+1:]...)
			}
		}
	}
	return str
}

func LowerCase(str []string) []string {
	// iterate through the slice of str usinde the index i
	for i := 0; i < len(str); i++ {
		// check if if the current string in the slice str[i] contais the substring "low" and "low,"
		if strings.Contains(str[i], "low") {
			if strings.Contains(str[i], "low,") {
				// convert the substring str[i+1] to an integer excluding the last character
				num, _ := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				// iterate through the elements of str from i-num to i-1 and convert each element of of that slice str[j] to lowercase using the strings.ToLower
				for j := i - num; j < i; j++ {
					str[j] = strings.ToLower(str[j])
				}
				// remove elements at index i and index i+1 from the slice by appending slice before index i and after index i+2
				str = append(str[:i], str[i+2:]...)

			} else {
				str[i-1] = strings.ToLower(str[i-1])
				str = append(str[:i], str[i+1:]...)
			}
		}
	}
	return str
}
