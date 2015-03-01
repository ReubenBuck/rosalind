// turning DNA into RNA

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	string := os.Args[1]
	r := strings.NewReplacer("T", "U")
	fmt.Println(r.Replace(string))
}
