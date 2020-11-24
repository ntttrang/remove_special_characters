package handlerString

import (
	"fmt"
	"strings"
)

// RemoveSpecialCharacters
func RemoveSpecialCharacters(input string) string {
	var readableString string
	/* 	List Special Charactor which are unable to save to DB (Refernece from Profile)
		#  (ASCII: 35)
		~  (ASCII: 126)
		|  (ASCII: 124)
		:; (ASCII: 58 59)
		;; (ASCII: 59 59)
		@@ (ASCII: 64)
		## (ASCII: 35 35)
		#; (ASCII: 35 59)
		^  (ASCII: 94)
		Non-Printable characters:
		ASCII 0->31
		ASCII 127->160
	    ASCII 219->222
	*/
	// Replace Printable characters
	printableReplacerList:= []string{
		"#", "",
		"~", "",
		"|", "",
		":", "",
		";", "",
		"@", "",
		"^", "",
	}
	r := strings.NewReplacer(printableReplacerList...)
	readableString = r.Replace(input)

	var specialNonPrintableASCIIList []rune
	for _, s := range readableString {
		// Replace Non-Printable characters
		isNonPrintable := (s >= 0 && s <= 31) ||               // ASCII 0->31
			              (s >= 127 && s <= 160) ||            // ASCII 127->160
			              (s >= 219 && s <= 222)               // ASCII 219->222
		if isNonPrintable  {
			specialNonPrintableASCIIList = append(specialNonPrintableASCIIList, s)
			readableString = strings.ReplaceAll(readableString, string(s),"")
		}
	}
	fmt.Println("specialNonPrintableASCIIList = ", specialNonPrintableASCIIList)
	fmt.Printf("input=%s, Len=%d\n", input, len(input))
	fmt.Printf("readableString=%s, Len=%d\n", readableString, len(readableString))
	return readableString
}
