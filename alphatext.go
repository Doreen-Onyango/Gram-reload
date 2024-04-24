package reload

import (
	"fmt"
	"strconv"
	"strings"
)

// checks for instace of cap or cap, and capitalizes the first letter of the previous slice of string and removes the cap and cap,.
func Capital(str []string) []string {
	for i := 0; i < len(str); i++ {
		if strings.Contains(str[i], "cap") {
			if strings.Contains(str[i], "cap,") {
				num, err := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				if err != nil || num > len(str) || num < 0 {
					fmt.Println("This is an error")
					continue
				}
				for j := i - num; j < i; j++ {
					str[j] = strings.ToUpper(string(str[j][0])) + strings.ToLower(string(str[j][1:]))
				}
				str = append(str[:i], str[i+2:]...)
			} else {
				str[i-1] = strings.ToUpper(string(str[i-1][0])) + strings.ToLower(string(str[i-1][1:]))
				str = append(str[:i], str[i+1:]...)
			}
		}
	}
	return str
}

// checks for any instances of up and up, and changes the previous slice of string to uppercase and removes the up and up,.
func UpperCase(str []string) []string {
	for i := 0; i < len(str); i++ {
		if strings.Contains(str[i], "(up") {
			if strings.Contains(str[i], "(up,") {
				num, err := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				if err != nil || num > len(str) || num < 0 {
					fmt.Println("This is an error")
					continue
				}
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
		if strings.Contains(str[i], "(low") {
			if strings.Contains(str[i], "(low,") {
				num, err := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				if err != nil || num > len(str) || num < 0 {
					fmt.Println("this is an error")
					continue
				}
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
