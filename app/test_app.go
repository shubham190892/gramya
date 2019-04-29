package main

import (
	"fmt"
	"regexp"
	"strings"
)

var alphaWithComma = regexp.MustCompile(`^[a-zA-Z, ]+$`).MatchString
var dateYYYYmmDD = regexp.MustCompile(`^[0-9]{4}-(0[0-9]|1[0-2])-([0-2][0-9]|3[0-1])$`).MatchString

func main() {
	fmt.Println(alphaWithComma(""))
	fmt.Println(dateYYYYmmDD("1234-02-31"))
	a := "abc,d"
	b := strings.Split(a, ",")
	fmt.Println(b)
	fmt.Println(len(b))

}
