package reload

import (
	"fmt"
	"strconv"
)

func HexBinDec(str []string) []string {
	// iterate usinf for loop, over each element of the str slice wher the index is our current element and the val i our current value
	for index, val := range str {
		// check if the current element val is equal to string "(hex)"
		if val == "(hex)" {
			// convert the string before the current element at index-1 into integer and assign the result to nval
			nval, _ := strconv.ParseInt(str[index-1], 16, 64)
			// update the string before val at index-1 with the string representation of nval
			str[index-1] = fmt.Sprint(nval)
			// remove the current element (hex) from the str slice and concatinates the rest of the str slices
			str = append(str[:index], str[index+1:]...)
		} else if val == "(bin)" {
			nval, _ := strconv.ParseInt(str[index-1], 2, 64)
			str[index-1] = fmt.Sprint(nval)
			str = append(str[:index], str[index+1:]...)
		}
	}
	return str
}

func GrammerVow(str []string) []string {
	vow := []string{"a", "e", "i", "o", "u", "h"}
	for index, val := range str {
		for _, value := range vow {
			if val == "a" && string(str[index+1][0]) == value {
				str[index] = "an"
			} else if val == "A" && string(str[index+1][0]) == value {
				str[index] = "An"
			}
		}
	}
	return str
}
