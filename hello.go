package main

import (
	"fmt"
	"github.com/ninjrok/stringutil"
)

//Prints hello world and reverses a string using stringutils package
func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("!oG ,olleH reversed is:- ")
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Printf("\n")
}
