package main

import (
	"fmt"
	"strings"
	"unicode"
)

func stringIsUpper(input string) bool {
	a := []rune(input)

	for i := 0; i < len(a); i++ {
		ch := a[i]
		if unicode.IsLetter(ch) {
			if unicode.IsUpper(ch) == false {
				return false
			}
		}

	}
	return true
}

//func fun1(input string) string {
//	if input[len(input)-1:] == "?" {
//		return "Sure."
//	} else if len(strings.TrimSpace(input)) == 0 {
//		return "Fine. Be that way!"
//	} else if stringIsUpper(input) && input[len(input)-1:] == "?" {
//		return "Calm down, I know what I'm doing!"
//	} else if stringIsUpper(input) {
//		return "Whoa, chill out!"
//	} else {
//		return "Whatever."
//	}
//}
func fun1(input string) interface{} {
	switch {
	case len(strings.TrimSpace(input)) == 0:
		return "Fine. Be that way!"
	case strings.ToUpper(input) == input && input[len(input)-1] == '?':
		return "Calm down, I know what I'm doing!"
	case strings.ToUpper(input) == input:
		return "Whoa, chill out!"
	case input[len(input)-1] == '?':
		return "Sure"
	default:
		return "Whatever"
	}
}
func main() {
	fmt.Println(fun1("How are you?"))
	fmt.Println(fun1(""))
	fmt.Println(fun1("HEYY?"))
	fmt.Println(fun1("HELLO"))
	fmt.Println(fun1("What's poppin?"))
}
