package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	file, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))
}
