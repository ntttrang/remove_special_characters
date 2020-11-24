package main

import (
	"fmt"

	"github.com/ntttrang/remove_special_characters/handlerString"
)

func main() {
	input := "TUNLAYA-ANUKIT;; RATTANAPENG"
	afterRemoveString := handlerString.RemoveSpecialCharacters(input)
	fmt.Println(afterRemoveString)
}
