// find function

package main

import (
	"fmt"
	"regexp"
)

func main() {

	r, err := regexp.Compile("1")
	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return
	}

	a := []string{"01", "1", "2", "3", "1", "0", "1"}

	if r.MatchString(a[2]) == true {
		fmt.Println("yes")

	} else {
		fmt.Println("no")

	}
}
