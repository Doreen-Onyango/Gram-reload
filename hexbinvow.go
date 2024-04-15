package reload

import (
	"fmt"
	"strconv"
)

func HexBinDec(str []string) []string {
	// iterate using for range, over each element of the str slice where the index is our current element and the val i our current value
	for index, val := range str {
		// check if the current element val is equal to string "(hex)"
		if val == "(hex)" {
			// convert the string before the current element at index-1 into integer and initialize the result to nval
			nval, _ := strconv.ParseInt(str[index-1], 16, 64)
			// update the string before val at index-1 with the string representation of nval
			str[index-1] = fmt.Sprint(nval)
			// remove the current element (hex) from the str slice and concatinate the rest of the str slices
			str = append(str[:index], str[index+1:]...)
		} else if val == "(bin)" {
			// check if the current element at val is equal to the string "(bin)"
			// convert the string before the current element at index-1 into an integer and initialize the result to nval
			// update the string at index-1 with the string representation of nval
			nval, _ := strconv.ParseInt(str[index-1], 2, 64)
			str[index-1] = fmt.Sprint(nval)
			// remove the current element "(bin)" from the str slice and concatinate the str slice
			str = append(str[:index], str[index+1:]...)
		}
	}
	return str
}

func GrammerVow(str []string) []string {
	// initialize a slice of strings named vow that contains the vowels and "h"
	vow := []string{"a", "e", "i", "o", "u", "h"}
	// iterate over each element val in the str slice and the index is the element being iterared over
	for index, val := range str {
		// iterate over each element value of the vow slice, in this case ignore the index _,
		for _, value := range vow {
			// checks if the current element val is equal to 'a and if the characterof the next element str[index+1][0] is equal to the curent vow value
			if val == "a" && string(str[index+1][0]) == value {
				// if that condition is true then relpace the val in the str slice with the string 'an'
				// check for 'A' and repeat the same then relace it with 'An'
				str[index] = "an"
			} else if val == "A" && string(str[index+1][0]) == value {
				str[index] = "An"
			}
		}
	}
	return str
}
