package main

import (
	"fmt"
	"strings"
)

func main() {
	var replacer = strings.NewReplacer(",", "", ".", "")

	str := "a, space, tab-separated string."
	str = replacer.Replace(str)
	fmt.Println(str)

}
