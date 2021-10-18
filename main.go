package main

import (
	"fmt"

	"github.com/cortes-/stringmod/quotes"
	"github.com/cortes-/stringmod/strings"
)

func main() {
	o, e := strings.CountOddEven("12345")
	fmt.Println(o, e) // 3 2
	fmt.Println(quotes.GetEmoji("turtle"))
}
