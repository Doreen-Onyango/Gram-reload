package reload

import (
	"strconv"
	"strings"
)

func CapiTal(str []string) []string {
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
