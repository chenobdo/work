package main

import (
	"fmt"
)

func main() {
	var s = "goas"
	var r = []rune(s)
	fmt.Printf("%c\n", s[1])
	fmt.Printf("%c\n", r[1])
	fmt.Printf("%c\n", s[3])
	fmt.Printf("%c\n", r[3])
}
