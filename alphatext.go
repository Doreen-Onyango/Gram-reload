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
					str = append(str[:i], str[i+2:]...)
				}
			} else {
				str[i-1] = strings.Title(str[i-1])
				str = append(str[:i], str[i+1:]...)
			}
		}
	}
	return str
}
