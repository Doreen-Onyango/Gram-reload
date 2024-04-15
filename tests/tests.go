package main

import (
	"fmt"
	"os"
	"strings"

	"reload"
)

func main() {
	args := os.Args[1:]
	file, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))
	strs := strings.Fields(string(file))
	strs = reload.CapiTal(strs)
	strs = reload.UpperCase(strs)
	strs = reload.LowerCase(strs)
	strs = reload.HexBinDec(strs)
	strs = reload.GrammerVow(strs)
	final := strings.Join(strs, " ")
	finals := []byte(final)
	os.WriteFile(args[1], finals, 0o644)
}
