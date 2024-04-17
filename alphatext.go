package reload

import (
	"strconv"
	"strings"
)

// checks for instace of cap or cap, and capitalizes the first letter of the previous slice of string and removes the cap and cap,.
func Capital(str []string) []string {
	for i := 0; i < len(str); i++ {
		if strings.Contains(str[i], "cap") {
			if strings.Contains(str[i], "cap,") {
				num, _ := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				for j := i - num; j < i; j++ {
					str[j] = strings.Title(str[j])
				}
				str = append(str[:i], str[i+2:]...)
			} else {
				str[i-1] = strings.Title(str[i-1])
				str = append(str[:i], str[i+1:]...)
			}
		}
	}
	return str
}

// checks for any instances of up and up, and changes the previous slice of string to uppercase and removes the up and up,.
func UpperCase(str []string) []string {
	for i := 0; i < len(str); i++ {
		if strings.Contains(str[i], "up") {
			if strings.Contains(str[i], "up,") {
				num, _ := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				for k := i - num; k < i; k++ {
					str[k] = strings.ToUpper(str[k])
				}
				str = append(str[:i], str[i+2:]...)
			} else {
				str[i-1] = strings.ToUpper(str[i-1])
				str = append(str[:i], str[i+1:]...)
			}
		}
	}
	return str
}

// checks for instance of low and low, and changes the previous word to lowercase then removes the low and low,.
func LowerCase(str []string) []string {
	for i := 0; i < len(str); i++ {
		if strings.Contains(str[i], "low") {
			if strings.Contains(str[i], "low,") {
				num, _ := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				for j := i - num; j < i; j++ {
					str[j] = strings.ToLower(str[j])
				}
				str = append(str[:i], str[i+2:]...)

			} else {
				str[i-1] = strings.ToLower(str[i-1])
				str = append(str[:i], str[i+1:]...)
			}
		}
	}
	return str
}
