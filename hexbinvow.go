package reload

import (
	"fmt"
	"strconv"
)

func HexBinDec(str []string) []string {
	for index, val := range str {
		if val == "(hex)" {
			nval, _ := strconv.ParseInt(str[index-1], 16, 64)
			str[index-1] = fmt.Sprint(nval)
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
