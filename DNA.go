// How can I calculate how many of each charcter are in a string

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	string := os.Args[1]
	fmt.Println("total length =", len(string))
	fmt.Println("A =", strings.Count(string, "A"))
	fmt.Println("C =", strings.Count(string, "C"))
	fmt.Println("T =", strings.Count(string, "T"))
	fmt.Println("G =", strings.Count(string, "G"))
	fmt.Println(strings.Count(string, "A"), strings.Count(string, "C"), strings.Count(string, "G"), strings.Count(string, "T"))

}
